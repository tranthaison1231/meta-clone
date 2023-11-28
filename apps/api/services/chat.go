package services

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/tranthaison1231/meta-clone/api/db"
	h "github.com/tranthaison1231/meta-clone/api/helpers"
	"github.com/tranthaison1231/meta-clone/api/models"
)

func GetChats(request *models.BasePaginationRequest, memberIds []uint) (*models.BasePaginationResponse[models.Chat], error) {
	var chats []models.Chat

	query := db.DB.Model(&models.Chat{}).Preload("LastMessage").Preload("Members")
	if memberIds != nil {
		var chatId string

		chatUserQuery := db.DB.Raw("SELECT chat_id FROM chat_users WHERE user_id IN (?) GROUP BY chat_id HAVING COUNT(*) = ?", memberIds, len(memberIds)).First(&chatId)

		if err := chatUserQuery.Error; err != nil {
			return nil, err
		}

		err := query.Where("id = ?", chatId).Find(&chats).Error

		if err != nil {
			return nil, err
		}
	}

	if err := query.Error; err != nil {
		fmt.Println("errors", err)
		return nil, err
	}

	pagination := h.Paginate(&chats, query, request)

	return &models.BasePaginationResponse[models.Chat]{
		Items:       chats,
		CurrentPage: pagination.CurrentPage,
		Count:       pagination.Count,
		TotalPages:  pagination.TotalPages,
	}, nil
}

func CreateChat(memberIds []uint) (*models.Chat, error) {
	var users []*models.User

	err := db.DB.Where("id IN (?)", memberIds).Find(&users).Error

	if err != nil {
		return nil, err
	}

	newChat := &models.Chat{
		Members: users,
	}

	createErr := db.DB.Create(newChat).Error

	if createErr != nil {
		return nil, createErr
	}
	return newChat, nil
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

func UpdateLastMessage(chatID string, message models.Message) error {
	var chat models.Chat
	err := db.DB.Model(&chat).Where("id = ?", chatID).First(&chat).Error

	chat.LastMessageID = message.ID

	err = db.DB.Save(&chat).Error

	return err
}

func GetChatMessages(request *models.GetChatMessagesRequest) (*models.BasePaginationResponse[models.Message], error) {
	var messages []models.Message
	var targetMessage models.Message

	targetMessageQuery := db.DB.Where("id = ?", request.TargetMessageID).First(&targetMessage)

	if err := targetMessageQuery.Error; err != nil {
		return nil, err
	}

	if targetMessageQuery.RowsAffected < 1 {
		return nil, errors.New("Target message not found")
	}

	query := db.DB.Model(&models.Message{}).Where("chat_id = ?", request.ChatID)

	if request.IsUp {
		query.Where("created_at <= ?", targetMessage.CreatedAt)
	} else {
		query.Where("created_at => ?", targetMessage.CreatedAt)
	}

	pagination := h.Paginate(&messages, query, &request.PaginateRequest)

	return &models.BasePaginationResponse[models.Message]{
		Items:       messages,
		CurrentPage: pagination.CurrentPage,
		Count:       pagination.Count,
		TotalPages:  pagination.TotalPages,
	}, nil
}
