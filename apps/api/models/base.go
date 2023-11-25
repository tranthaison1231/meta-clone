package models

import "time"

type Base struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `sql:"index" json:"deletedAt"`
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
