package profession_usecase

import (
	pkgprofession "construir_mais_barato/app/domain/profession"
)

type FindProfessionsWithCountIdUC struct {
	Service pkgprofession.ProfessionService
}

type FindProfessionsWithCountIdUCParams struct {
	Service pkgprofession.ProfessionService
}

func NewFindProfessionsWithCountIdUC(params FindProfessionsWithCountIdUCParams) FindProfessionsWithCountIdUC {
	return FindProfessionsWithCountIdUC{
		Service: params.Service,
	}
}

func (uc *FindProfessionsWithCountIdUC) Execute() (*[]ProfessionWithCountPresenter, error) {

	result, err := uc.Service.FindProfessionsWithCount()
	if err != nil {
		return nil, err
	}

	presenters := make([]ProfessionWithCountPresenter, 0)
	for _, res := range result {

		presenter := ProfessionWithCountPresenter{
			Name:  res["profession"].(string),
			Count: int(res["count"].(int64)),
		}

		presenters = append(presenters, presenter)
	}

	return &presenters, nil
}
