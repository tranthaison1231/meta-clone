package models

type FriendRequest struct {
	Base
	UserID   uint `json:"userId"`
	FriendID uint `json:"friendId"`
}

type AddFriendRequest struct {
	UserID   uint `json:"userId" validate:"required"`
	FriendID uint `json:"friendId" validate:"required"`
}

type AcceptFriendRequest struct {
	AddFriendRequest
}
