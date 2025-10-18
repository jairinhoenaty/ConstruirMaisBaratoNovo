package unlockedbudget

type UnlockedBudgetRepository interface {
	FindByProfessionalAndBudget(professionalID uint, budgetID uint) (*UnlockedBudget, error)

	FindPaidByProfessionalAndBudget(professionalID uint, budgetID uint) (*UnlockedBudget, error)

	FindByPaymentID(paymentID string) (*UnlockedBudget, error)

	FindAllByProfessional(professionalID uint) ([]*UnlockedBudget, error)

	Save(unlockedBudget UnlockedBudget) (*UnlockedBudget, error)

	Update(unlockedBudget *UnlockedBudget) error
}
