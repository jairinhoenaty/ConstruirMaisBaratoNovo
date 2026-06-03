package budget_usecase

import (
	pkgbudget "construir_mais_barato/app/domain/budget"
	pkgprofessional "construir_mais_barato/app/domain/professional"
	pkgstore "construir_mais_barato/app/domain/store"
	pkguser "construir_mais_barato/app/domain/user"
	pkgnotification "construir_mais_barato/app/usecase/notification"
	"fmt"
)

type SaveBudgetUC struct {
	Service                     pkgbudget.BudgetService
	StoreService                pkgstore.StoreService
	ProfessionalService         pkgprofessional.ProfessionalService
	UserService                 pkguser.UserService
	SendAppNotificationUCParams pkgnotification.SendAppNotificationUCParams
	Assembler                   *BudgetAssembler
}

type SaveBudgetUCParams struct {
	Service                     pkgbudget.BudgetService
	StoreService                pkgstore.StoreService
	ProfessionalService         pkgprofessional.ProfessionalService
	UserService                 pkguser.UserService
	SendAppNotificationUCParams pkgnotification.SendAppNotificationUCParams
}

func NewSaveBudgetUC(params SaveBudgetUCParams) SaveBudgetUC {
	return SaveBudgetUC{
		Service:                     params.Service,
		StoreService:                params.StoreService,
		ProfessionalService:         params.ProfessionalService,
		UserService:                 params.UserService,
		SendAppNotificationUCParams: params.SendAppNotificationUCParams,
	}
}

func (uc *SaveBudgetUC) Execute() (*BudgetPresenter, error) {

	budget := GenerateBudget(uc.Assembler)
	budgetSaved, err := uc.Service.Save(budget)
	if err != nil {
		fmt.Println("Erro ao salvar orçamento", err.Error())
		return nil, err
	}
	budgetPresenter := GenerateBudgetPresenter(budgetSaved)

	if uc.Assembler.Approved && uc.Assembler.StoresId != nil && len(*uc.Assembler.StoresId) > 0 {
		uc.notifyStoreOwners(*uc.Assembler.StoresId)
	} else if uc.Assembler.Approved && uc.Assembler.ProfessionalsId != nil && len(*uc.Assembler.ProfessionalsId) > 0 {
		uc.notifyProfessionals(*uc.Assembler.ProfessionalsId)
	}

	return &budgetPresenter, nil
}

func (uc *SaveBudgetUC) notifyStoreOwners(storeIds []uint) {
	if uc.StoreService == nil || uc.UserService == nil {
		return
	}

	var userIds []uint
	for _, storeId := range storeIds {
		store, err := uc.StoreService.FindById(storeId)
		if err != nil {
			fmt.Printf("Erro ao buscar loja %d: %v\n", storeId, err)
			continue
		}

		user, err := uc.UserService.FindByEmail(store.Email)
		if err != nil {
			fmt.Printf("Usuário não encontrado para o email da loja %s: %v\n", store.Email, err)
			continue
		}

		userIds = append(userIds, user.ID)
	}

	if len(userIds) == 0 {
		return
	}

	notificationUC := pkgnotification.NewSendAppNotificationUC(uc.SendAppNotificationUCParams)
	notificationUC.Assembler = &pkgnotification.SendAppNotificationAssembler{
		Ids:    userIds,
		IDType: pkguser.IDTypeUser,
		Title:  "Chegou um novo orçamento!",
		Body:   "Confira o orçamento de produtos para a sua loja.",
	}

	if err := notificationUC.Execute(); err != nil {
		fmt.Printf("Erro ao enviar notificação para lojistas: %v\n", err)
	}
}

func (uc *SaveBudgetUC) notifyProfessionals(professionalIds []uint) {
	if uc.ProfessionalService == nil || uc.UserService == nil {
		return
	}

	var userIds []uint
	for _, professionalId := range professionalIds {
		professional, err := uc.ProfessionalService.FindById(professionalId)
		if err != nil {
			fmt.Printf("Erro ao buscar profissional %d: %v\n", professionalId, err)
			continue
		}

		user, err := uc.UserService.FindByEmail(professional.Email)
		if err != nil {
			fmt.Printf("Usuário não encontrado para o email do profissional %s: %v\n", professional.Email, err)
			continue
		}

		userIds = append(userIds, user.ID)
	}

	if len(userIds) == 0 {
		return
	}

	notificationUC := pkgnotification.NewSendAppNotificationUC(uc.SendAppNotificationUCParams)
	notificationUC.Assembler = &pkgnotification.SendAppNotificationAssembler{
		Ids:    userIds,
		IDType: pkguser.IDTypeUser,
		Title:  "Chegou um novo orçamento!",
		Body:   "Confira o orçamento de serviços disponível para você.",
	}

	if err := notificationUC.Execute(); err != nil {
		fmt.Printf("Erro ao enviar notificação para profissionais: %v\n", err)
	}
}
