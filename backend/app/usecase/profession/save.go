package profession_usecase

import pkgprofession "construir_mais_barato/app/domain/profession"

type SaveProfessionUC struct {
	Service   pkgprofession.ProfessionService
	Assembler *ProfessionAssembler
}

type SaveProfessionUCParams struct {
	Service pkgprofession.ProfessionService
}

func NewSaveProfessionUC(params SaveProfessionUCParams) SaveProfessionUC {
	return SaveProfessionUC{
		Service: params.Service,
	}
}

func (uc *SaveProfessionUC) Execute() (*ProfessionPresenter, error) {
	profission := GenerateProfession(uc.Assembler)
	profissionSaved, err := uc.Service.Save(profission)
	if err != nil {
		return nil, err
	}
	profissionPresenter := GenerateProfessionPresenter(profissionSaved)

	return &profissionPresenter, nil

}
