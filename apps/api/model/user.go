package model

type User struct {
	ID       uint   `gorm:"primary_key;autoIncrement"`
	Email    string `gorm:"uniqueIndex;not null;type:varchar(255)"`
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
