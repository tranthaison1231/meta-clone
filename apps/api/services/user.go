package services

import (
	"github.com/tranthaison1231/messenger-clone/api/models"

	"github.com/tranthaison1231/messenger-clone/api/db"
)

func GetUserByMail(mail string) (*models.User, error) {
	var user models.User
	result := db.DB.Where("mail = ?", mail).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
