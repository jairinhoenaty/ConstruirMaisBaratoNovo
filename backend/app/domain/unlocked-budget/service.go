package unlockedbudget

type UnlockedBudgetService interface {
	FindByProfessionalAndBudget(professionalID uint, budgetID uint) (*UnlockedBudget, error)
	FindPaidByProfessionalAndBudget(professionalID uint, budgetID uint) (*UnlockedBudget, error)
	FindByPaymentID(paymentID string) (*UnlockedBudget, error)
	FindAllByProfessional(professionalID uint) ([]*UnlockedBudget, error)
	Save(unlockedBudget UnlockedBudget) (*UnlockedBudget, error)
	Update(unlockedBudget *UnlockedBudget) error
}

type unlockedBudgetService struct {
	repository UnlockedBudgetRepository
}

func NewUnlockedBudgetService(repository UnlockedBudgetRepository) UnlockedBudgetService {
	return &unlockedBudgetService{
		repository: repository,
	}
}

func (s *unlockedBudgetService) FindByProfessionalAndBudget(professionalID uint, budgetID uint) (*UnlockedBudget, error) {
	return s.repository.FindByProfessionalAndBudget(professionalID, budgetID)
}

func (s *unlockedBudgetService) FindPaidByProfessionalAndBudget(professionalID uint, budgetID uint) (*UnlockedBudget, error) {
	return s.repository.FindPaidByProfessionalAndBudget(professionalID, budgetID)
}

func (s *unlockedBudgetService) FindByPaymentID(paymentID string) (*UnlockedBudget, error) {
	return s.repository.FindByPaymentID(paymentID)
}

func (s *unlockedBudgetService) FindAllByProfessional(professionalID uint) ([]*UnlockedBudget, error) {
	return s.repository.FindAllByProfessional(professionalID)
}

func (s *unlockedBudgetService) Save(unlockedBudget UnlockedBudget) (*UnlockedBudget, error) {
	return s.repository.Save(unlockedBudget)
}

func (s *unlockedBudgetService) Update(unlockedBudget *UnlockedBudget) error {
	return s.repository.Update(unlockedBudget)
}
