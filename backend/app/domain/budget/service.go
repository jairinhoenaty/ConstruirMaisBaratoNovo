package budget

import "time"

type BudgetService interface {
	FindAll(limit int, offset int) ([]*Budget, int64, error)
	FindById(id uint) (*Budget, error)
	FindByEmail(email string) (*Budget, error)
	FindBudgetsByMonthAndProfessionalID(month string, professionalID uint, clientID int, page int, pageSize int) ([]*Budget, error)
	FindExpiredBudgets(before time.Time) ([]*Budget, error)
	Save(budget Budget) (*Budget, error)
	Remove(id uint) error
}

type budgetService struct {
	repository BudgetRepository
}

func NewBudgetService(repository BudgetRepository) BudgetService {
	return &budgetService{
		repository: repository,
	}
}

func (s *budgetService) FindExpiredBudgets(before time.Time) ([]*Budget, error) {
	budgets, err := s.repository.FindExpiredBudgets(before)
	if err != nil {
		return nil, err
	}
	return budgets, nil
}

func (s *budgetService) FindBudgetsByMonthAndProfessionalID(month string, professionalID uint, clientID int, page int, pageSize int) ([]*Budget, error) {
	budgets, err := s.repository.FindBudgetsByMonthAndProfessionalID(month, professionalID, clientID, page, pageSize)
	if err != nil {
		return nil, err
	}
	return budgets, nil
}

func (s *budgetService) FindAll(limit int, offset int) ([]*Budget, int64, error) {
	budgets, total, err := s.repository.FindAll(limit, offset)
	if err != nil {
		return nil, total, err
	}
	return budgets, total, nil
}

func (s *budgetService) FindById(id uint) (*Budget, error) {
	budget, err := s.repository.FindById(id)
	if err != nil {
		return nil, err
	}
	return budget, nil
}

func (s *budgetService) FindByEmail(email string) (*Budget, error) {
	budget, err := s.repository.FindByEmail(email)
	if err != nil {
		return nil, err
	}
	return budget, nil
}

func (s *budgetService) Save(budget Budget) (*Budget, error) {
	newBudget, err := s.repository.Save(budget)
	if err != nil {
		return nil, err
	}
	return newBudget, nil
}

func (s *budgetService) Remove(id uint) error {
	return s.repository.Remove(id)
}
