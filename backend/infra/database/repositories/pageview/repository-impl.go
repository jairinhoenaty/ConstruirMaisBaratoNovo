package pageview_repository_impl

import (
	"errors"
	"time"

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

func getTodayDate() time.Time {
	now := time.Now()
	return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
}

func (r *repository) Increment(path string) (*pkgpageview.PageView, error) {
	viewDate := getTodayDate()
	var pageView pkgpageview.PageView
	err := r.DB.Where("path = ? AND view_date = ?", path, viewDate).First(&pageView).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		pageView = pkgpageview.PageView{
			Path:     path,
			ViewDate: viewDate,
			Count:    1,
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
	if err := r.DB.Order("view_date desc, count desc").Find(&pageViews).Error; err != nil {
		return nil, err
	}
	return pageViews, nil
}
