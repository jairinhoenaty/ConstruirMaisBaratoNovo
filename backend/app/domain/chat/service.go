package chat

type ChatService interface {
	FindAll(limit, offset int) ([]*Chat, int64, error)
	FindByUser(limit, offset int, professionalID int, clienteID int) ([]*Chat, int64, error)
	//FindAll(limit, offset int) ([]*Professional, int64, error)
	//FindById(id uint) (*Chat, error)
	Save(Chat Chat) (*Chat, error)
	Remove(id uint) error
}

type contactService struct {
	repository ChatRepository
}

func NewChatService(repository ChatRepository) ChatService {
	return &contactService{
		repository: repository,
	}
}

func (s *contactService) FindAll(limit, offset int) ([]*Chat, int64, error) {
	contacts, total, err := s.repository.FindAll(limit, offset)
	if err != nil {
		return nil, 0, err
	}
	return contacts, total, nil
}

func (s *contactService) FindByUser(limit, offset int, professionalID int, clienteID int) ([]*Chat, int64, error) {
	contacts, total, err := s.repository.FindByUser(limit, offset, professionalID, clienteID)
	if err != nil {
		return nil, 0, err
	}
	return contacts, total, nil
}

/*func (s *contactService) FindById(id uint) (*Chat, error) {
	contact, err := s.repository.FindById(id)
	if err != nil {
		return nil, err
	}
	return contact, nil
}*/

func (s *contactService) Save(Chat Chat) (*Chat, error) {
	newChat, err := s.repository.Save(Chat)
	if err != nil {
		return nil, err
	}
	return newChat, nil
}

func (s *contactService) Remove(id uint) error {
	return s.repository.Remove(id)
}
