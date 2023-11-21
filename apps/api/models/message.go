package models

type Message struct {
	Base
	Content string `gorm:"not null;type:varchar(255)" json:"content"`
	ChatID  uint64 `gorm:"not null" json:"chat_id"`
	OwnerID uint   `gorm:"not null" json:"owner_id"`
}

type MessageRequest struct {
	Content string `json:"content" validate:"required"`
}
