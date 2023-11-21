package models

type Message struct {
	ID      uint   `gorm:"primary_key;autoIncrement"`
	Content string `gorm:"not null;type:varchar(255)"`
}
