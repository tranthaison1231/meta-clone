package models

type Message struct {
	Base
	Content string `gorm:"not null;type:varchar(255)" json:"content"`
	ChatID  uint64 `gorm:"not null" json:"chatId"`
	OwnerID uint   `gorm:"not null" json:"ownerId"`
}

type MessageRequest struct {
	Content string `json:"content" validate:"required"`
}
