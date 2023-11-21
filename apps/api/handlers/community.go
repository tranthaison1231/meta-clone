package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tranthaison1231/messenger-clone/api/models"
	"github.com/tranthaison1231/messenger-clone/api/services"
)

type Community struct {
	Name string
	Logo string
}

func GetCommunities(c *gin.Context) {
	user := c.MustGet("user").(*models.User)

	communities, err := services.GetCommunities(user.ID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{
		"communities": communities,
	})
}

func CreateCommunity(c *gin.Context) {
	var req models.CreateCommunityRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	community, err := services.CreateCommunity(models.Community{
		Name:    req.Name,
		Logo:    req.Logo,
		OwnerID: c.MustGet("user").(*models.User).ID,
	})

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"community": community,
	})
}
