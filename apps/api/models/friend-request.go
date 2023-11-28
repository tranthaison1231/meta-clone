package models

type FriendRequest struct {
	Base
	UserID   string `gorm:"not null" json:"userId"`
	FriendID string `gorm:"not null" json:"friendId"`
}

type AddFriendRequest struct {
	UserID   string `json:"userId" validate:"required"`
	FriendID string `json:"friendId" validate:"required"`
}

type AcceptFriendRequest struct {
	AddFriendRequest
	IsRejecting bool `json:"isRejecting"`
}
