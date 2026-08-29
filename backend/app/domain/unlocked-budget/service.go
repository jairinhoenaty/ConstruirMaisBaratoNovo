package unlockedbudget

type UnlockedBudgetService interface {
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
	FindByStatusToken(statusToken string) (*UnlockedBudget, error)
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

func (s *unlockedBudgetService) FindByStatusToken(statusToken string) (*UnlockedBudget, error) {
	return s.repository.FindByStatusToken(statusToken)
}

func (s *unlockedBudgetService) FindAllByProfessional(professionalID uint) ([]*UnlockedBudget, error) {
	return s.repository.FindAllByProfessional(professionalID)
}

func (s *unlockedBudgetService) FindByStoreAndBudget(storeID uint, budgetID uint) (*UnlockedBudget, error) {
	return s.repository.FindByStoreAndBudget(storeID, budgetID)
}

func (s *unlockedBudgetService) FindPaidByStoreAndBudget(storeID uint, budgetID uint) (*UnlockedBudget, error) {
	return s.repository.FindPaidByStoreAndBudget(storeID, budgetID)
}

func (s *unlockedBudgetService) FindAllByStore(storeID uint) ([]*UnlockedBudget, error) {
	return s.repository.FindAllByStore(storeID)
}

func (s *unlockedBudgetService) Save(unlockedBudget UnlockedBudget) (*UnlockedBudget, error) {
	return s.repository.Save(unlockedBudget)
}

func (s *unlockedBudgetService) Update(unlockedBudget *UnlockedBudget) error {
	return s.repository.Update(unlockedBudget)
}
