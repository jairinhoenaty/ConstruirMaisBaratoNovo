package profession

import (
	"gorm.io/gorm"
)

type Profession struct {
	gorm.Model
	Name        string
	Description string
	Icon        string
}
