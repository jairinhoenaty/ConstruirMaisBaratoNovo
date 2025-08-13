package profession_usecase

import (
	pkgprofession "construir_mais_barato/app/domain/profession"
	"fmt"
)

type DeleteProfessionUC struct {
	Service pkgprofession.ProfessionService
	ID      *uint
}

type DeleteProfessionUCParams struct {
	Service pkgprofession.ProfessionService
}

func NewDeleteProfessionUC(params DeleteProfessionUCParams) DeleteProfessionUC {
	return DeleteProfessionUC{
		Service: params.Service,
	}
}

func (uc *DeleteProfessionUC) Execute() error {
	if uc.ID == nil {
		return fmt.Errorf("invalid id")
	}

	err := uc.Service.Remove(*uc.ID)
	if err != nil {
		return err
	}
	return nil
}
