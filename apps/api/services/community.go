package services

import (
	"github.com/tranthaison1231/meta-clone/api/db"
	"github.com/tranthaison1231/meta-clone/api/models"
)

func GetCommunities(userID string) (*[]models.Community, error) {
	var communities []models.Community
	err := db.DB.Model(&models.Community{}).Where("owner_id = ?", userID).Find(&communities).Error

	if err != nil {
		return nil, err
	}
	return &communities, nil
}

func CreateCommunity(newCommunity models.Community) (*models.Community, error) {
	err := db.DB.Create(&newCommunity).Error

	if err != nil {
		return nil, err
	}
	return &newCommunity, nil
}
