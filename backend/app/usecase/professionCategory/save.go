package professionCategory_usecase

import pkgprofessionCategory "construir_mais_barato/app/domain/professionCategory"

type SaveUC struct {
	Service   pkgprofessionCategory.ProfessionCategoryService
	Assembler *ProfessionCategoryAssembler
}

type SaveUCParams struct {
	Service pkgprofessionCategory.ProfessionCategoryService
}

func NewSaveUC(params SaveUCParams) SaveUC {
	return SaveUC{Service: params.Service}
}

func (uc *SaveUC) Execute() (*ProfessionCategoryPresenter, error) {
	category := GenerateProfessionCategory(uc.Assembler)

	saved, err := uc.Service.Save(category)
	if err != nil {
		return nil, err
	}

	presenter := GenerateProfessionCategoryPresenter(saved)
	return &presenter, nil
}
