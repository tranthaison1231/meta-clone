package handlers

import (
	"github.com/gin-gonic/gin"
	h "github.com/tranthaison1231/meta-clone/api/helpers"
	"github.com/tranthaison1231/meta-clone/api/models"
	"github.com/tranthaison1231/meta-clone/api/services"
)

// @Summary Get News Feed
// @ID get-news-feed
// @Security BearerAuth
// @Success 200 {string} {"status": "success", data: { "posts": []models.Post }, "code": 200}
// @Router /posts [get]
func GetNewsFeed(c *gin.Context) {
	user := c.MustGet("user").(*models.User)

	posts, err := services.GetPosts(user.ID)

	if err != nil {
		h.Fail400(c, err.Error())
	}

	h.Success(c, gin.H{
		"posts": posts,
	})
}
