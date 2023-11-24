package handlers

import (
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/tranthaison1231/meta-clone/api/db"
	h "github.com/tranthaison1231/meta-clone/api/helpers"
	"github.com/tranthaison1231/meta-clone/api/models"
	"github.com/tranthaison1231/meta-clone/api/services"
	"golang.org/x/crypto/bcrypt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SignIn(c *gin.Context) {
	var req models.SignInRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.Fail400(c, err.Error())
		return
	}
	v := validator.New()
	err := v.Struct(req)
	if err != nil {
		h.Fail400(c, err.Error())
		return
	}

	user, err := services.GetUserByMail(req.Email)

	if err != nil {
		h.Fail400(c, err.Error())
		return
	}

	if err := user.ValidatePwdStaticHash(req.Password); err != nil {
		h.Fail400(c, "You entered an invalid password")
		return
	}

	token, err := services.GenerateToken(user)
	if err != nil {
		h.Fail400(c, err.Error())
		return
	}

	h.Success(c, gin.H{
		"token": token,
	})
}

func SignUp(c *gin.Context) {
	var req models.SignUpRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.Fail400(c, err.Error())
		return
	}
	v := validator.New()
	err := v.Struct(req)
	if err != nil {
		h.Fail400(c, err.Error())
		return
	}
	pwd, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		h.Fail400(c, err.Error())
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
		h.Fail409(c, "User with that email already exists")
		return
	} else if result.Error != nil {
		h.Fail502(c, result.Error.Error())
		return
	}
	token, err := services.GenerateToken(&newUser)
	if err != nil {
		h.Fail400(c, err.Error())
		return
	}

	h.Success(c, gin.H{
		"token": token,
	})
}

func GetMe(c *gin.Context) {
	user := c.MustGet("user").(*models.User)

	h.Success(c, gin.H{
		"user": user,
	})
}

func UpdateMe(c *gin.Context) {
	var req models.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.Fail400(c, err.Error())
		return
	}
	v := validator.New()
	err := v.Struct(req)
	if err != nil {
		h.Fail400(c, err.Error())
		return
	}

	user := c.MustGet("user").(*models.User)

	if req.FirstName != "" {
		user.FirstName = req.FirstName
	}
	if req.LastName != "" {
		user.LastName = req.LastName
	}
	if req.Avatar != "" {
		user.Avatar = req.Avatar
	}
	if req.Gender != "" {
		user.Gender = req.Gender
	}
	db.DB.Save(&user)

	h.Success(c, gin.H{
		"user": user,
	})
}
