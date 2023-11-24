package models

type FriendRequest struct {
	Base
	UserID   uint `json:"user_id"`
	FriendID uint `json:"friend_id"`
}

type AddFriendRequest struct {
	UserID   uint `json:"user_id" validate:"required"`
	FriendID uint `json:"friend_id" validate:"required"`
}

type AcceptFriendRequest struct {
	AddFriendRequest
}
