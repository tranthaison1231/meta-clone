package handlers

import (
	"github.com/gin-gonic/gin"
	h "github.com/tranthaison1231/meta-clone/api/helpers"
	"github.com/tranthaison1231/meta-clone/api/models"
	"github.com/tranthaison1231/meta-clone/api/services"
)

// @Summary Get Posts
// @ID get-posts
// @Security BearerAuth
// @Success 200 {string} {"status": "success", data: { "posts": []models.Post }, "code": 200}
// @Router /posts [get]
func GetPosts(c *gin.Context) {
	user := c.MustGet("user").(*models.User)

	posts, err := services.GetPosts(user.ID)

	if err != nil {
		h.Fail400(c, err.Error())
	}

	h.Success(c, gin.H{
		"posts": posts,
	})
}

func CreatePost(c *gin.Context) {
	var req models.CreatePostRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		h.Fail400(c, err.Error())
		return
	}

	post, err := services.CreatePost(models.Post{
		Content: req.Content,
		OwnerID: c.MustGet("user").(*models.User).ID,
	})

	if err != nil {
		h.Fail400(c, err.Error())
		return
	}

	h.Success(c, gin.H{
		"post": post,
	})
}
