package contact_repository_impl

import (
	"fmt"

	"gorm.io/gorm"

	pkgcontact "construir_mais_barato/app/domain/contact"
)

type repository struct {
	DB *gorm.DB
}

func NewContactRepositoryImpl(db *gorm.DB) pkgcontact.ContactRepository {
	return &repository{
		DB: db,
	}
}

func (r *repository) FindAll(limit int, offset int) ([]*pkgcontact.Contact, int64, error) {
	var contacts []*pkgcontact.Contact
	var total int64 = 0

	if err := r.DB.Model(&pkgcontact.Contact{}).
		Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := r.DB.Preload("Product").
		Preload("City").
		Preload("Professional").
		Preload("Professional.City").
		Preload("Client").
		Preload("Client.City").
		Preload("Store").
		Preload("Store.City").
		Order("created_at desc").
		Limit(limit).
		Offset(offset).Find(&contacts).Error; err != nil {
		return nil, total, err
	}
	return contacts, total, nil
}

func (r *repository) FindByUser(limit int, offset int, professionalID int, clientID int, storeID int) ([]*pkgcontact.Contact, int64, error) {
	fmt.Println("Busca Mensagens por User")
	var total int64 = 0

	var contacts []*pkgcontact.Contact

	if professionalID != 0 {
		if err := r.DB.Model(&pkgcontact.Contact{}).
			Where("professional_id", professionalID).
			Where("approved", true).
			Count(&total).Error; err != nil {
			return nil, 0, err
		}

		if err := r.DB.Preload("Product").
			Preload("City").
			Preload("Client").
			Preload("Client.City").
			Preload("Store").
			Preload("Store.City").
			Preload("Professional").
			Preload("Professional.City").
			Where("professional_id", professionalID).
			Where("approved", true).
			Limit(limit).
			Offset(offset).
			Order("created_at desc").Find(&contacts).Error; err != nil {
			return nil, 0, err
		}
	}

	if clientID != 0 {
		if err := r.DB.Model(&pkgcontact.Contact{}).
			Where("client_id", clientID).
			Where("approved", true).
			Count(&total).Error; err != nil {
			return nil, 0, err
		}
		if err := r.DB.Preload("Product").
			Preload("City").
			Preload("Client").
			Preload("Client.City").
			Preload("Store").
			Preload("Store.City").
			Preload("Professional").
			Preload("Professional.City").
			Where("client_id", clientID).
			Where("approved", true).
			Order("created_at desc").
			Limit(limit).
			Offset(offset).
			Find(&contacts).Error; err != nil {
			return nil, 0, err
		}
	}

	if storeID != 0 {
		if err := r.DB.Model(&pkgcontact.Contact{}).
			Where("store_id", storeID).
			Where("approved", true).
			Count(&total).Error; err != nil {
			return nil, 0, err
		}

		if err := r.DB.Preload("Product").Preload("City").
			Preload("Client").
			Preload("Client.City").
			Preload("Store").
			Preload("Store.City").
			Preload("Professional").
			Preload("Professional.City").
			Where("store_id", storeID).
			Where("approved", true).
			Limit(limit).
			Offset(offset).
			Order("created_at desc").Find(&contacts).Error; err != nil {
			return nil, 0, err
		}
	}

	return contacts, total, nil
}

func (r *repository) FindById(id uint) (*pkgcontact.Contact, error) {
	contact := pkgcontact.Contact{}
	if err := r.DB.First(&contact, id).Error; err != nil {
		return nil, err
	}
	return &contact, nil
}

func (r *repository) Save(contact pkgcontact.Contact) (*pkgcontact.Contact, error) {
	//fmt.Println(contact)
	if err := r.DB.Save(&contact).Error; err != nil {
		fmt.Println("Error saving contact:", err)
		return nil, err
	}
	return &contact, nil
}

func (r *repository) Remove(id uint) error {

	if err := r.DB.Delete(&pkgcontact.Contact{}, id).Error; err != nil {
		return err
	}
	return nil
}
