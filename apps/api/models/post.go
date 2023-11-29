package models

type Post struct {
	Base
	Content string `gorm:"not null;type:text" json:"content"`
	OwnerID string `gorm:"not null" json:"ownerId"`
	Owner   User   `gorm:"foreignKey:OwnerID;references:ID" json:"owner"`
}

type CreatePostRequest struct {
	Content string `json:"content" validate:"required"`
}
