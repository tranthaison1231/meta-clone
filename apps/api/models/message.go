package models

type Message struct {
	Base
	Content string `gorm:"not null;type:varchar(255)" json:"content"`
	ChatID  string `gorm:"not null" json:"chatId"`
	Chat    Chat   `gorm:"foreignKey:ChatID;references:ID" json:"chat"`
	OwnerID string `gorm:"not null;type:uuid" json:"ownerId"`
	Owner   User   `gorm:"foreignKey:OwnerID;references:ID" json:"owner"`
}

type MessageRequest struct {
	Content string `json:"content" validate:"required"`
}
