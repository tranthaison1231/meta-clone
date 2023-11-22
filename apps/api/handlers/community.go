package handlers

import (
	"github.com/gin-gonic/gin"
	h "github.com/tranthaison1231/meta-clone/api/helpers"
	"github.com/tranthaison1231/meta-clone/api/models"
	"github.com/tranthaison1231/meta-clone/api/services"
)

type Community struct {
	Name string
	Logo string
}

func GetCommunities(c *gin.Context) {
	user := c.MustGet("user").(*models.User)

	communities, err := services.GetCommunities(user.ID)

	if err != nil {
		h.Fail400(c, err.Error())
	}

	h.Success(c, gin.H{
		"communities": communities,
	})
}

func CreateCommunity(c *gin.Context) {
	var req models.CreateCommunityRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		h.Fail400(c, err.Error())
		return
	}

	community, err := services.CreateCommunity(models.Community{
		Name:    req.Name,
		Logo:    req.Logo,
		OwnerID: c.MustGet("user").(*models.User).ID,
	})

	if err != nil {
		h.Fail400(c, err.Error())
		return
	}

	h.Success(c, gin.H{
		"community": community,
	})
}
