package contact_usecase

import (
	pkgcontact "construir_mais_barato/app/domain/contact"
	"fmt"
)

type DeleteContactUC struct {
	Service pkgcontact.ContactService
	ID      *uint
}

type DeleteContactUCParams struct {
	Service pkgcontact.ContactService
}

func NewDeleteContactUC(params DeleteContactUCParams) DeleteContactUC {
	return DeleteContactUC{
		Service: params.Service,
	}
}

func (uc *DeleteContactUC) Execute() error {
	if uc.ID == nil {
		return fmt.Errorf("invalid id")
	}

	err := uc.Service.Remove(*uc.ID)
	if err != nil {
		return err
	}
	return nil
}
