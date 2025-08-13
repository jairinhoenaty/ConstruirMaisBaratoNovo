package chat

type ChatRepository interface {
	FindAll(limit int, offset int) ([]*Chat, int64, error)
	FindByUser(limit, offset int, professionalID int, clienteID int) ([]*Chat, int64, error)
	//FindById(id uint) (*Chat, error)
	Save(Chat Chat) (*Chat, error)
	Remove(id uint) error
}
