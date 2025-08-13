package contact_usecase

import (
	pkgcontact "construir_mais_barato/app/domain/contact"
)

type FindByUserContactUC struct {
	Service pkgcontact.ContactService
}

type FindByUserContactUCParams struct {
	Service pkgcontact.ContactService
}

func NewFindByUserUC(params FindByUserContactUCParams) FindByUserContactUC {
	return FindByUserContactUC{
		Service: params.Service,
	}
}

func (uc *FindByUserContactUC) Execute(limit int, offset int, professionalID int, clienteID int, storeID int) (*[]ContactPresenter, int64, error) {

	contacts, total, err := uc.Service.FindByUser(limit, offset, professionalID, clienteID, storeID)

	if err != nil {
		return nil, 0, err
	}
	presenters := make([]ContactPresenter, 0)
	if len(contacts) > 0 {
		for _, contact := range contacts {
			presenters = append(presenters, ContactPresenter{
				ID:             contact.ID,
				Name:           contact.Name,
				Telephone:      contact.Telephone,
				Email:          contact.Email,
				Message:        contact.Message,
				Status:         contact.Status,
				CityID:         contact.CityID,
				City:           contact.City,
				ProfessionalID: contact.ProfessionalID,
				Professional:   contact.Professional,
				ClientID:       contact.ClientID,
				Client:         contact.Client,
				StoreID:        contact.StoreID,
				Store:          contact.Store,
				ProductID:      contact.ProductID,
				Product:        contact.Product,
				CreatedAt:      contact.CreatedAt,
				Approved:       contact.Approved,
			})
		}
	}
	return &presenters, total, nil
}
