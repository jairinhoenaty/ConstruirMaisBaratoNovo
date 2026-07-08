package pageview_usecase

import (
	pkgpageview "construir_mais_barato/app/domain/pageview"
)

type PageViewPresenter struct {
	ID    uint   `json:"id"`
	Path  string `json:"path"`
	Count int64  `json:"count"`
}

func GeneratePageViewPresenter(pageView *pkgpageview.PageView) PageViewPresenter {
	if pageView == nil {
		return PageViewPresenter{}
	}
	return PageViewPresenter{
		ID:    pageView.ID,
		Path:  pageView.Path,
		Count: pageView.Count,
	}
}
