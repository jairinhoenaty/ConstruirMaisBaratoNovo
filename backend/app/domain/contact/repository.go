package contact

type ContactRepository interface {
	FindAll(limit int, offset int) ([]*Contact, int64, error)
	FindByUser(limit, offset int, professionalID int, clienteID int, storeID int) ([]*Contact, int64, error)
	FindById(id uint) (*Contact, error)
	Save(Contact Contact) (*Contact, error)
	Remove(id uint) error
}
