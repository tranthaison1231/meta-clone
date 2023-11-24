package models

import (
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Base
	FirstName   string      `gorm:"type:varchar(255)" json:"first_name"`
	LastName    string      `gorm:"type:varchar(255)" json:"last_name"`
	Email       string      `gorm:"uniqueIndex;not null;type:varchar(255)" json:"email"`
	Salt        string      `gorm:"not null" json:"-"`
	Avatar      string      `gorm:"type:text" json:"avatar"`
	Password    string      `gorm:"not null" json:"-"`
	Gender      string      `gorm:"not null" json:"gender"`
	Chats       []*Chat     `gorm:"many2many:chat_users;" json:"chats"`
	Communities []Community `gorm:"many2many:user_communities;" json:"communities"`
}

type SignInRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type Gender string

type SignUpRequest struct {
	SignInRequest
	Gender string `json:"gender" validate:"required" enum:"male,female"`
	Avatar string `json:"avatar"`
}

type UpdateUserRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Avatar    string `json:"avatar"`
	Gender    string `json:"gender" enum:"male,female"`
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
