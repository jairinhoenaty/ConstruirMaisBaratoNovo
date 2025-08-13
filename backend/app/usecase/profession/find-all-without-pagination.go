package profession_usecase

import pkgprofession "construir_mais_barato/app/domain/profession"

type FindAllWithoutPaginationProfessionUC struct {
	Service pkgprofession.ProfessionService
}

type FindAllWithoutPaginationProfessionParams struct {
	Service pkgprofession.ProfessionService
}

func NewFindAllWithoutPaginationProfessionUC(params FindAllWithoutPaginationProfessionParams) FindAllWithoutPaginationProfessionUC {
	return FindAllWithoutPaginationProfessionUC{
		Service: params.Service,
	}
}

func (uc *FindAllWithoutPaginationProfessionUC) Execute() (*[]ProfessionPresenter, error) {

	professions, err := uc.Service.FindAllWithoutPagination()
	if err != nil {
		return nil, err
	}
	presenters := make([]ProfessionPresenter, 0)
	if len(professions) > 0 {
		for _, profession := range professions {
			presenters = append(presenters, ProfessionPresenter{
				ID:          profession.ID,
				Name:        profession.Name,
				Description: profession.Description,
				Icon:        profession.Icon,
			})
		}
	}
	return &presenters, nil
}
