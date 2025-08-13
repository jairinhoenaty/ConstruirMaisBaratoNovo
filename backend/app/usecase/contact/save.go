package contact_usecase

import pkgcontact "construir_mais_barato/app/domain/contact"

type SaveContactUC struct {
	Service   pkgcontact.ContactService
	Assembler *ContactAssembler
}

type SaveContactUCParams struct {
	Service pkgcontact.ContactService
}

func NewSaveContactUC(params SaveContactUCParams) SaveContactUC {
	return SaveContactUC{
		Service: params.Service,
	}
}

func (uc *SaveContactUC) Execute() (*ContactPresenter, error) {
	contact := GenerateContact(uc.Assembler)
	contactSaved, err := uc.Service.Save(contact)
	if err != nil {
		return nil, err
	}
	contactPresenter := GenerateContactPresenter(contactSaved)

	return &contactPresenter, nil

}
