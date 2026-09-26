package professionCategory_usecase

import pkgprofessionCategory "construir_mais_barato/app/domain/professionCategory"

func GenerateProfessionCategory(assembler *ProfessionCategoryAssembler) pkgprofessionCategory.ProfessionCategory {
	category := pkgprofessionCategory.ProfessionCategory{}
	if assembler == nil {
		return category
	}

	category.Name = assembler.Name
	category.Description = assembler.Description
	category.Icon = assembler.Icon
	category.Position = assembler.Position
	category.Active = assembler.Active
	category.ClientTravels = assembler.ClientTravels
	if assembler.ID > 0 {
		category.ID = assembler.ID
	}
	return category
}

func GenerateProfessionCategoryPresenter(category *pkgprofessionCategory.ProfessionCategory) ProfessionCategoryPresenter {
	presenter := ProfessionCategoryPresenter{}
	if category != nil {
		presenter.ID = category.ID
		presenter.Name = category.Name
		presenter.Description = category.Description
		presenter.Icon = category.Icon
		presenter.Position = category.Position
		presenter.Active = category.Active
		presenter.ClientTravels = category.ClientTravels
	}
	return presenter
}

// A contagem vai junto para a tela não precisar de uma chamada por card.
func GenerateProfessionCategoryPresenters(
	categories []*pkgprofessionCategory.ProfessionCategory,
	counts map[uint]int64,
) []ProfessionCategoryPresenter {
	presenters := make([]ProfessionCategoryPresenter, 0, len(categories))
	for _, category := range categories {
		presenter := GenerateProfessionCategoryPresenter(category)
		presenter.ProfessionCount = counts[category.ID]
		presenters = append(presenters, presenter)
	}
	return presenters
}

func CategoryIDs(categories []*pkgprofessionCategory.ProfessionCategory) []uint {
	ids := make([]uint, 0, len(categories))
	for _, category := range categories {
		ids = append(ids, category.ID)
	}
	return ids
}
