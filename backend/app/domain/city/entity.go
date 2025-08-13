package city

import "gorm.io/gorm"

type City struct {
	gorm.Model
	Name string
	UF   string `gorm:"index"`
}
