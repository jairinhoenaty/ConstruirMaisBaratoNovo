package profession_usecase

import pkgprofession "construir_mais_barato/app/domain/profession"

type FindAllProfessionUC struct {
	Service   pkgprofession.ProfessionService
	Assembler FindWithPaginationProfessionAssembler
}

type FindAllProfessionUCParams struct {
	Service pkgprofession.ProfessionService
}

func NewFindAllProfessionUC(params FindAllProfessionUCParams) FindAllProfessionUC {
	return FindAllProfessionUC{
		Service: params.Service,
	}
}

func (uc *FindAllProfessionUC) Execute() (*[]ProfessionPresenter, int64, error) {

	professions, total, err := uc.Service.FindAll(uc.Assembler.Limit, uc.Assembler.Offset)
	if err != nil {
		return nil, 0, err
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
	return &presenters, total, nil
}
