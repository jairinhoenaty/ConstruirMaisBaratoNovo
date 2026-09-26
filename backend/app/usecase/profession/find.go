package profession_usecase

import pkgprofession "construir_mais_barato/app/domain/profession"

type FindProfessionUC struct {
	Service             pkgprofession.ProfessionService
	QuantityProfessions uint
}

type FindProfessionUCParams struct {
	Service pkgprofession.ProfessionService
}

func NewFindProfessionUC(params FindProfessionUCParams) FindProfessionUC {
	return FindProfessionUC{
		Service: params.Service,
	}
}

func (uc *FindProfessionUC) Execute() (*[]ProfessionPresenter, error) {

	professions, err := uc.Service.Find(uc.QuantityProfessions)
	if err != nil {
		return nil, err
	}
	presenters := GenerateProfessionPresenters(professions)
	return &presenters, nil
}
