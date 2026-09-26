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
	presenters := GenerateProfessionPresenters(professions)
	return &presenters, total, nil
}
