package models

type Community struct {
	Base
	Name    string  `gorm:"not null;type:varchar(255)" json:"name"`
	Logo    string  `gorm:"not null;type:text;" json:"logo"`
	OwnerID string  `gorm:"not null" json:"ownerId"`
	Owner   *User   `gorm:"foreignKey:OwnerID;references:ID" json:"owner"`
	Users   []*User `gorm:"many2many:user_communities;" json:"users"`
}

type CreateCommunityRequest struct {
	Name string `json:"name" validate:"required"`
	Logo string `json:"logo" validate:"required"`
}
