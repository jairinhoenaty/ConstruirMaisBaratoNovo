package budget_usecase

import (
	pkgbudget "construir_mais_barato/app/domain/budget"
	pkgprofession "construir_mais_barato/app/domain/profession"
	pkgprofessional "construir_mais_barato/app/domain/professional"
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
	if budget != nil {

		professionalsPresenter := generateProfessionalPresenter(budget.Professionals)

		presenter.ID = budget.ID
		presenter.Name = budget.Name
		presenter.Email = budget.Email
		presenter.Telephone = budget.Telephone
		presenter.Description = budget.Description
		presenter.CreatedAt = budget.CreatedAt
		presenter.Professionals = professionalsPresenter
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

	}
	return presenter
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
