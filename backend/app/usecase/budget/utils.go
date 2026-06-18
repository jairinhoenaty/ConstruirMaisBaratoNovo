package budget_usecase

import (
	pkgbudget "construir_mais_barato/app/domain/budget"
	pkgprofession "construir_mais_barato/app/domain/profession"
	pkgprofessional "construir_mais_barato/app/domain/professional"
	pkgstore "construir_mais_barato/app/domain/store"
)

func GenerateBudget(assembler *BudgetAssembler) pkgbudget.Budget {
	budget := pkgbudget.Budget{}
	if assembler != nil {

		// professionals := generateProfessionals(assembler.Professionals)

		budget.ID = assembler.ID
		budget.Name = assembler.Name
		budget.Email = assembler.Email
		budget.Telephone = assembler.Telephone
		budget.Description = assembler.Description
		budget.ProfessionalIDs = assembler.ProfessionalsId
		budget.StoresIDs = assembler.StoresId
		budget.CityID = assembler.CityID
		budget.TermResponsabilityAccepted = assembler.TermResponsabilityAccepted
		// budget.ClientID = assembler.ClientID
		budget.Approved = assembler.Approved
	}
	return budget
}

// func generateProfessionals(professionals *[]ProfessionalAssembler) []pkgprofessional.Professional {
// 	list := make([]pkgprofessional.Professional, 0)
// 	if professionals != nil && len(*professionals) > 0 {
// 		for _, professional := range *professionals {

// 			professions := getProfessions(professional.Professions)

// 			prof := pkgprofessional.Professional{}

// 			prof.ID = professional.ID
// 			prof.Name = professional.Name
// 			prof.Email = professional.Email
// 			prof.Telephone = professional.Telephone
// 			prof.Professions = professions

// 			list = append(list, prof)
// 		}
// 	}

// 	return list
// }

// func getProfessions(professions *[]ProfessionAssembler) []pkgprofession.Profession {
// 	list := make([]pkgprofession.Profession, 0)
// 	if professions != nil && len(*professions) > 0 {

// 		for _, profession := range *professions {
// 			profess := pkgprofession.Profession{}
// 			profess.ID = profession.ID
// 			profess.Name = profession.Name
// 			list = append(list, profess)
// 		}
// 	}
// 	return list
// }

func GenerateBudgetPresenter(budget *pkgbudget.Budget) BudgetPresenter {
	presenter := BudgetPresenter{}
	var professionalsPresenter *[]ProfessionalPresenter
	var storePresenter *[]StorePresenter
	if budget != nil {

		// Analisar para no futuro não trazer dados de outros profissionais ou somente o profissional em questão, talvez criar uma função específica
		// para o presenter de orçamentos de profissional. Para o administrador, acho legal manter como está.
		if len(budget.Professionals) > 0 {
			professionalsPresenter = generateProfessionalPresenter(budget.Professionals)

		} else if len(budget.Stores) > 0 {
			storePresenter = generateStorePresenter(budget.Stores)

		}

		presenter.ID = budget.ID
		presenter.Name = budget.Name
		presenter.Email = budget.Email
		presenter.Telephone = budget.Telephone
		presenter.Description = budget.Description
		presenter.CreatedAt = budget.CreatedAt
		presenter.Professionals = professionalsPresenter
		presenter.Stores = storePresenter
		presenter.CityID = budget.CityID
		presenter.City = CityPresenter{
			Name: budget.City.Name,
			UF:   budget.City.UF,
		}
		presenter.TermResponsabilityAccepted = budget.TermResponsabilityAccepted
		// presenter.ClientID = budget.ClientID
		// presenter.Client = ClientPresenter{
		// 	Name:      budget.Client.Name,
		// 	Email:     budget.Client.Email,
		// 	Telephone: budget.Client.Telephone,
		// 	City: CityPresenter{
		// 		Name: budget.Client.City.Name,
		// 		UF:   budget.Client.City.UF,
		// 	},
		// }
		presenter.Approved = budget.Approved
		// Recusado só conta para profissional; lojista não marca recusado p/ admin.
		refusedCount := 0
		for _, refusal := range budget.Refusals {
			if refusal.RecipientType == pkgbudget.RecipientTypeProfessional {
				refusedCount++
			}
		}
		presenter.RefusedCount = refusedCount
		presenter.Refused = refusedCount > 0

	}
	return presenter
}

func generateStorePresenter(stores []pkgstore.Store) *[]StorePresenter {
	list := make([]StorePresenter, 0)
	if stores != nil && len(stores) > 0 {
		for _, store := range stores {

			presenter := StorePresenter{}

			presenter.ID = store.ID
			presenter.Name = store.Name
			presenter.Email = store.Email
			presenter.Telephone = store.Telephone
			presenter.City = CityPresenter{
				Name: store.City.Name,
				UF:   store.City.UF,
			}

			list = append(list, presenter)
		}
	}

	return &list
}
func generateProfessionalPresenter(professionals []pkgprofessional.Professional) *[]ProfessionalPresenter {
	list := make([]ProfessionalPresenter, 0)
	if professionals != nil && len(professionals) > 0 {
		for _, professional := range professionals {

			professions := getProfessionsPresenter(professional.Professions)

			presenter := ProfessionalPresenter{}

			presenter.ID = professional.ID
			presenter.Name = professional.Name
			presenter.Email = professional.Email
			presenter.Telephone = professional.Telephone
			presenter.Professions = professions
			presenter.City = CityPresenter{
				Name: professional.City.Name,
				UF:   professional.City.UF,
			}

			list = append(list, presenter)
		}
	}

	return &list
}

func getProfessionsPresenter(professions []pkgprofession.Profession) *[]ProfessionPresenter {
	list := make([]ProfessionPresenter, 0)
	if professions != nil && len(professions) > 0 {

		for _, profession := range professions {
			profess := ProfessionPresenter{}
			profess.ID = profession.ID
			profess.Name = profession.Name
			list = append(list, profess)
		}
	}
	return &list
}
