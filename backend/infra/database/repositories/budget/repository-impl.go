package budget_repository_impl

import (
	"fmt"
	"strconv"
	"time"

	"gorm.io/gorm"

	pkgpbudget "construir_mais_barato/app/domain/budget"
	pkgprofessional "construir_mais_barato/app/domain/professional"
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

func (r *repository) FindBudgetsByMonthAndProfessionalID(month string, professionalID uint, clientID int, page int, pageSize int) ([]*pkgpbudget.Budget, error) {
	var budgets []*pkgpbudget.Budget
	offset := (page - 1) * pageSize

	// Obter o número do mês
	monthNumber := getMonthNumber(month)

	if monthNumber == "" {
		return nil, fmt.Errorf("mês inválido: %s", month)
	}

	// Obter o ano corrente
	currentYear := time.Now().Year()
	// Formatar a string do ano e mês
	yearMonth := fmt.Sprintf("%d-%s", currentYear, monthNumber)

	fmt.Println("REPOSITORIO => ano e mes da consulta => ", yearMonth)
	filtro := "0=0"

	if clientID != 0 {
		filtro = filtro + " and client_id=" + strconv.Itoa(clientID)
	}
	if professionalID != 0 {
		filtro = filtro + " and budgets_professionals.professional_id=" + strconv.FormatUint(uint64(professionalID), 10)

	}

	err := r.DB.
		Joins("JOIN budgets_professionals ON budgets.id = budgets_professionals.budget_id").
		//Where("budgets_professionals.professional_id = ?", professionalID).
		//Where("NOT budgets.deleted_at IS NULL").
		//Where("DATE_FORMAT(budgets.created_at, '%Y-%m') = ?", yearMonth).
		Where(filtro).
		Where("approved", true).
		Preload("City").
		Preload("Client").
		Preload("Client.City").
		Preload("Professionals").
		Preload("Professionals.Professions").
		Preload("Professionals.City").
		Order("created_at desc").
		Limit(pageSize).
		Offset(offset).
		Find(&budgets).Debug().Error // Adiciona .Debug() para imprimir a consulta SQL gerada

	if err != nil {
		return nil, err
	}
	// if err := r.DB.
	// 	Joins("JOIN budgets_professionals ON budgets.id = budgets_professionals.budget_id").
	// 	Where("budgets_professionals.professional_id = ?", professionalID).
	// 	Where("budgets.deleted_at IS NULL").
	// 	Where(gorm.Expr("DATE_FORMAT(budgets.created_at, '%Y-%m') = ?", fmt.Sprintf("2024-%s", monthNumber))).
	// 	Preload("Professionals").
	// 	Preload("Professionals").
	// 	Preload("Professionals.Professions").
	// 	Limit(pageSize).
	// 	Offset(offset).
	// 	Find(&budgets).Error; err != nil {
	// 	return nil, err
	// }
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
		Preload("Professionals.Professions").
		Preload("Professionals.City").
		Preload("City").
		Preload("Client").
		Preload("Client.City").
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
		if budget.ProfessionalIDs != nil {
			var professionals []pkgprofessional.Professional
			if err := tx.Where("id IN ?", *budget.ProfessionalIDs).Find(&professionals).Error; err != nil {
				return err
			}
			if err := tx.Model(&budget).Association("Professionals").Replace(&professionals); err != nil {
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
