package models

type Connection struct {
	Base
	ConnectionID string `gorm:"not null,unique" json:"connectionId"`
	UserID       string `gorm:"not null,unique" json:"userId"`
	User         User   `gorm:"foreignKey:UserID;references:ID;not null" json:"user"`
}
