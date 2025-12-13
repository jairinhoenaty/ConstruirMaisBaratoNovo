package budget_repository_impl

import (
	"fmt"
	"time"

	"gorm.io/gorm"

	pkgpbudget "construir_mais_barato/app/domain/budget"
	pkgprofessional "construir_mais_barato/app/domain/professional"
	pkgstore "construir_mais_barato/app/domain/store"
)

type repository struct {
	DB *gorm.DB
}

func NewBudgetRepositoryImpl(db *gorm.DB) pkgpbudget.BudgetRepository {
	return &repository{
		DB: db,
	}
}

func (r *repository) FindExpiredBudgets(before time.Time) ([]*pkgpbudget.Budget, error) {
	var expiredBudgets []*pkgpbudget.Budget
	err := r.DB.Where("created_at < ?", before).Find(&expiredBudgets).Error

	return expiredBudgets, err
}

func (r *repository) FindBudgetsByMonthAndProfessionalID(
	month string,
	professionalID uint,
	storeID uint,
	clientID int,
	page int,
	pageSize int,
) ([]*pkgpbudget.Budget, error) {
	var budgets []*pkgpbudget.Budget
	offset := (page - 1) * pageSize

	monthNumber := getMonthNumber(month)
	if monthNumber == "" {
		return nil, fmt.Errorf("mês inválido: %s", month)
	}

	// currentYear := time.Now().Year()
	// yearMonth := fmt.Sprintf("%d-%s", currentYear, monthNumber)

	// construir condicoes e args de forma segura
	tx := r.DB.Model(&pkgpbudget.Budget{})

	// JOIN dinâmico
	if professionalID != 0 {
		tx = tx.Joins("JOIN budgets_professionals ON budgets.id = budgets_professionals.budget_id")
	}

	if storeID != 0 {
		tx = tx.Joins("JOIN budgets_stores ON budgets.id = budgets_stores.budget_id")
	}

	// Filtros
	if professionalID != 0 {
		tx = tx.Where("budgets_professionals.professional_id = ?", professionalID)
	}
	if storeID != 0 {
		tx = tx.Where("budgets_stores.store_id = ?", storeID)
	}
	if clientID != 0 {
		tx = tx.Where("budgets.client_id = ?", clientID)
	}
	if storeID == 0 && professionalID == 0 {
		return nil, fmt.Errorf("Nenhum orçamento encontrado")
	}
	// Filtro por mês
	// tx = tx.Where("DATE_FORMAT(budgets.created_at, '%Y-%m') = ?", yearMonth)

	// Só aprovados
	tx = tx.Where("budgets.approved = ?", true)

	// Preloads
	tx = tx.
		Preload("City").
		Preload("Professionals").
		Preload("Professionals.Professions").
		Preload("Professionals.City")

	// Execução final
	if err := tx.
		Distinct("budgets.*").
		Order("budgets.created_at DESC").
		Limit(pageSize).
		Offset(offset).
		Find(&budgets).
		Error; err != nil {
		return nil, err
	}

	return budgets, nil
}

func (r *repository) FindAll(limit, offset int) ([]*pkgpbudget.Budget, int64, error) {
	var total int64

	// Contagem total de orçamentos
	if err := r.DB.Model(&pkgpbudget.Budget{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	var budgets []*pkgpbudget.Budget

	if err := r.DB.
		Preload("Professionals").
		Preload("Stores").
		Preload("Professionals.Professions").
		Preload("Professionals.City").
		Preload("City").
		Distinct("budgets.*").
		// Preload("Client").
		// Preload("Client.City").
		Where("deleted_at IS NULL").
		Limit(limit).
		Offset(offset).
		Order("id DESC").
		Find(&budgets).Error; err != nil {
		return nil, 0, err
	}
	return budgets, total, nil
}

func (r *repository) FindById(id uint) (*pkgpbudget.Budget, error) {
	budget := pkgpbudget.Budget{}
	if err := r.DB.
		Preload("Professionals").
		Preload("Professionals.Professions").
		Where("deleted_at IS NULL").
		First(&budget, id).Error; err != nil {
		return nil, err
	}
	return &budget, nil
}

func (r *repository) FindByEmail(email string) (*pkgpbudget.Budget, error) {
	budget := pkgpbudget.Budget{}
	if err := r.DB.
		Preload("Professionals").
		Preload("Professionals.Professions").
		Preload("City").
		Where("email = ? AND deleted_at IS NULL", email).
		First(&budget).Error; err != nil {
		return nil, err
	}
	return &budget, nil
}

func (r *repository) Save(budget pkgpbudget.Budget) (*pkgpbudget.Budget, error) {
	err := r.DB.Transaction(func(tx *gorm.DB) error {
		// Salve o orçamento primeiro
		if err := tx.Save(&budget).Error; err != nil {
			return err
		}

		// Se houver IDs de profissionais, associe-os ao orçamento
		if budget.ProfessionalIDs != nil && len(*budget.ProfessionalIDs) > 0 {
			var professionals []pkgprofessional.Professional
			if err := tx.Where("id IN ?", *budget.ProfessionalIDs).Find(&professionals).Error; err != nil {
				return err
			}
			if err := tx.Model(&budget).Association("Professionals").Replace(&professionals); err != nil {
				return err
			}
		} else if budget.StoresIDs != nil && len(*budget.StoresIDs) > 0 {
			var Stores []pkgstore.Store
			if err := tx.Where("id IN ?", *budget.StoresIDs).Find(&Stores).Error; err != nil {
				return err
			}
			if err := tx.Model(&budget).Association("Stores").Replace(&Stores); err != nil {
				return err
			}

		}

		return nil
	})

	if err != nil {
		return nil, err
	}
	return &budget, nil

}

func (r *repository) Remove(id uint) error {

	// Primeiro, obtém o orçamento que você quer remover
	var budget pkgpbudget.Budget
	r.DB.Preload("Professionals").First(&budget, id)

	// Remover o relacionamento many-to-many entre Budget e Professionals
	// Remove os profissionais associados do orçamento
	r.DB.Model(&budget).Association("Professionals").Clear()

	if err := r.DB.Delete(&pkgpbudget.Budget{}, id).Error; err != nil {
		return err
	}
	return nil
}

var monthMap = map[string]string{
	"January":   "01",
	"February":  "02",
	"March":     "03",
	"April":     "04",
	"May":       "05",
	"June":      "06",
	"July":      "07",
	"August":    "08",
	"September": "09",
	"October":   "10",
	"November":  "11",
	"December":  "12",
}

func getMonthNumber(month string) string {
	if num, ok := monthMap[month]; ok {
		return num
	}
	return ""
}
