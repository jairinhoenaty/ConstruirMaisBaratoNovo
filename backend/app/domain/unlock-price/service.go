package unlockprice

type UnlockPriceService interface {
	FindByUserType(userType UserType) (*UnlockPrice, error)
	FindActiveByUserType(userType UserType) (*UnlockPrice, error)
	FindAll() ([]*UnlockPrice, error)
	Save(unlockPrice UnlockPrice) (*UnlockPrice, error)
	Update(unlockPrice *UnlockPrice) error
}

type unlockPriceService struct {
	repository UnlockPriceRepository
}

func NewUnlockPriceService(repository UnlockPriceRepository) UnlockPriceService {
	return &unlockPriceService{
		repository: repository,
	}
}

func (s *unlockPriceService) FindByUserType(userType UserType) (*UnlockPrice, error) {
	return s.repository.FindByUserType(userType)
}

func (s *unlockPriceService) FindActiveByUserType(userType UserType) (*UnlockPrice, error) {
	return s.repository.FindActiveByUserType(userType)
}

func (s *unlockPriceService) FindAll() ([]*UnlockPrice, error) {
	return s.repository.FindAll()
}

func (s *unlockPriceService) Save(unlockPrice UnlockPrice) (*UnlockPrice, error) {
	return s.repository.Save(unlockPrice)
}

func (s *unlockPriceService) Update(unlockPrice *UnlockPrice) error {
	return s.repository.Update(unlockPrice)
}
