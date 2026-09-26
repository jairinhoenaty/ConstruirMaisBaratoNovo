package professionCategory_usecase

import pkgprofessionCategory "construir_mais_barato/app/domain/professionCategory"

type FindActiveUC struct {
	Service pkgprofessionCategory.ProfessionCategoryService
}

type FindActiveUCParams struct {
	Service pkgprofessionCategory.ProfessionCategoryService
}

func NewFindActiveUC(params FindActiveUCParams) FindActiveUC {
	return FindActiveUC{Service: params.Service}
}

func (uc *FindActiveUC) Execute() ([]ProfessionCategoryPresenter, error) {
	categories, err := uc.Service.FindActive()
	if err != nil {
		return nil, err
	}

	counts, err := uc.Service.CountProfessions(CategoryIDs(categories))
	if err != nil {
		return nil, err
	}

	return GenerateProfessionCategoryPresenters(categories, counts), nil
}
