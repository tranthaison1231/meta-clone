package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type SignInRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func SignIn(c *gin.Context) {
	var req SignInRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	v := validator.New()
	err := v.Struct(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"token": "123456",
	})
}

func SignUp(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"token": "123456",
	})
}

func GetMe(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"id":     "123456",
		"name":   "Son Tran",
		"avatar": "https://scontent.fhan2-4.fna.fbcdn.net/v/t39.30808-1/275227304_2189265651227084_6306293141393566490_n.jpg?stp=dst-jpg_p100x100&_nc_cat=100&ccb=1-7&_nc_sid=5f2048&_nc_eui2=AeGSrLdIzPJQLWlPGZnNs2k60UBRHL1Rt6nRQFEcvVG3qchcOQx5CR7QUYLj9gnSwxmMdT1TVOXg8lKn0RSMS2cU&_nc_ohc=iOcN-3H4vTkAX_84_ws&_nc_ad=z-m&_nc_cid=0&_nc_ht=scontent.fhan2-4.fna&oh=00_AfB3O50uwv_-Omrjj1WgK2sKF9Y2yeAAG-HKNd1Vcn6Ymg&oe=655DD252",
	})
}
