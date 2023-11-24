package models

import (
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Base
	Email          string          `gorm:"uniqueIndex;not null;type:varchar(255)" json:"email"`
	Salt           string          `gorm:"not null" json:"-"`
	Avatar         string          `gorm:"type:text" json:"avatar"`
	Password       string          `gorm:"not null" json:"-"`
	Gender         string          `gorm:"not null" json:"gender"`
	Chats          []*Chat         `gorm:"many2many:chat_users;" json:"chats"`
	Communities    []Community     `gorm:"many2many:user_communities;" json:"communities"`
	Friends        []*User         `gorm:"many2many:user_friends" json:"friends"`
	FriendRequests []FriendRequest `json:"friend_requests"`
}

type UserFriend struct {
	Base
	UserID   string `gorm:"not null" json:"user_id"`
	FriendID string `gorm:"not null" json:"friend_id"`
}

type SignInRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type Gender string

type SignUpRequest struct {
	SignInRequest
	Gender string `json:"gender" validate:"required"`
	Avatar string `json:"avatar"`
}

type GetUserFriendRequest struct {
	UserID uint `json:"user_id" validate:"required"`
}

type GetUserResponse struct {
	User
	FriendStatus string `json:"friend_status"`
}

func (u *User) ValidatePwdStaticHash(password string) error {
	if password == "" {
		return errors.WithStack(errors.New("Password cannot be empty"))
	}
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}
