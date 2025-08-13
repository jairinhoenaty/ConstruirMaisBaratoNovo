package budget

import "time"

type BudgetRepository interface {
	FindAll(limit int, offset int) ([]*Budget, int64, error)
	FindById(CityID uint) (*Budget, error)
	FindByEmail(email string) (*Budget, error)
	FindBudgetsByMonthAndProfessionalID(month string, professionalID uint, clientID int, page int, pageSize int) ([]*Budget, error)
	FindExpiredBudgets(before time.Time) ([]*Budget, error)
	Save(orcamento Budget) (*Budget, error)
	Remove(id uint) error
}
