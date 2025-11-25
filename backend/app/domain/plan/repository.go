package plan

type PlanRepository interface {
	FindAll() ([]*Plan, error)
	FindByID(id uint) (*Plan, error)
	FindByUserType(userType UserType) (*Plan, error)
	FindAllActive() ([]*Plan, error)
}
