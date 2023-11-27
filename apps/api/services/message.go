package services

import (
	"github.com/pkg/errors"
	"github.com/tranthaison1231/meta-clone/api/db"
	h "github.com/tranthaison1231/meta-clone/api/helpers"
	"github.com/tranthaison1231/meta-clone/api/models"
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

func GetBatchMessages(messageId uint, isUp bool, request *models.BasePaginationRequest) (*models.BasePaginationResponse[models.Message], error) {
	var messages []models.Message
	var targetMessage models.Message

	targetMessageQuery := db.DB.Where("id = ?", messageId).First(&targetMessage)

	if err := targetMessageQuery.Error; err != nil {
		return nil, err
	}

	if targetMessageQuery.RowsAffected < 1 {
		return nil, errors.New("Target message not found")
	}

	query := db.DB.Model(&models.Message{})

	if isUp {
		query.Where("created_at < ?", targetMessage.CreatedAt)
	} else {
		query.Where("created_at > ?", targetMessage.CreatedAt)
	}

	pagination := h.Paginate(&messages, query, request)

	return &models.BasePaginationResponse[models.Message]{
		Items:       messages,
		CurrentPage: pagination.CurrentPage,
		Count:       pagination.Count,
		TotalPages:  pagination.TotalPages,
	}, nil
}
