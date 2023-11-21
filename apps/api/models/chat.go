package models

import "gorm.io/gorm"

type Chat struct {
	gorm.Model
	ID   uint   `gorm:"primary_key;autoIncrement"`
	Name string `gorm:"not null;type:varchar(255)"`
	// Members []*User `gorm:"many2many:chat_users;"`
	OwnerID uint `gorm:"not null"`
}

type CreateChatRequest struct {
	Name string `json:"name" validate:"required"`
}

type AddMemberToChatRequest struct {
	UserID uint `json:"user_id" validate:"required"`
}
