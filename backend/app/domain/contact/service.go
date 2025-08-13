package contact

type ContactService interface {
	FindAll(limit, offset int) ([]*Contact, int64, error)
	FindByUser(limit, offset int, professionalID int, clienteID int, storeID int) ([]*Contact, int64, error)
	//FindAll(limit, offset int) ([]*Professional, int64, error)
	FindById(id uint) (*Contact, error)
	Save(Contact Contact) (*Contact, error)
	Remove(id uint) error
}

type contactService struct {
	repository ContactRepository
}

func NewContactService(repository ContactRepository) ContactService {
	return &contactService{
		repository: repository,
	}
}

func (s *contactService) FindAll(limit, offset int) ([]*Contact, int64, error) {
	contacts, total, err := s.repository.FindAll(limit, offset)
	if err != nil {
		return nil, 0, err
	}
	return contacts, total, nil
}

func (s *contactService) FindByUser(limit, offset int, professionalID int, clienteID int, storeID int) ([]*Contact, int64, error) {
	contacts, total, err := s.repository.FindByUser(limit, offset, professionalID, clienteID, storeID)
	if err != nil {
		return nil, 0, err
	}
	return contacts, total, nil
}

func (s *contactService) FindById(id uint) (*Contact, error) {
	contact, err := s.repository.FindById(id)
	if err != nil {
		return nil, err
	}
	return contact, nil
}

func (s *contactService) Save(Contact Contact) (*Contact, error) {
	newContact, err := s.repository.Save(Contact)
	if err != nil {
		return nil, err
	}
	return newContact, nil
}

func (s *contactService) Remove(id uint) error {
	return s.repository.Remove(id)
}
