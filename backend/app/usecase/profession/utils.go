package profession_usecase

import pkgprofession "construir_mais_barato/app/domain/profession"

func GenerateProfession(assembler *ProfessionAssembler) pkgprofession.Profession {
	profession := pkgprofession.Profession{}
	if assembler != nil {
		profession.Name = assembler.Name
		profession.Description = assembler.Description
		profession.Icon = assembler.Icon
		profession.CategoryID = assembler.CategoryID
		profession.ClientTravels = assembler.ClientTravels
	}
	if assembler.ID > 0 {
		profession.ID = assembler.ID
	}

	return profession
}

func GenerateProfessionPresenter(profession *pkgprofession.Profession) ProfessionPresenter {
	presenter := ProfessionPresenter{}
	if profession != nil {
		presenter.ID = profession.ID
		presenter.Name = profession.Name
		presenter.Description = profession.Description
		presenter.Icon = profession.Icon
		presenter.CategoryID = profession.CategoryID
		presenter.ClientTravels = profession.ClientTravelsResolved()
		presenter.ClientTravelsOverride = profession.ClientTravels
	}
	return presenter
}

func GenerateProfessionPresenters(professions []*pkgprofession.Profession) []ProfessionPresenter {
	presenters := make([]ProfessionPresenter, 0, len(professions))
	for _, profession := range professions {
		presenters = append(presenters, GenerateProfessionPresenter(profession))
	}
	return presenters
}
