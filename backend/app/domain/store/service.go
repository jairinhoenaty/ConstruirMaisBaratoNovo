package store

type StoreService interface {
	FindAll(limit, offset int) ([]*Store, int64, error)
	FindById(id uint) (*Store, error)
	FindByName(email string) ([]*Store, error)

	FindByEmail(email string) (*Store, error)
	/*
		FindByName(email string) ([]*Store, error)
		FindByCityAndProfession(cityID, professionID uint, limit, offset int) ([]*Store, int64, error)
		FindByNameAndCityAndProfession(name string, cityID, professionID uint, limit, offset int) ([]*Store, error)
		CountStoresByProfession() ([]ProfessionCount, error)
		CountStoresByState(uf string, limit, offset int) ([]CityStoreCount, *int64, error)
		CountStoresByProfessionInCity(cityID uint) ([]ProfessionCount, error)
	*/
	FindLastStores(quantityRecords int) ([]Store, error)

	Save(store Store) (*Store, error)
	Remove(id uint) error
	ExportXLSX() ([]*Store, error)
}

type storeService struct {
	repository StoreRepository
}

func NewStoreService(repository StoreRepository) StoreService {
	return &storeService{
		repository: repository,
	}
}

func (s *storeService) ExportXLSX() ([]*Store, error) {
	stores, err := s.repository.ExportXLSX()
	if err != nil {
		return nil, err
	}
	return stores, nil
}

/*
	func (s *storeService) FindByNameAndCityAndProfession(name string, cityID, professionID uint, limit, offset int) ([]*Store, error) {
		stores, err := s.repository.FindByNameAndCityAndProfession(name, cityID, professionID, limit, offset)
		if err != nil {
			return nil, err
		}
		return stores, nil
	}

	func (s *storeService) CountStoresByProfessionInCity(cityID uint) ([]ProfessionCount, error) {
		results, err := s.repository.CountStoresByProfessionInCity(cityID)
		if err != nil {
			return nil, err
		}
		return results, nil
	}

	func (s *storeService) CountStoresByState(uf string, limit, offset int) ([]CityStoreCount, *int64, error) {
		results, total, err := s.repository.CountStoresByState(uf, limit, offset)
		if err != nil {
			return nil, nil, err
		}
		return results, total, nil
	}

	func (s *storeService) CountStoresByProfession() ([]ProfessionCount, error) {
		results, err := s.repository.CountStoresByProfession()
		if err != nil {
			return nil, err
		}
		return results, nil
	}
*/
func (s *storeService) FindLastStores(quantityRecords int) ([]Store, error) {
	stores, err := s.repository.FindLastStores(quantityRecords)
	if err != nil {
		return nil, err
	}
	return stores, nil
}

/*
	func (s *storeService) FindByCityAndProfession(cityID, professionID uint, limit, offset int) ([]*Store, int64, error) {
		stores, total, err := s.repository.FindByCityAndProfession(cityID, professionID, limit, offset)
		if err != nil {
			return nil, 0, err
		}
		return stores, total, nil
	}
*/
func (s *storeService) FindAll(limit, offset int) ([]*Store, int64, error) {
	stores, count, err := s.repository.FindAll(limit, offset)
	if err != nil {
		return nil, 0, err
	}
	return stores, count, nil
}

func (s *storeService) FindById(id uint) (*Store, error) {
	store, err := s.repository.FindById(id)
	if err != nil {
		return nil, err
	}
	return store, nil
}

func (s *storeService) FindByEmail(email string) (*Store, error) {
	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *storeService) FindByName(name string) ([]*Store, error) {
	stores, err := s.repository.FindByName(name)
	if err != nil {
		return nil, err
	}
	return stores, nil
}

func (s *storeService) Save(store Store) (*Store, error) {
	newstore, err := s.repository.Save(store)
	if err != nil {
		return nil, err
	}
	return newstore, nil
}

func (s *storeService) Remove(id uint) error {
	return s.repository.Remove(id)
}
