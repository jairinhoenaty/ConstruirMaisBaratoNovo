package profession_usecase

import pkgprofession "construir_mais_barato/app/domain/profession"

func GenerateProfession(assembler *ProfessionAssembler) pkgprofession.Profession {
	profession := pkgprofession.Profession{}
	if assembler != nil {
		profession.Name = assembler.Name
		profession.Description = assembler.Description
		profession.Icon = assembler.Icon
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
	}
	return presenter
}
