package unlockedbudget

type UnlockedBudgetRepository interface {
	// Professional methods
	FindByProfessionalAndBudget(professionalID uint, budgetID uint) (*UnlockedBudget, error)
	FindPaidByProfessionalAndBudget(professionalID uint, budgetID uint) (*UnlockedBudget, error)
	FindAllByProfessional(professionalID uint) ([]*UnlockedBudget, error)

	// Store methods
	FindByStoreAndBudget(storeID uint, budgetID uint) (*UnlockedBudget, error)
	FindPaidByStoreAndBudget(storeID uint, budgetID uint) (*UnlockedBudget, error)
	FindAllByStore(storeID uint) ([]*UnlockedBudget, error)

	// Common methods
	FindByPaymentID(paymentID string) (*UnlockedBudget, error)
	Save(unlockedBudget UnlockedBudget) (*UnlockedBudget, error)
	Update(unlockedBudget *UnlockedBudget) error
}
