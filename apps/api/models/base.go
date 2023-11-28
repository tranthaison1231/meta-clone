package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Base struct {
	ID        string     `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `sql:"index" json:"deletedAt"`
}

func (base *Base) BeforeCreate(tx *gorm.DB) (err error) {
	base.ID = uuid.New().String()
	return nil
}

type BasePaginationRequest struct {
	Page    int    `json:"page"`
	Limit   int    `json:"limit"`
	OrderBy string `json:"orderBy"`
}

type BasePaginationResponse[T any] struct {
	Items       []T `json:"items"`
	CurrentPage int `json:"currentPage"`
	Count       int `json:"count"`
	TotalPage   int `json:"totalPage"`
}
