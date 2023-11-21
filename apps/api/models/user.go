package models

import (
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       uint   `gorm:"primary_key;autoIncrement"`
	Email    string `gorm:"uniqueIndex;not null;type:varchar(255)"`
	Salt     string `gorm:"not null"`
	Password string `gorm:"not null"`
	Gender   string `gorm:"not null"`
}

type UserResponse struct {
	ID     uint   `json:"id,omitempty"`
	Email  string `json:"email,omitempty"`
	Gender string `json:"gender,omitempty"`
}

type SignInRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type Gender string

type SignUpRequest struct {
	SignInRequest
	Gender string `json:"gender" validate:"required"`
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
