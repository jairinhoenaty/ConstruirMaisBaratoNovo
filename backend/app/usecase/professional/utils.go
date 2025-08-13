package professional_usecase

import (
	// pkgprofession "construir_mais_barato/app/domain/profession"
	pkgprofessional "construir_mais_barato/app/domain/professional"
)

func GenerateProfessional(assembler *ProfessionalAssembler) pkgprofessional.Professional {
	professional := pkgprofessional.Professional{}
	if assembler != nil {

		// professions := make([]pkgprofession.Profession, 0)
		// if len(assembler.Professions) > 0 {
		// 	for _, profession := range assembler.Professions {
		// 		professions = append(professions, pkgprofession.Profession{
		// 			Name:        profession.Name,
		// 			Description: profession.Description,
		// 			Icon:        profession.Icon,
		// 		})

		// 	}
		// }

		professional.ID = assembler.ID
		professional.Name = assembler.Name
		professional.Email = assembler.Email
		professional.Telephone = assembler.Telephone
		professional.LgpdAceito = assembler.LgpdAceito
		professional.CityID = assembler.CityID
		professional.ProfessionIDs = assembler.ProfessionIDs
		professional.Cep = assembler.Cep
		professional.Street = assembler.Street
		professional.Neighborhood = assembler.Neighborhood
		professional.Image = assembler.Image
		professional.Verified = assembler.Verified
		professional.OnLine = assembler.OnLine

	}
	return professional
}

func GenerateProfessionalPresenter(professional *pkgprofessional.Professional) ProfessionalPresenter {
	presenter := ProfessionalPresenter{}
	if professional != nil {

		professionPresenter := make([]ProfissionPresenter, 0)
		if len(professional.Professions) > 0 {
			for _, profession := range professional.Professions {
				professionPresenter = append(professionPresenter, ProfissionPresenter{
					ID:          profession.ID,
					Name:        profession.Name,
					Description: profession.Description,
					Icon:        profession.Icon,
				})

			}
		}

		cidadePresenter := CidadePresenter{
			ID:   professional.CityID,
			Name: professional.City.Name,
			UF:   professional.City.UF,
		}

		presenter.ID = professional.ID
		presenter.Name = professional.Name
		presenter.Email = professional.Email
		presenter.Telephone = professional.Telephone
		presenter.LgpdAceito = professional.LgpdAceito
		presenter.CreatedAt = professional.CreatedAt
		presenter.Cidade = cidadePresenter
		presenter.Professions = professionPresenter
		presenter.Cep = professional.Cep
		presenter.Street = professional.Street
		presenter.Neighborhood = professional.Neighborhood
		presenter.Image = professional.Image
		presenter.Distance = professional.Distance
		presenter.OnLine = professional.OnLine
		presenter.Verified = professional.Verified

	}
	return presenter
}
