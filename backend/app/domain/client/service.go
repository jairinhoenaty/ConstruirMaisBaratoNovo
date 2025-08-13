package client

type ClientService interface {
	FindAll(limit, offset int) ([]*Client, int64, error)
	FindById(id uint) (*Client, error)
	FindByEmail(email string) (*Client, error)
	FindByName(email string) ([]*Client, error)
	/*
		FindByCityAndProfession(cityID, professionID uint, limit, offset int) ([]*Client, int64, error)
		FindByNameAndCityAndProfession(name string, cityID, professionID uint, limit, offset int) ([]*Client, error)
		CountClientsByProfession() ([]ProfessionCount, error)
		CountClientsByState(uf string, limit, offset int) ([]CityClientCount, *int64, error)
		CountClientsByProfessionInCity(cityID uint) ([]ProfessionCount, error)
	*/
	FindLastClients(quantityRecords int) ([]Client, error)
	Save(client Client) (*Client, error)

	Remove(id uint) error
	ExportXLSX() ([]*Client, error)
}

type clientService struct {
	repository ClientRepository
}

func NewClientService(repository ClientRepository) ClientService {
	return &clientService{
		repository: repository,
	}
}

func (s *clientService) ExportXLSX() ([]*Client, error) {
	clients, err := s.repository.ExportXLSX()
	if err != nil {
		return nil, err
	}
	return clients, nil
}

/*
	func (s *clientService) FindByNameAndCityAndProfession(name string, cityID, professionID uint, limit, offset int) ([]*Client, error) {
		clients, err := s.repository.FindByNameAndCityAndProfession(name, cityID, professionID, limit, offset)
		if err != nil {
			return nil, err
		}
		return clients, nil
	}

	func (s *clientService) CountClientsByProfessionInCity(cityID uint) ([]ProfessionCount, error) {
		results, err := s.repository.CountClientsByProfessionInCity(cityID)
		if err != nil {
			return nil, err
		}
		return results, nil
	}

	func (s *clientService) CountClientsByState(uf string, limit, offset int) ([]CityClientCount, *int64, error) {
		results, total, err := s.repository.CountClientsByState(uf, limit, offset)
		if err != nil {
			return nil, nil, err
		}
		return results, total, nil
	}

	func (s *clientService) CountClientsByProfession() ([]ProfessionCount, error) {
		results, err := s.repository.CountClientsByProfession()
		if err != nil {
			return nil, err
		}
		return results, nil
	}
*/
func (s *clientService) FindLastClients(quantityRecords int) ([]Client, error) {
	clients, err := s.repository.FindLastClients(quantityRecords)
	if err != nil {
		return nil, err
	}
	return clients, nil
}

/*
	func (s *clientService) FindByCityAndProfession(cityID, professionID uint, limit, offset int) ([]*Client, int64, error) {
		clients, total, err := s.repository.FindByCityAndProfession(cityID, professionID, limit, offset)
		if err != nil {
			return nil, 0, err
		}
		return clients, total, nil
	}
*/
func (s *clientService) FindAll(limit, offset int) ([]*Client, int64, error) {
	clients, count, err := s.repository.FindAll(limit, offset)
	if err != nil {
		return nil, 0, err
	}
	return clients, count, nil
}

func (s *clientService) FindById(id uint) (*Client, error) {
	client, err := s.repository.FindById(id)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func (s *clientService) FindByEmail(email string) (*Client, error) {
	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *clientService) FindByName(name string) ([]*Client, error) {
	clients, err := s.repository.FindByName(name)
	if err != nil {
		return nil, err
	}
	return clients, nil
}

func (s *clientService) Save(client Client) (*Client, error) {
	newclient, err := s.repository.Save(client)
	if err != nil {
		return nil, err
	}
	return newclient, nil
}

func (s *clientService) Remove(id uint) error {
	return s.repository.Remove(id)
}
