package professionCategory_usecase

import (
	"fmt"

	pkgprofession "construir_mais_barato/app/domain/profession"
	pkgprofessionCategory "construir_mais_barato/app/domain/professionCategory"
	pkgprofessionuc "construir_mais_barato/app/usecase/profession"
)

// Devolve a categoria com as profissões dela e o client_travels já resolvido.
type FindProfessionsUC struct {
	Service           pkgprofessionCategory.ProfessionCategoryService
	ProfessionService pkgprofession.ProfessionService
	CategoryID        *uint
}

type FindProfessionsUCParams struct {
	Service           pkgprofessionCategory.ProfessionCategoryService
	ProfessionService pkgprofession.ProfessionService
}

func NewFindProfessionsUC(params FindProfessionsUCParams) FindProfessionsUC {
	return FindProfessionsUC{
		Service:           params.Service,
		ProfessionService: params.ProfessionService,
	}
}

func (uc *FindProfessionsUC) Execute() (*CategoryWithProfessionsPresenter, error) {
	if uc.CategoryID == nil {
		return nil, fmt.Errorf("invalid id")
	}

	category, err := uc.Service.FindById(*uc.CategoryID)
	if err != nil {
		return nil, err
	}

	professions, err := uc.ProfessionService.FindByCategory(category.ID)
	if err != nil {
		return nil, err
	}

	presenter := GenerateProfessionCategoryPresenter(category)
	presenter.ProfessionCount = int64(len(professions))

	return &CategoryWithProfessionsPresenter{
		ProfessionCategoryPresenter: presenter,
		Professions:                 pkgprofessionuc.GenerateProfessionPresenters(professions),
	}, nil
}
