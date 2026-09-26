package professionCategory_usecase

import (
	"fmt"

	pkgprofessionCategory "construir_mais_barato/app/domain/professionCategory"
)

type FindByIdUC struct {
	Service pkgprofessionCategory.ProfessionCategoryService
	ID      *uint
}

type FindByIdUCParams struct {
	Service pkgprofessionCategory.ProfessionCategoryService
}

func NewFindByIdUC(params FindByIdUCParams) FindByIdUC {
	return FindByIdUC{Service: params.Service}
}

func (uc *FindByIdUC) Execute() (*ProfessionCategoryPresenter, error) {
	if uc.ID == nil {
		return nil, fmt.Errorf("invalid id")
	}

	category, err := uc.Service.FindById(*uc.ID)
	if err != nil {
		return nil, err
	}

	presenter := GenerateProfessionCategoryPresenter(category)
	return &presenter, nil
}
