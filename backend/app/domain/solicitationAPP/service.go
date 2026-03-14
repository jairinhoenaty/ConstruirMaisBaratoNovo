package solicitationapp

type SolicitationAppService interface {
	Save(solicitationApp SolicitationApp) (*SolicitationApp, error)
}

type solicitationAppService struct {
	repository SolicitationAppRepository
}

func NewSolicitationAppService(repository SolicitationAppRepository) SolicitationAppService {
	return &solicitationAppService{
		repository: repository,
	}
}

func (s *solicitationAppService) Save(solicitationApp SolicitationApp) (*SolicitationApp, error) {
	newsolicitationApp, err := s.repository.Save(solicitationApp)
	if err != nil {
		return nil, err
	}
	return newsolicitationApp, nil
}
