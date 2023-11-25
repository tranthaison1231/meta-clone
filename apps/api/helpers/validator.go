package h

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func CheckBindAndValidate(req interface{}, c *gin.Context) error {
	if err := c.ShouldBindJSON(&req); err != nil {
		Fail400(c, err.Error())
		return err
	}
	v := validator.New()
	err := v.Struct(req)
	if err != nil {
		Fail400(c, err.Error())
		return err
	}
	return nil
}
