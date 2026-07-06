package pageview_repository_impl

import (
	"errors"

	"gorm.io/gorm"

	pkgpageview "construir_mais_barato/app/domain/pageview"
)

type repository struct {
	DB *gorm.DB
}

func NewPageViewRepositoryImpl(db *gorm.DB) pkgpageview.PageViewRepository {
	return &repository{
		DB: db,
	}
}

func (r *repository) Increment(path string) (*pkgpageview.PageView, error) {
	var pageView pkgpageview.PageView
	err := r.DB.Where("path = ?", path).First(&pageView).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		pageView = pkgpageview.PageView{
			Path:  path,
			Count: 1,
		}
		if createErr := r.DB.Create(&pageView).Error; createErr != nil {
			return nil, createErr
		}
		return &pageView, nil
	}
	if err != nil {
		return nil, err
	}
	pageView.Count++
	if saveErr := r.DB.Save(&pageView).Error; saveErr != nil {
		return nil, saveErr
	}
	return &pageView, nil
}

func (r *repository) FindAll() ([]*pkgpageview.PageView, error) {
	var pageViews []*pkgpageview.PageView
	if err := r.DB.Order("count desc").Find(&pageViews).Error; err != nil {
		return nil, err
	}
	return pageViews, nil
}
