package services

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/tranthaison1231/messenger-clone/api/models"

	"github.com/jinzhu/copier"
	"github.com/tranthaison1231/messenger-clone/api/db"
)

func GetUserByMail(mail string) (*models.User, error) {
	var user models.User
	result := db.DB.Preload("FriendRequests").Where("email = ?", mail).First(&user)

	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func GetUsers(currentUser *models.User) (*[]models.GetUserResponse, error) {
	var users []models.User

	err := db.DB.Preload("Friends").Preload("FriendRequests").Where("id <> ?", currentUser.ID).Find(&users).Error

	if err != nil {
		return nil, err
	}

	var getUserResponse []models.GetUserResponse
	for _, user := range users {
		var response models.GetUserResponse

		copier.Copy(&response, &user)

		response.FriendStatus = "None"

		for _, friendRequest := range currentUser.FriendRequests {
			if friendRequest.FriendID == user.ID {
				response.FriendStatus = "Pending"
			}
		}

		for _, friendRequest := range user.FriendRequests {
			if friendRequest.FriendID == currentUser.ID {
				response.FriendStatus = "RequireAccept"
			}
		}

		for _, friend := range user.Friends {
			if friend.ID == currentUser.ID {
				response.FriendStatus = "Friend"
			}
		}

		getUserResponse = append(getUserResponse, response)
	}

	return &getUserResponse, nil
}

func GetUserFriends(userId uint) (*[]models.User, error) {
	var users []models.User
	err := db.DB.Model(&users).Where("id = ?", userId).Association("Friends").Find(&users)

	if err != nil {
		return nil, err
	}

	return &users, nil
}

func AddFriend(userId uint, friendId uint) (*models.FriendRequest, error) {
	var friendRequest models.FriendRequest
	var user models.User

	result := db.DB.Where("user_id = ? AND friend_id = ? OR user_id = ? AND friend_id = ?", userId, friendId, friendId, userId).Find(&friendRequest)

	if err := result.Error; err != nil {
		return nil, err
	}

	if result.RowsAffected > 0 {
		return nil, errors.New("Friend request already sent!")
	}

	userErr := db.DB.Where("id = ?", userId).Find(&user).Error

	if userErr != nil {
		return nil, userErr
	}

	for _, value := range user.FriendRequests {
		if value.UserID == userId && value.FriendID == friendId {
			return nil, errors.New("You already added this friend!")
		}
	}

	new_friend_request := models.FriendRequest{
		UserID:   userId,
		FriendID: friendId,
	}

	err := db.DB.Create(&new_friend_request).Error

	if err != nil {
		return nil, err

	}

	return &new_friend_request, nil
}

func AcceptFriend(userId uint, friendId uint) (string, error) {
	var friendRequest models.FriendRequest
	var user models.User
	var friend models.User

	result := db.DB.Where("user_id = ? AND friend_id = ?", userId, friendId).First(&friendRequest)

	fmt.Println("=======", userId, friendId)

	if err := result.Error; err != nil {
		return "", err
	}

	if result.RowsAffected < 1 {
		return "", errors.New("Friend request doesn't exist!")
	}

	err1 := db.DB.Where("id = ?", userId).First(&user).Error
	err2 := db.DB.Where("id = ?", friendId).First(&friend).Error

	if err1 != nil {
		return "", err1
	}

	if err2 != nil {
		return "", err2
	}

	user.Friends = append(user.Friends, &friend)
	friend.Friends = append(friend.Friends, &user)

	db.DB.Save(&user)
	db.DB.Save(&friend)

	db.DB.Delete(&friendRequest)

	return "Added Friend", nil
}

// func DenyFriend(userId uint, friendId uint)
