package models

import "time"

type Base struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`
}

type BasePaginationRequest struct {
	Page    int    `json:"page"`
	Limit   int    `json:"limit"`
	OrderBy string `json:"order_by"`
}

type BasePaginationResponse[T any] struct {
	Items       []T `json:"items"`
	CurrentPage int `json:"current_page"`
	Count       int `json:"count"`
	TotalPages  int `json:"total_pages"`
}
