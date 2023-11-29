package models

type Media struct {
	Base
	URL string `gorm:"not null;type:text" json:"url"`
}
