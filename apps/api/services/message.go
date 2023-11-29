package services

import (
	"github.com/tranthaison1231/meta-clone/api/db"
	"github.com/tranthaison1231/meta-clone/api/models"
)

func CreateMessage(newMessage models.Message) (*models.Message, error) {
	err := db.DB.Create(&newMessage).Error

	if err != nil {
		return nil, err
	}
	return &newMessage, nil
}
