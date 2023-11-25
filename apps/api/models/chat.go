package models

type Chat struct {
	Base
	Name        string  `gorm:"not null;type:varchar(255)" json:"name"`
	Members     []*User `gorm:"many2many:chat_users;" json:"members"`
	OwnerID     uint    `gorm:"not null" json:"ownerId"`
	LastMessage Message `gorm:"foreignKey:ChatID" json:"lastMessage"`
}

type CreateChatRequest struct {
	Name string `json:"name" validate:"required"`
}

type AddMemberToChatRequest struct {
	UserID uint `json:"userId" validate:"required"`
}
