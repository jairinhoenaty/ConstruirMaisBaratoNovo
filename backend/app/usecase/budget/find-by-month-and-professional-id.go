package budget_usecase

import (
	pkgbudget "construir_mais_barato/app/domain/budget"
	pkgprofessional "construir_mais_barato/app/domain/professional"
	pkguser "construir_mais_barato/app/domain/user"
	pkgprofessionaluc "construir_mais_barato/app/usecase/professional"
	pkguseruc "construir_mais_barato/app/usecase/user"
	"fmt"
)

type FindByMonthAndProfessionalIDUC struct {
	Service             pkgbudget.BudgetService
	ServiceUser         pkguser.UserService
	ServiceProfessional pkgprofessional.ProfessionalService
	Assembler           *FindBudgetByMontAndProfessionalIDAssembler
}

type FindByMonthAndProfessionalIDUCParams struct {
	Service             pkgbudget.BudgetService
	ServiceUser         pkguser.UserService
	ServiceProfessional pkgprofessional.ProfessionalService
}

func NewFindByMonthAndProfessionalIDUC(params FindByMonthAndProfessionalIDUCParams) FindByMonthAndProfessionalIDUC {
	return FindByMonthAndProfessionalIDUC{
		Service:             params.Service,
		ServiceUser:         params.ServiceUser,
		ServiceProfessional: params.ServiceProfessional,
	}
}

func (uc *FindByMonthAndProfessionalIDUC) Execute() ([]BudgetPresenter, error) {
	fmt.Println("Executando Caso de uso de Orçamentos do profissional")

	//o id do profissional recebido é o id do profissional da tabela de usuário
	userParams := pkguseruc.FindByIdUCParamns{
		Service: uc.ServiceUser,
	}
	userUC := pkguseruc.NewFindByIdUC(userParams)

	if uint64(*uc.Assembler.ProfessionalID) != 0 {
		userUC.ID = uc.Assembler.ProfessionalID
	}
	if uc.Assembler.ClientID != 0 {
		uClientID := uint(uc.Assembler.ClientID)
		userUC.ID = &uClientID
	}
	user, err := userUC.Execute()
	if err != nil {
		return nil, fmt.Errorf("Usuário não encontrado com o id informado")
	}

	if uint64(*uc.Assembler.ProfessionalID) != 0 {
		//então devo pesquisar na tabela de profissional pelo nome profissional vinculado ao orçamento
		professionalParams := pkgprofessionaluc.FindByNamedUCParams{
			Service: uc.ServiceProfessional,
		}
		professionalUC := pkgprofessionaluc.NewFindByNamedUC(professionalParams)
		professionalUC.Assembler = &pkgprofessionaluc.FindByNameAssembler{
			Name: user.Name,
		}
		foundProfessionals, err := professionalUC.Execute()
		if err != nil {
			return nil, fmt.Errorf("Profissional não encontrado com o id informado")
		}
		// Verifica se a lista de profissionais não está vazia
		if len(*foundProfessionals) == 0 {
			return nil, fmt.Errorf("Nenhum profissional encontrado")
		}

		// Desreferencia o ponteiro para acessar a fatia real
		professionals := *foundProfessionals
		//fmt.Println(professionals)

		// Obtém o ID do primeiro profissional
		firstProfessionalID := professionals[0].ID

		// Faz algo com o ID (por exemplo, retorná-lo ou usá-lo em outra lógica)
		fmt.Printf("ID do primeiro profissional: %d\n", firstProfessionalID)
		uc.Assembler.ProfessionalID = &firstProfessionalID
	}

	listBudgets := make([]BudgetPresenter, 0)
	if uc.Assembler == nil {
		return nil, fmt.Errorf("invalid data")
	}

	fmt.Println("Parametros da consulta")
	fmt.Println("Mes => ", uc.Assembler.Month)
	fmt.Println("Id do profissional => ", int(*uc.Assembler.ProfessionalID))

	budgets, err := uc.Service.FindBudgetsByMonthAndProfessionalID(
		uc.Assembler.Month,
		uint(*uc.Assembler.ProfessionalID),
		uc.Assembler.ClientID,
		int(uc.Assembler.Page),
		int(uc.Assembler.PageSize),
	)

	if err != nil {
		return nil, err
	}

	if len(budgets) > 0 {
		for _, bg := range budgets {
			budgetPresenter := GenerateBudgetPresenter(bg)
			listBudgets = append(listBudgets, budgetPresenter)
		}
	}

	return listBudgets, nil
}
