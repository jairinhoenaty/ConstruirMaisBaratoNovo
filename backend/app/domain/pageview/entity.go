package pageview

import (
	"time"

	"gorm.io/gorm"
)

type PageView struct {
	gorm.Model
	Path     string    `gorm:"uniqueIndex:idx_pageview_path_date;size:255;not null"`
	ViewDate time.Time `gorm:"uniqueIndex:idx_pageview_path_date;type:date;not null"`
	Count    int64     `gorm:"default:0"`
}
