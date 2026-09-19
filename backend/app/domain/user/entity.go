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
	// Síndico(a) ou representante de condomínio, marcado no cadastro pelo app.
	// Fica aqui, e não em Client, porque vale para qualquer perfil.
	IsCondoManager bool      `gorm:"default:false"`
	CreatedAt      time.Time `gorm:"<-:create"`
	//DeletedAt time.Time `gorm:"unique"` //`gorm:"index:idx_name,unique"`
}
