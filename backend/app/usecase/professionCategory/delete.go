package professionCategory_usecase

import (
	"fmt"

	pkgprofession "construir_mais_barato/app/domain/profession"
	pkgprofessionCategory "construir_mais_barato/app/domain/professionCategory"
)

type DeleteUC struct {
	Service           pkgprofessionCategory.ProfessionCategoryService
	ProfessionService pkgprofession.ProfessionService
	ID                *uint
}

type DeleteUCParams struct {
	Service           pkgprofessionCategory.ProfessionCategoryService
	ProfessionService pkgprofession.ProfessionService
}

func NewDeleteUC(params DeleteUCParams) DeleteUC {
	return DeleteUC{
		Service:           params.Service,
		ProfessionService: params.ProfessionService,
	}
}

func (uc *DeleteUC) Execute() error {
	if uc.ID == nil {
		return fmt.Errorf("invalid id")
	}

	// Sem esta trava a exclusão deixaria profissões órfãs, e elas sumiriam da
	// busca do app — que só enxerga profissão através de uma categoria.
	professions, err := uc.ProfessionService.FindByCategory(*uc.ID)
	if err != nil {
		return err
	}
	if len(professions) > 0 {
		return fmt.Errorf("categoria possui %d profissão(ões) vinculada(s)", len(professions))
	}

	return uc.Service.Remove(*uc.ID)
}
