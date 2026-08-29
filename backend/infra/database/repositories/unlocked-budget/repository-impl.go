package unlocked_budget_repository_impl

import (
	unlockedbudget "construir_mais_barato/app/domain/unlocked-budget"

	"gorm.io/gorm"
)

type repository struct {
	DB *gorm.DB
}

func NewUnlockedBudgetRepositoryImpl(db *gorm.DB) unlockedbudget.UnlockedBudgetRepository {
	return &repository{
		DB: db,
	}
}

func (r *repository) FindByProfessionalAndBudget(professionalID uint, budgetID uint) (*unlockedbudget.UnlockedBudget, error) {
	var unlocked unlockedbudget.UnlockedBudget
	err := r.DB.Where("professional_id = ? AND budget_id = ?", professionalID, budgetID).First(&unlocked).Error
	if err != nil {
		return nil, err
	}
	return &unlocked, nil
}

func (r *repository) FindPaidByProfessionalAndBudget(professionalID uint, budgetID uint) (*unlockedbudget.UnlockedBudget, error) {
	var unlocked unlockedbudget.UnlockedBudget
	err := r.DB.Where("professional_id = ? AND budget_id = ? AND status = ?", professionalID, budgetID, "paid").First(&unlocked).Error
	if err != nil {
		return nil, err
	}
	return &unlocked, nil
}

func (r *repository) FindByPaymentID(paymentID string) (*unlockedbudget.UnlockedBudget, error) {
	var unlocked unlockedbudget.UnlockedBudget
	err := r.DB.Where("payment_id = ?", paymentID).First(&unlocked).Error
	if err != nil {
		return nil, err
	}
	return &unlocked, nil
}

func (r *repository) FindByStatusToken(statusToken string) (*unlockedbudget.UnlockedBudget, error) {
	var unlocked unlockedbudget.UnlockedBudget
	err := r.DB.Where("status_token = ?", statusToken).First(&unlocked).Error
	if err != nil {
		return nil, err
	}
	return &unlocked, nil
}

func (r *repository) FindAllByProfessional(professionalID uint) ([]*unlockedbudget.UnlockedBudget, error) {
	var unlocked []*unlockedbudget.UnlockedBudget
	err := r.DB.Where("professional_id = ?", professionalID).Order("created_at DESC").Find(&unlocked).Error
	if err != nil {
		return nil, err
	}
	return unlocked, nil
}

func (r *repository) FindByStoreAndBudget(storeID uint, budgetID uint) (*unlockedbudget.UnlockedBudget, error) {
	var unlocked unlockedbudget.UnlockedBudget
	err := r.DB.Where("store_id = ? AND budget_id = ?", storeID, budgetID).First(&unlocked).Error
	if err != nil {
		return nil, err
	}
	return &unlocked, nil
}

func (r *repository) FindPaidByStoreAndBudget(storeID uint, budgetID uint) (*unlockedbudget.UnlockedBudget, error) {
	var unlocked unlockedbudget.UnlockedBudget
	err := r.DB.Where("store_id = ? AND budget_id = ? AND status = ?", storeID, budgetID, "paid").First(&unlocked).Error
	if err != nil {
		return nil, err
	}
	return &unlocked, nil
}

func (r *repository) FindAllByStore(storeID uint) ([]*unlockedbudget.UnlockedBudget, error) {
	var unlocked []*unlockedbudget.UnlockedBudget
	err := r.DB.Where("store_id = ?", storeID).Order("created_at DESC").Find(&unlocked).Error
	if err != nil {
		return nil, err
	}
	return unlocked, nil
}

func (r *repository) Save(unlockedBudget unlockedbudget.UnlockedBudget) (*unlockedbudget.UnlockedBudget, error) {
	var existingUnlocked unlockedbudget.UnlockedBudget

	// Verificar se o ID está presente para decidir entre atualizar ou criar
	if unlockedBudget.ID != 0 {
		// Tentar encontrar o registro existente
		if err := r.DB.Where("id = ?", unlockedBudget.ID).First(&existingUnlocked).Error; err != nil {
			return nil, err
		}
	}

	err := r.DB.Transaction(func(tx *gorm.DB) error {
		// Se o registro existente for encontrado, atualizar
		if existingUnlocked.ID != 0 {
			if err := tx.Model(&existingUnlocked).Updates(unlockedBudget).Error; err != nil {
				return err
			}
			unlockedBudget.ID = existingUnlocked.ID
		} else {
			// Criar um novo registro
			if err := tx.Create(&unlockedBudget).Error; err != nil {
				return err
			}
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return &unlockedBudget, nil
}

func (r *repository) Update(unlockedBudget *unlockedbudget.UnlockedBudget) error {
	return r.DB.Save(unlockedBudget).Error
}
