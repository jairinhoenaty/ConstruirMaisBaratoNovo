package banner_usecase

import (
	pkgbanner "construir_mais_barato/app/domain/banner"
	"fmt"
)

type DeleteBannerUC struct {
	Service pkgbanner.BannerService
	ID      *uint
}

type DeleteBannerUCParams struct {
	Service pkgbanner.BannerService
}

func NewDeleteBannerUC(params DeleteBannerUCParams) DeleteBannerUC {
	return DeleteBannerUC{
		Service: params.Service,
	}
}

func (uc *DeleteBannerUC) Execute() error {
	if uc.ID == nil {
		return fmt.Errorf("invalid id")
	}

	err := uc.Service.Remove(*uc.ID)
	if err != nil {
		return err
	}
	return nil
}
