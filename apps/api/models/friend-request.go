package models

type FriendRequest struct {
	Base
	UserID   uint `gorm:"not null" json:"userId"`
	FriendID uint `gorm:"not null" json:"friendId"`
}

type AddFriendRequest struct {
	UserID   uint `json:"userId" validate:"required"`
	FriendID uint `json:"friendId" validate:"required"`
}

type AcceptFriendRequest struct {
	AddFriendRequest
	IsRejecting bool `json:"isRejecting"`
}
