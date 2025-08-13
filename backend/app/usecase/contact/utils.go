package contact_usecase

import (
	pkgcontact "construir_mais_barato/app/domain/contact"
)

func GenerateContact(assembler *ContactAssembler) pkgcontact.Contact {
	contact := pkgcontact.Contact{}
	if assembler != nil {
		contact.ID = assembler.ID
		contact.Name = assembler.Name
		contact.Telephone = assembler.Telephone
		contact.Email = assembler.Email
		contact.Message = assembler.Mensagem
		contact.Status = assembler.Status
		contact.CityID = assembler.CityID
		contact.ProfessionalID = assembler.ProfessionalID
		contact.ClientID = assembler.ClientID
		contact.StoreID = assembler.StoreID
		contact.ProductID = assembler.ProductID
		contact.Approved = assembler.Approved

	}
	return contact
}

func GenerateContactPresenter(contact *pkgcontact.Contact) ContactPresenter {
	presenter := ContactPresenter{}
	if contact != nil {
		presenter.ID = contact.ID
		presenter.Name = contact.Name
		presenter.Telephone = contact.Telephone
		presenter.Email = contact.Email
		presenter.Message = contact.Message
		presenter.Status = contact.Status
		presenter.CityID = contact.CityID
		presenter.City = contact.City
		presenter.ProfessionalID = contact.ProfessionalID
		presenter.Professional = contact.Professional
		presenter.ClientID = contact.ClientID
		presenter.Client = contact.Client
		presenter.StoreID = contact.StoreID
		presenter.Store = contact.Store
		presenter.ProductID = contact.ProductID
		presenter.Product = contact.Product
		presenter.CreatedAt = contact.CreatedAt
		presenter.Approved = contact.Approved

	}

	return presenter
}
