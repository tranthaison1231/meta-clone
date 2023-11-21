package services

import (
	"github.com/tranthaison1231/messenger-clone/api/db"
	"github.com/tranthaison1231/messenger-clone/api/models"
)

func GetMessages(chatID int64) (*[]models.Message, error) {
	var messages []models.Message
	err := db.DB.Model(&models.Chat{}).Where("owner_id = ?", chatID).Find(&messages).Error

	if err != nil {
		return nil, err
	}
	return &messages, nil
}

func CreateMessage(newMessage models.Message) (*models.Message, error) {
	err := db.DB.Create(&newMessage).Error

	if err != nil {
		return nil, err
	}
	return &newMessage, nil
}
