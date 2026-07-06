package pageview

import "gorm.io/gorm"

type PageView struct {
	gorm.Model
	Path  string `gorm:"uniqueIndex;size:255;not null"`
	Count int64  `gorm:"default:0"`
}
