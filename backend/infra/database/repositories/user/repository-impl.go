package user_repository_impl

import (
	"gorm.io/gorm"

	pkguser "construir_mais_barato/app/domain/user"
)

type repository struct {
	DB *gorm.DB
}

func NewUserRepositoryImpl(db *gorm.DB) pkguser.UserRepository {
	return &repository{
		DB: db,
	}
}

func (r *repository) FindAll() ([]*pkguser.User, error) {
	var users []*pkguser.User
	if err := r.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *repository) FindById(id uint) (*pkguser.User, error) {
	user := pkguser.User{}
	if err := r.DB.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *repository) FindByEmail(email string) (*pkguser.User, error) {
	user := pkguser.User{}
	if err := r.DB.Where("email = ? ", email).
		Unscoped().
		First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil

}

func (r *repository) Save(user pkguser.User) (*pkguser.User, error) {
	if err := r.DB.Save(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *repository) Remove(id uint) error {

	if err := r.DB.Delete(&pkguser.User{}, id).Error; err != nil {
		return err
	}
	return nil
}
