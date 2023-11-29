package models

type FriendRequest struct {
	Base
	UserID   string `gorm:"not null" json:"userId"`
	User     User   `gorm:"foreignKey:UserID;references:ID;not null" json:"user"`
	FriendID string `gorm:"not null" json:"friendId"`
	Friend   User   `gorm:"foreignKey:FriendID;references:ID;not null" json:"friend"`
}

type AddFriendRequest struct {
	UserID   string `json:"userId" validate:"required"`
	FriendID string `json:"friendId" validate:"required"`
}

type AcceptFriendRequest struct {
	AddFriendRequest
	IsRejecting bool `json:"isRejecting"`
}
