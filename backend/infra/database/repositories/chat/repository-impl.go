package chat_repository_impl

import (
	"fmt"

	"gorm.io/gorm"

	pkgchat "construir_mais_barato/app/domain/chat"
)

type repository struct {
	DB *gorm.DB
}

func NewChatRepositoryImpl(db *gorm.DB) pkgchat.ChatRepository {
	return &repository{
		DB: db,
	}
}

func (r *repository) FindAll(limit int, offset int) ([]*pkgchat.Chat, int64, error) {
	var chats []*pkgchat.Chat
	var total int64 = 0

	if err := r.DB.Model(&pkgchat.Chat{}).
		Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := r.DB.
		Preload("Professional").
		Preload("Professional.City").
		Preload("Client").
		Preload("Client.City").
		Order("created_at desc").
		Limit(limit).
		Offset(offset).Find(&chats).Error; err != nil {
		return nil, total, err
	}
	return chats, total, nil
}

func (r *repository) FindByUser(limit int, offset int, professionalID int, clientID int) ([]*pkgchat.Chat, int64, error) {
	fmt.Println("Busca Mensagens por User")
	var total int64 = 0

	var chats []*pkgchat.Chat

	if professionalID != 0 {
		if err := r.DB.Model(&pkgchat.Chat{}).
			Where("professional_id", professionalID).
			Count(&total).Error; err != nil {
			return nil, 0, err
		}

		if err := r.DB.
			Preload("Client").
			Preload("Client.City").
			Preload("Professional").
			Preload("Professional.City").
			Where("professional_id", professionalID).
			Limit(limit).
			Offset(offset).
			Order("created_at desc").Find(&chats).Error; err != nil {
			return nil, 0, err
		}
	}

	if clientID != 0 {
		if err := r.DB.Model(&pkgchat.Chat{}).
			Where("client_id", clientID).
			Count(&total).Error; err != nil {
			return nil, 0, err
		}
		if err := r.DB.
			Preload("Client").
			Preload("Client.City").
			Preload("Professional").
			Preload("Professional.City").
			Where("client_id", clientID).
			Order("created_at desc").
			Limit(limit).
			Offset(offset).
			Find(&chats).Error; err != nil {
			return nil, 0, err
		}
	}


	return chats, total, nil
}

func (r *repository) FindById(id uint) (*pkgchat.Chat, error) {
	chat := pkgchat.Chat{}
	if err := r.DB.First(&chat, id).Error; err != nil {
		return nil, err
	}
	return &chat, nil
}

func (r *repository) Save(chat pkgchat.Chat) (*pkgchat.Chat, error) {
	//fmt.Println(chat)
	if err := r.DB.Save(&chat).Error; err != nil {
		fmt.Println("Error saving chat:", err)
		return nil, err
	}
	return &chat, nil
}

func (r *repository) Remove(id uint) error {

	if err := r.DB.Delete(&pkgchat.Chat{}, id).Error; err != nil {
		return err
	}
	return nil
}
