package models

type Message struct {
	Base
	Content        string   `gorm:"not null;type:varchar(255)" json:"content"`
	ChatID         string   `gorm:"not null" json:"chatId"`
	Chat           Chat     `gorm:"foreignKey:ChatID;references:ID" json:"chat"`
	OwnerID        string   `gorm:"not null;type:uuid" json:"ownerId"`
	Owner          User     `gorm:"foreignKey:OwnerID;references:ID" json:"owner"`
	IsSeen         bool     `gorm:"not null" json:"isSeen"`
	ReplyMessageID string   `json:"replyMessageId"`
	ReplyMessage   *Message `foreignKey:"ReplyMessageID" json:"replyMessage"`
	MessageType    string   `gorm:"not null" json:"messageType"`
}

type SendMessageInputDto struct {
	Content        string `json:"content" validate:"required"`
	MessageType    string `json:"messageType"`
	ReplyMessageID string `json:"replyMessageId"`
}

type SeenMessageInputDto struct {
	MessageID string `json:"messageId"`
}
