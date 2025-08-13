package user

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name        string
	Email       string `gorm:"unique"`
	Password    string
	Profile     string
	GoogleToken string
	CreatedAt   time.Time `gorm:"<-:create"`
	//DeletedAt time.Time `gorm:"unique"` //`gorm:"index:idx_name,unique"`
}
