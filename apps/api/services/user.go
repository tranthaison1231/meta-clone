package services

import (
	"fmt"

	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	h "github.com/tranthaison1231/meta-clone/api/helpers"
	"github.com/tranthaison1231/meta-clone/api/models"

	"github.com/tranthaison1231/meta-clone/api/db"
)

func GetUserByMail(mail string) (*models.User, error) {
	var user models.User
	result := db.DB.Preload("FriendRequests").Where("email = ?", mail).First(&user)

	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func GetUsers(request *models.BasePaginationRequest, currentUser *models.User) (*models.BasePaginationResponse[models.GetUserResponse], error) {
	var users []models.User

	query := db.DB.Model(&models.User{}).Where("id <> ?", currentUser.ID)

	if err := query.Error; err != nil {
		return nil, err
	}

	pagination := h.Paginate(&users, query, request)

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
		friendRequestQuery := db.DB.Model(&models.FriendRequest{}).Where("friend_id = ? AND user_id = ?", currentUser.ID, user.ID).First(&models.FriendRequest{})

		if friendRequestQuery.RowsAffected == 1 {
			response.FriendStatus = "RequireAccept"

		}

		var userFriend models.UserFriend

		query := db.DB.Raw("SELECT * FROM user_friends WHERE user_id = ? AND friend_id = ?", user.ID, currentUser.ID).Scan(&userFriend)

		if query.RowsAffected == 1 {
			response.FriendStatus = "Friend"
		}

		getUserResponse = append(getUserResponse, response)
	}

	return &models.BasePaginationResponse[models.GetUserResponse]{
		Items:       getUserResponse,
		CurrentPage: pagination.CurrentPage,
		Count:       pagination.Count,
		TotalPages:  pagination.TotalPages,
	}, nil
}

func GetUserFriends(userId string, request *models.BasePaginationRequest) (*models.BasePaginationResponse[models.User], error) {
	var users []models.User
	query := db.DB.Model(&users).Joins("LEFT JOIN user_friends ON user_friends.user_id = users.id").Where("friend_id = ?", userId)

	if err := query.Error; err != nil {
		return nil, err
	}

	pagination := h.Paginate(&users, query, request)

	return &models.BasePaginationResponse[models.User]{
		Items:       users,
		CurrentPage: pagination.CurrentPage,
		Count:       pagination.Count,
		TotalPages:  pagination.TotalPages,
	}, nil
}

func AddFriend(userId string, friendId string) (*models.FriendRequest, error) {
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

	newFriendRequest := models.FriendRequest{
		UserID:   userId,
		FriendID: friendId,
	}

	err := db.DB.Create(&newFriendRequest).Error

	if err != nil {
		return nil, err
	}

	return &newFriendRequest, nil
}

func AcceptFriend(userId string, friendId string, isRejecting bool) (string, error) {
	var friendRequest models.FriendRequest
	var user models.User
	var friend *models.User

	result := db.DB.Where("user_id = ? AND friend_id = ?", userId, friendId).First(&friendRequest)

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

	if !isRejecting {
		err1 := db.DB.Exec("INSERT INTO user_friends (user_id, friend_id) VALUES(?, ?)", userId, friendId).Error
		if err1 != nil {
			fmt.Println(err1)
			return "", err1
		}
		err2 := db.DB.Exec("INSERT INTO user_friends (user_id, friend_id) VALUES(?, ?)", friendId, userId).Error
		if err2 != nil {
			fmt.Println(err1)
			return "", err2
		}
		db.DB.Delete(&friendRequest)
		return "Added Friend", nil
	}

	db.DB.Delete(&friendRequest)
	return "Rejected Friend", nil
}

func GetUserByID(id string) (*models.User, error) {
	var user models.User
	result := db.DB.Where("id = ?", id).First(&user)

	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
