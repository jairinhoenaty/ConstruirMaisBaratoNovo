package accountdeletion_usecase

import (
	pkgaccountdeletion "construir_mais_barato/app/domain/accountDeletion"
	"errors"
	"strings"
)

type CreateAccountDeletionUC struct {
	Service   pkgaccountdeletion.AccountDeletionService
	Assembler *CreateAccountDeletionAssembler
}

type CreateAccountDeletionUCParams struct {
	Service pkgaccountdeletion.AccountDeletionService
}

func NewCreateAccountDeletionUC(
	params CreateAccountDeletionUCParams,
) CreateAccountDeletionUC {
	return CreateAccountDeletionUC{
		Service: params.Service,
	}
}

func (uc *CreateAccountDeletionUC) Execute() (
	*pkgaccountdeletion.AccountDeletionRequest,
	error,
) {
	if uc.Assembler == nil {
		return nil, errors.New("dados da solicitação não informados")
	}

	name := strings.TrimSpace(uc.Assembler.Name)
	email := strings.TrimSpace(strings.ToLower(uc.Assembler.Email))

	if name == "" {
		return nil, errors.New("nome é obrigatório")
	}

	if email == "" {
		return nil, errors.New("e-mail é obrigatório")
	}

	request := pkgaccountdeletion.AccountDeletionRequest{
		UserID:    uc.Assembler.UserID,
		Name:      name,
		Email:     email,
		Telephone: strings.TrimSpace(uc.Assembler.Telephone),
		Reason:    strings.TrimSpace(uc.Assembler.Reason),
		Status:    pkgaccountdeletion.StatusReceived,
	}

	return uc.Service.Save(request)
}
