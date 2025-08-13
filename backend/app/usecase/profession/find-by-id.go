package profession_usecase

import (
	pkgprofession "construir_mais_barato/app/domain/profession"
	"fmt"
)

type FindByIdUC struct {
	Service pkgprofession.ProfessionService
	ID      *uint
}

type FindByIdUCParams struct {
	Service pkgprofession.ProfessionService
}

func NewFindByIdUC(params FindByIdUCParams) FindByIdUC {
	return FindByIdUC{
		Service: params.Service,
	}
}

func (uc *FindByIdUC) Execute() (*ProfessionPresenter, error) {
	if uc.ID == nil {
		return nil, fmt.Errorf("invalid id")
	}

	profission, err := uc.Service.FindById(*uc.ID)
	if err != nil {
		return nil, err
	}

	profissionPresenter := GenerateProfessionPresenter(profission)
	return &profissionPresenter, nil
}
