package handlers

import (
	"net/http"
	"strings"

	"github.com/tranthaison1231/messenger-clone/api/db"
	"github.com/tranthaison1231/messenger-clone/api/models"
	"github.com/tranthaison1231/messenger-clone/api/services"
	"golang.org/x/crypto/bcrypt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func SignIn(c *gin.Context) {
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

	user, err := services.GetUserByMail(req.Email)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	if err := user.ValidatePwdStaticHash(req.Password); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "You entered an invalid password"})
		return
	}

	token, err := services.GenerateToken(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func SignUp(c *gin.Context) {
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
	pwd, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	newUser := models.User{
		Email:    strings.ToLower(req.Email),
		Password: string(pwd),
		Gender:   req.Gender,
		Avatar:   req.Avatar,
	}
	result := db.DB.Create(&newUser)

	if result.Error != nil && result.Error == gorm.ErrDuplicatedKey {
		c.JSON(http.StatusConflict, gin.H{"status": "fail", "message": "User with that email already exists"})
		return
	} else if result.Error != nil {
		c.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": result.Error.Error()})
		return
	}
	token, err := services.GenerateToken(&newUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func GetMe(c *gin.Context) {
	user := c.MustGet("user").(*models.User)

	userResponse := &models.UserResponse{
		ID:        user.ID,
		Email:     user.Email,
		Gender:    user.Gender,
		Avatar:    user.Avatar,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "data": gin.H{"user": userResponse}})
}
