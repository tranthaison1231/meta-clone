package services

import (
	"github.com/tranthaison1231/messenger-clone/api/db"
	"github.com/tranthaison1231/messenger-clone/api/models"
)

func GetChats(userID uint) (*[]models.Chat, error) {
	var chats []models.Chat
	err := db.DB.Model(&models.Chat{}).Preload("Members").Where("owner_id = ?", userID).Find(&chats).Error

	if err != nil {
		return nil, err
	}
	return &chats, nil
}

func CreateChat(newChat models.Chat) (*models.Chat, error) {
	var chat models.Chat
	err := db.DB.Create(&newChat).Error

	if err != nil {
		return nil, err
	}
	return &chat, nil
}

func AddMemberToChat(userID uint, memberID uint) (*models.Chat, error) {
	var chat models.Chat
	err := db.DB.Model(&chat).Preload("Members").Where("id = ?", userID).First(&chat).Error

	if err != nil {
		return nil, err
	}
	return &chat, nil
}
