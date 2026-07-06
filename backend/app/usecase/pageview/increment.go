package pageview_usecase

import (
	pkgpageview "construir_mais_barato/app/domain/pageview"
)

type IncrementPageViewUCParams struct {
	Service pkgpageview.PageViewService
}

type IncrementPageViewAssembler struct {
	Path string `json:"path"`
}

type IncrementPageViewUC struct {
	Service   pkgpageview.PageViewService
	Assembler *IncrementPageViewAssembler
}

func NewIncrementPageViewUC(params IncrementPageViewUCParams) IncrementPageViewUC {
	return IncrementPageViewUC{
		Service: params.Service,
	}
}

func (uc *IncrementPageViewUC) Execute() (*PageViewPresenter, error) {
	pageView, err := uc.Service.Increment(uc.Assembler.Path)
	if err != nil {
		return nil, err
	}
	presenter := GeneratePageViewPresenter(pageView)
	return &presenter, nil
}
