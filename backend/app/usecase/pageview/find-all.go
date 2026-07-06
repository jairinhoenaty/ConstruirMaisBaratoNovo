package pageview_usecase

import (
	pkgpageview "construir_mais_barato/app/domain/pageview"
)

type FindAllPageViewsUCParams struct {
	Service pkgpageview.PageViewService
}

type FindAllPageViewsUC struct {
	Service pkgpageview.PageViewService
}

func NewFindAllPageViewsUC(params FindAllPageViewsUCParams) FindAllPageViewsUC {
	return FindAllPageViewsUC{
		Service: params.Service,
	}
}

func (uc *FindAllPageViewsUC) Execute() ([]PageViewPresenter, error) {
	pageViews, err := uc.Service.FindAll()
	if err != nil {
		return nil, err
	}
	presenters := make([]PageViewPresenter, 0, len(pageViews))
	for _, pageView := range pageViews {
		presenters = append(presenters, GeneratePageViewPresenter(pageView))
	}
	return presenters, nil
}
