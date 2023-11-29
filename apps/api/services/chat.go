package services

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/tranthaison1231/meta-clone/api/db"
	h "github.com/tranthaison1231/meta-clone/api/helpers"
	"github.com/tranthaison1231/meta-clone/api/models"
)

func GetChat(chatID string) (*models.Chat, error) {
	var chat models.Chat

	err := db.DB.Where("id = ?", chatID).Preload("LastMessage").Preload("Members").First(&chat).Error

	if err != nil {
		return nil, err
	}

	return &chat, nil
}

func GetChats(request *models.GetChatsRequest) (*models.BasePaginationResponse[models.Chat], error) {
	var chats []models.Chat

	query := db.DB.Debug().Model(&models.Chat{}).Preload("LastMessage").Preload("Members")
	if request.MemberIds != nil {
		var chatIds []string
		var lengthCheck int
		if request.IsSingleChat {
			lengthCheck = 2
		} else {
			lengthCheck = len(request.MemberIds)
		}

		chatQuery := db.DB.Raw("SELECT chat_id FROM chat_users WHERE user_id IN ? GROUP BY chat_id HAVING COUNT(*) = ?", request.MemberIds, lengthCheck).Find(&chatIds)

		fmt.Println("query.RowsAffected", query.RowsAffected)

		if chatQuery.RowsAffected > 0 {
			err := query.Where("id IN ?", chatIds).Find(&chats).Error
			if err != nil {
				return nil, err
			}
		} else {
			query.Where("id = ?", "-")
		}
	}

	if err := query.Error; err != nil {
		return nil, err
	}

	fmt.Println("chats", chats)

	pagination := h.Paginate(&chats, query, &request.PaginateRequest)

	return &models.BasePaginationResponse[models.Chat]{
		Items:       chats,
		CurrentPage: pagination.CurrentPage,
		Count:       pagination.Count,
		TotalPages:  pagination.TotalPages,
	}, nil
}

func CreateChat(memberIds []string) (*models.Chat, error) {
	var users []*models.User

	err := db.DB.Where("id IN (?)", memberIds).Find(&users).Error

	if err != nil {
		return nil, err
	}

	newChat := &models.Chat{}

	createErr := db.DB.Omit("owner_id", "last_message_id").Create(newChat).Error

	if createErr != nil {
		return nil, createErr
	}

	for _, user := range users {
		err := db.DB.Exec("INSERT INTO chat_users(chat_id, user_id) VALUES (?,?)", newChat.ID, user.ID).Error

		if err != nil {
			return nil, err
		}
	}

	return newChat, nil
}

func AddMemberToChat(chatID uint64, memberID string) (*models.Chat, error) {
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

func UpdateLastMessage(chatID string, message *models.Message) error {
	var chat models.Chat
	err := db.DB.Model(&chat).Where("id = ?", chatID).First(&chat).Error

	if err != nil {
		return err
	}

	chat.LastMessageID = message.ID

	err = db.DB.Omit("owner_id").Save(&chat).Error

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
