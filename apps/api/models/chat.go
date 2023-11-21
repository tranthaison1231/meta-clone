package models

type Chat struct {
	Base
	Name        string  `gorm:"not null;type:varchar(255)" json:"name"`
	Members     []*User `gorm:"many2many:chat_users;" json:"members"`
	OwnerID     uint    `gorm:"not null" json:"owner_id"`
	LastMessage Message `gorm:"foreignKey:ChatID" json:"last_message"`
}

type CreateChatRequest struct {
	Name string `json:"name" validate:"required"`
}

type AddMemberToChatRequest struct {
	UserID uint `json:"user_id" validate:"required"`
}
