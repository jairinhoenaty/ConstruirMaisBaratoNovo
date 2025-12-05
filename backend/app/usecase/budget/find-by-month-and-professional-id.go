package budget_usecase

import (
	pkgbudget "construir_mais_barato/app/domain/budget"
	pkgprofessional "construir_mais_barato/app/domain/professional"
	pkgstore "construir_mais_barato/app/domain/store"
	pkguser "construir_mais_barato/app/domain/user"
	pkgprofessionaluc "construir_mais_barato/app/usecase/professional"
	pkgstoreuc "construir_mais_barato/app/usecase/store"
	pkguseruc "construir_mais_barato/app/usecase/user"
	"fmt"
)

type FindByMonthAndProfessionalIDUC struct {
	Service             pkgbudget.BudgetService
	ServiceUser         pkguser.UserService
	ServiceProfessional pkgprofessional.ProfessionalService
	ServiceStore        pkgstore.StoreService
	Assembler           *FindBudgetByMontAndProfessionalIDAssembler
}

type FindByMonthAndProfessionalIDUCParams struct {
	Service             pkgbudget.BudgetService
	ServiceUser         pkguser.UserService
	ServiceStore        pkgstore.StoreService
	ServiceProfessional pkgprofessional.ProfessionalService
}

func NewFindByMonthAndProfessionalIDUC(params FindByMonthAndProfessionalIDUCParams) FindByMonthAndProfessionalIDUC {
	return FindByMonthAndProfessionalIDUC{
		Service:             params.Service,
		ServiceUser:         params.ServiceUser,
		ServiceProfessional: params.ServiceProfessional,
		ServiceStore:        params.ServiceStore,
	}
}

func (uc *FindByMonthAndProfessionalIDUC) Execute() ([]BudgetPresenter, error) {
	fmt.Println("Executando Caso de uso de Orçamentos do profissional")
	storeID := uint(0)
	profissionalID := uint(0)

	if uc.Assembler.ProfessionalID == nil {
		uc.Assembler.ProfessionalID = new(uint)
		*uc.Assembler.ProfessionalID = 0
	}

	//o id do profissional recebido é o id do profissional da tabela de usuário
	userParams := pkguseruc.FindByIdUCParamns{
		Service: uc.ServiceUser,
	}
	userUC := pkguseruc.NewFindByIdUC(userParams)

	if uc.Assembler.ProfessionalID != nil && uint64(*uc.Assembler.ProfessionalID) != 0 {
		userUC.ID = uc.Assembler.ProfessionalID
	} else if uc.Assembler.StoreID != nil && uint64(*uc.Assembler.StoreID) != 0 {
		userUC.ID = uc.Assembler.StoreID
	}
	if uc.Assembler.ClientID != 0 {
		uClientID := uint(uc.Assembler.ClientID)
		userUC.ID = &uClientID
	}
	user, err := userUC.Execute()
	if err != nil {
		return nil, fmt.Errorf("Usuário não encontrado com o id informado")
	}

	err = uc.GetProfessionalOrStore(user)
	if err != nil {
		return nil, err
	}

	listBudgets := make([]BudgetPresenter, 0)
	if uc.Assembler == nil {
		return nil, fmt.Errorf("invalid data")
	}

	fmt.Println("Parametros da consulta")
	fmt.Println("Mes => ", uc.Assembler.Month)
	if uc.Assembler.ProfessionalID != nil && *uc.Assembler.ProfessionalID != 0 {
		fmt.Println("Id do profissional => ", int(*uc.Assembler.ProfessionalID))
		profissionalID = *uc.Assembler.ProfessionalID
	} else if uc.Assembler.StoreID != nil && *uc.Assembler.StoreID != 0 {
		storeID = *uc.Assembler.StoreID
	} else {
		return nil, fmt.Errorf("erro ao buscar orçamentos")
	}

	budgets, err := uc.Service.FindBudgetsByMonthAndProfessionalID(
		uc.Assembler.Month,
		profissionalID,
		storeID,
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

func (uc *FindByMonthAndProfessionalIDUC) GetProfessionalOrStore(user *pkguseruc.UserPresenter) error {
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
			return fmt.Errorf("Profissional não encontrado com o id informado")
		}
		// Verifica se a lista de profissionais não está vazia
		if len(*foundProfessionals) == 0 {
			return fmt.Errorf("Nenhum profissional encontrado")
		}

		// Desreferencia o ponteiro para acessar a fatia real
		professionals := *foundProfessionals
		//fmt.Println(professionals)

		// Obtém o ID do primeiro profissional
		firstProfessionalID := professionals[0].ID

		// Faz algo com o ID (por exemplo, retorná-lo ou usá-lo em outra lógica)
		fmt.Printf("ID do primeiro profissional: %d\n", firstProfessionalID)
		uc.Assembler.ProfessionalID = &firstProfessionalID
	} else if uc.Assembler.StoreID != nil && uint64(*uc.Assembler.StoreID) != 0 {

		storeParams := pkgstoreuc.FindByNamedUCParams{
			Service: uc.ServiceStore,
		}
		storeUC := pkgstoreuc.NewFindByNamedUC(storeParams)
		storeUC.Assembler = &pkgstoreuc.FindByNameAssembler{
			Name: user.Name,
		}
		foundStore, err := storeUC.Execute()
		if err != nil {
			return fmt.Errorf("Lojista não encontrado com o id informado")
		}
		Stores := *foundStore
		firstStore := Stores[0].ID
		uc.Assembler.StoreID = &firstStore
	}
	return nil
}
