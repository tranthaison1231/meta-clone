package services

import (
	"github.com/tranthaison1231/meta-clone/api/db"
	"github.com/tranthaison1231/meta-clone/api/models"
)

func GetPosts(userID string) (*[]models.Post, error) {
	var posts []models.Post
	err := db.DB.Model(&models.Community{}).Where("owner_id = ?", userID).Find(&posts).Error

	if err != nil {
		return nil, err
	}
	return &posts, nil
}

func CreatePost(newPost models.Post) (*models.Post, error) {
	err := db.DB.Create(&newPost).Error

	if err != nil {
		return nil, err
	}
	return &newPost, nil
}
