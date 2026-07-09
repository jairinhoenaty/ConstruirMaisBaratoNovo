package pageview_usecase

import (
	pkgpageview "construir_mais_barato/app/domain/pageview"
)

type PageViewPresenter struct {
	ID       uint   `json:"id"`
	Path     string `json:"path"`
	ViewDate string `json:"viewDate"`
	Count    int64  `json:"count"`
}

func GeneratePageViewPresenter(pageView *pkgpageview.PageView) PageViewPresenter {
	if pageView == nil {
		return PageViewPresenter{}
	}
	viewDate := ""
	if !pageView.ViewDate.IsZero() {
		viewDate = pageView.ViewDate.Format("2006-01-02")
	}
	return PageViewPresenter{
		ID:       pageView.ID,
		Path:     pageView.Path,
		ViewDate: viewDate,
		Count:    pageView.Count,
	}
}
