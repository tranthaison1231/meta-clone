package models

type Chat struct {
	Base
	Name          string   `gorm:"not null;type:varchar(255)" json:"name"`
	Members       []*User  `gorm:"many2many:chat_users;" json:"members"`
	OwnerID       string   `gorm:"not null" json:"ownerId"`
	Owner         User     `gorm:"foreignKey:OwnerID;references:ID" json:"owner"`
	LastMessageID string   `json:"lastMessageId"`
	LastMessage   *Message `gorm:"foreignKey:LastMessageID;references:ID" json:"lastMessage"`
}

type CreateChatRequest struct {
	Name string `json:"name" validate:"required"`
}

type AddMemberToChatRequest struct {
	UserID uint `json:"userId" validate:"required"`
}
