package plan

type PlanService interface {
	FindAll() ([]*Plan, error)
	FindByID(id uint) (*Plan, error)
	FindByUserType(userType UserType) (*Plan, error)
	FindAllActive() ([]*Plan, error)
}

type planService struct {
	repository PlanRepository
}

func NewPlanService(repository PlanRepository) PlanService {
	return &planService{
		repository: repository,
	}
}

func (s *planService) FindAll() ([]*Plan, error) {
	plans, err := s.repository.FindAll()
	if err != nil {
		return nil, err
	}
	return plans, nil
}

func (s *planService) FindByID(id uint) (*Plan, error) {
	plan, err := s.repository.FindByID(id)
	if err != nil {
		return nil, err
	}
	return plan, nil
}

func (s *planService) FindByUserType(userType UserType) (*Plan, error) {
	plan, err := s.repository.FindByUserType(userType)
	if err != nil {
		return nil, err
	}
	return plan, nil
}

func (s *planService) FindAllActive() ([]*Plan, error) {
	plans, err := s.repository.FindAllActive()
	if err != nil {
		return nil, err
	}
	return plans, nil
}
