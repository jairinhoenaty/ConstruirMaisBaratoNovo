package professional_usecase

import (
	pkgprofessional "construir_mais_barato/app/domain/professional"
	pkguser "construir_mais_barato/app/domain/user"
	pkguseruc "construir_mais_barato/app/usecase/user"
	"fmt"
)

type FindByIdUC struct {
	Service     pkgprofessional.ProfessionalService
	ServiceUser pkguser.UserService
	ID          *uint
}

type FindByIdUCParamns struct {
	Service     pkgprofessional.ProfessionalService
	ServiceUser pkguser.UserService
}

func NewFindByIdUC(params FindByIdUCParamns) FindByIdUC {
	return FindByIdUC{
		Service:     params.Service,
		ServiceUser: params.ServiceUser,
	}
}

func (uc *FindByIdUC) Execute() (*ProfessionalPresenter, error) {
	if uc.ID == nil {
		return nil, fmt.Errorf("invalid id")
	}

	//o id do profissional recebido é o id do profissional da tabela de usuário
	userParams := pkguseruc.FindByIdUCParamns{
		Service: uc.ServiceUser,
	}
	userUC := pkguseruc.NewFindByIdUC(userParams)
	userUC.ID = uc.ID
	user, err := userUC.Execute()
	if err != nil {
		return nil, fmt.Errorf("Usuário não encontrado com o id informado")
	}

	//pesquisar pelo nome do usuário na tabela de profissionais
	foundProfessionals, err := uc.Service.FindByEmail(user.Email)

	if err != nil {
		println("Searching for professional by email error: ", err.Error())
		return nil, err
	}

	// Verifica se a lista de profissionais não está vazia
	if foundProfessionals == nil {
		return nil, fmt.Errorf("Nenhum profissional encontrado")
	}

	// Obtém o ID do primeiro profissional
	professional := foundProfessionals

	professionalPresenter := GenerateProfessionalPresenter(professional)

	return &professionalPresenter, nil
}
