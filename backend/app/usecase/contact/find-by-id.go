package contact_usecase

import (
	pkgcontact "construir_mais_barato/app/domain/contact"
	"fmt"
)

type FindByIdUC struct {
	Service pkgcontact.ContactService
	ID      *uint
}

type FindByIdUCParams struct {
	Service pkgcontact.ContactService
}

func NewFindByIdUC(params FindByIdUCParams) FindByIdUC {
	return FindByIdUC{
		Service: params.Service,
	}
}

func (uc *FindByIdUC) Execute() (*ContactPresenter, error) {
	if uc.ID == nil {
		return nil, fmt.Errorf("invalid id")
	}

	contact, err := uc.Service.FindById(*uc.ID)
	if err != nil {
		return nil, err
	}

	contactPresenter := GenerateContactPresenter(contact)
	return &contactPresenter, nil
}
