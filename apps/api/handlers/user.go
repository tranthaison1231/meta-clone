package handlers

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	h "github.com/tranthaison1231/meta-clone/api/helpers"
	"github.com/tranthaison1231/meta-clone/api/models"
	"github.com/tranthaison1231/meta-clone/api/services"
)

func AddFriend(c *gin.Context) {
	var req models.AddFriendRequest
	if err := h.CheckBindAndValidate(&req, c); err != nil {
		h.Fail400(c, err.Error())
		return
	}

	friendRequest, err := services.AddFriend(req.UserID, req.FriendID)

	if err != nil {
		h.Fail400(c, err.Error())
		return
	}

	h.Success(c, gin.H{
		"friendRequest": friendRequest,
	})
}

func AcceptFriend(c *gin.Context) {
	var req models.AcceptFriendRequest
	if err := h.CheckBindAndValidate(&req, c); err != nil {
		return
	}

	message, err := services.AcceptFriend(req.UserID, req.FriendID, req.IsRejecting)

	if err != nil {
		h.Fail400(c, err.Error())
		return
	}

	h.Success(c, gin.H{
		"message": message,
	})
}

func GetUserFriends(c *gin.Context) {
	fmt.Println("Page", c.Param("page"), "Limit", c.Param("limit"), c.Params)

	requestParams := h.ConstructPaginateRequest(c)

	userId, err := strconv.ParseInt(c.Param("userID"), 10, 64)

	if err != nil {
		h.Fail400(c, err.Error())
		return
	}

	users, err := services.GetUserFriends(uint(userId), requestParams)

	if err != nil {
		h.Fail400(c, err.Error())
		return
	}

	h.Success(c, users)
}

func GetUsers(c *gin.Context) {
	requestParams := h.ConstructPaginateRequest(c)

	currentUser := c.MustGet("user").(*models.User)

	users, err := services.GetUsers(requestParams, currentUser)

	if err != nil {
		h.Fail400(c, err.Error())
		return
	}
	h.Success(c, users)
}
