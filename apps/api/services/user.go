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

	query := db.DB.Preload("FriendRequests").Where("id <> ?", currentUser.ID)

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

		for _, friendRequest := range user.FriendRequests {
			if friendRequest.FriendID == currentUser.ID {
				response.FriendStatus = "RequireAccept"
			}
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
		TotalPage:   pagination.TotalPage,
	}, nil
}

func GetUserFriends(userId uint, request *models.BasePaginationRequest) (*models.BasePaginationResponse[models.User], error) {
	var users []models.User
	query := db.DB.Model(&users).Joins("LEFT JOIN user_friends ON user_friends.user_id = users.id").Where("friend_id = ?", userId)

	if err := query.Error; err != nil {
		return nil, err
	}

	pagination := h.Paginate(&users, query, request)

	fmt.Println("users", users)

	return &models.BasePaginationResponse[models.User]{
		Items:       users,
		CurrentPage: pagination.CurrentPage,
		Count:       pagination.Count,
		TotalPage:   pagination.TotalPage,
	}, nil
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

func AcceptFriend(userId uint, friendId uint, isRejecting bool) (string, error) {
	var friendRequest models.FriendRequest
	var user models.User
	var friend models.User

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

	fmt.Println(isRejecting)

	if isRejecting {
		db.DB.Delete(&friendRequest)
	} else {
		user.Friends = append(user.Friends, &friend)
		friend.Friends = append(friend.Friends, &user)
		db.DB.Save(&user)
		db.DB.Save(&friend)
	}

	return "Added Friend", nil
}

func GetUserByID(id uint) (*models.User, error) {
	var user models.User
	result := db.DB.Where("id = ?", id).First(&user)

	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
