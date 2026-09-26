package professionCategory_usecase

import pkgprofessionCategory "construir_mais_barato/app/domain/professionCategory"

type FindAllUC struct {
	Service   pkgprofessionCategory.ProfessionCategoryService
	Assembler FindWithPaginationAssembler
}

type FindAllUCParams struct {
	Service pkgprofessionCategory.ProfessionCategoryService
}

func NewFindAllUC(params FindAllUCParams) FindAllUC {
	return FindAllUC{Service: params.Service}
}

func (uc *FindAllUC) Execute() ([]ProfessionCategoryPresenter, int64, error) {
	categories, total, err := uc.Service.FindAll(uc.Assembler.Limit, uc.Assembler.Offset)
	if err != nil {
		return nil, 0, err
	}

	counts, err := uc.Service.CountProfessions(CategoryIDs(categories))
	if err != nil {
		return nil, 0, err
	}

	return GenerateProfessionCategoryPresenters(categories, counts), total, nil
}
