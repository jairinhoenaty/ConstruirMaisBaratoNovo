package region_usecase

import (
	pkgregion "construir_mais_barato/app/domain/region"
	"fmt"
)

type DeleteRegionUC struct {
	Service pkgregion.RegionService
	ID      *uint
}

type DeleteRegionUCParams struct {
	Service pkgregion.RegionService
}

func NewDeleteRegionUC(params DeleteRegionUCParams) DeleteRegionUC {
	return DeleteRegionUC{
		Service: params.Service,
	}
}

func (uc *DeleteRegionUC) Execute() error {
	if uc.ID == nil {
		return fmt.Errorf("invalid id")
	}

	err := uc.Service.Remove(*uc.ID)
	if err != nil {
		return err
	}
	return nil
}
