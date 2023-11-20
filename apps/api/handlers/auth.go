package handlers

import (
	"net/http"
	"strings"

	"github.com/tranthaison1231/messenger-clone/api/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type AuthController struct {
	DB *gorm.DB
}

func InitAuthController(DB *gorm.DB) AuthController {
	return AuthController{DB}
}

func (ac *AuthController) SignIn(c *gin.Context) {
	var req models.SignInRequest
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

	var user models.User
	result := ac.DB.First(&user, "email = ?", strings.ToLower(req.Email))

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": "123456",
	})
}

func (ac *AuthController) SignUp(c *gin.Context) {
	var req models.SignUpRequest
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
	newUser := models.User{
		Email:    strings.ToLower(req.Email),
		Password: req.Password,
		Gender:   req.Gender,
	}
	result := ac.DB.Create(&newUser)

	if result.Error != nil && strings.Contains(result.Error.Error(), "duplicate key value violates unique") {
		c.JSON(http.StatusConflict, gin.H{"status": "fail", "message": "User with that email already exists"})
		return
	} else if result.Error != nil {
		c.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": "Something bad happened"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": "123456",
	})
}

func (ac *AuthController) GetMe(c *gin.Context) {
	var user models.User
	result := ac.DB.First(&user, "email = ?", strings.ToLower("ttson.1711@gmail.com"))

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Internal server error"})
		return
	}

	userResponse := &models.UserResponse{
		ID:     user.ID,
		Email:  user.Email,
		Gender: user.Gender,
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "data": gin.H{"user": userResponse}})
}
