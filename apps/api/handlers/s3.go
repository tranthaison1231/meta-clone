package handlers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	h "github.com/tranthaison1231/meta-clone/api/helpers"
	"github.com/tranthaison1231/meta-clone/api/models"
)

func GetPresignedUrl(c *gin.Context) {
	var req models.GetPresignedUrlRequestDto

	if err := h.CheckBindAndValidate(&req, c); err != nil {
		fmt.Println("++++++++++++ | GetPresignedUrl | CheckBindAndValidate err | ++++++++++++", err)
		return
	}

	s3PresignClient, err := h.InitS3PresignClient(c)

	if err != nil {
		fmt.Println("++++++++++++ | GetPresignedUrl | InitS3PresignClient err | ++++++++++++", err)
		h.Fail400(c, err.Error())
		return
	}

	url, err := s3PresignClient.GetPresignedUrl(c, req)

	if err != nil {
		fmt.Println("++++++++++++ | GetPresignedUrl | GetPresignedUrl err | ++++++++++++", err)

		h.Fail400(c, err.Error())
		return
	}

	h.Success(c, gin.H{
		"url": url,
	})

}
