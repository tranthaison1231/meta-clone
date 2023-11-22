package services

import (
	"github.com/tranthaison1231/meta-clone/api/db"
	"github.com/tranthaison1231/meta-clone/api/models"
)

func GetChats(userID uint) (*[]models.Chat, error) {
	var chats []models.Chat
	err := db.DB.Model(&models.Chat{}).Preload("LastMessage").Preload("Members").Where("owner_id = ?", userID).Find(&chats).Error

	if err != nil {
		return nil, err
	}
	return &chats, nil
}

func CreateChat(newChat models.Chat) (*models.Chat, error) {
	err := db.DB.Create(&newChat).Error

	if err != nil {
		return nil, err
	}
	return &newChat, nil
}

func AddMemberToChat(chatID uint64, memberID uint) (*models.Chat, error) {
	user, err := GetUserByID(memberID)

	if err != nil {
		return nil, err
	}

	var chat models.Chat
	err = db.DB.Model(&chat).Preload("Members").Where("id = ?", chatID).First(&chat).Error

	if err != nil {
		return nil, err
	}

	chat.Members = append(chat.Members, user)

	err = db.DB.Save(&chat).Error

	if err != nil {
		return nil, err
	}
	return &chat, nil
}

func UpdateLastMessage(chatID uint64, message models.Message) error {
	var chat models.Chat
	err := db.DB.Model(&chat).Where("id = ?", chatID).First(&chat).Error

	chat.LastMessage = message

	err = db.DB.Save(&chat).Error

	return err
}
