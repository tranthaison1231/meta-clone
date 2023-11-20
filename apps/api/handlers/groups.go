package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Group struct {
	Name string
	Logo string
}

func GetMyGroups(c *gin.Context) {
	groups := []Group{
		{"Cộng đồng chứng khoán Việt Nam", "https://scontent.fhan2-3.fna.fbcdn.net/v/t39.30808-6/277302816_3092373541093206_6743323560183017342_n.jpg?stp=c39.0.100.100a_dst-jpg_p100x100&_nc_cat=101&ccb=1-7&_nc_sid=5f2048&_nc_eui2=AeElb4EpZT7c38zbKbqCPnA111SVc_ENioXXVJVz8Q2KhZzKtEh4_SM33BNUgszBtKr300v98tclbOvI-dUZ2_n9&_nc_ohc=OjwnHjFwFC0AX9fUHAd&_nc_ad=z-m&_nc_cid=0&_nc_ht=scontent.fhan2-3.fna&oh=00_AfAtTkcpE-tEV4JCSJGLychccvFe9ez95DfTyETYOfGIxQ&oe=655EA938"},
		{"Cộng đồng frontend Việt Nam", "https://scontent.fhan2-3.fna.fbcdn.net/v/t39.30808-6/240655429_4059163764195718_3652059374269714525_n.jpg?stp=c55.0.100.100a_dst-jpg_p100x100&_nc_cat=108&ccb=1-7&_nc_sid=5f2048&_nc_eui2=AeFdIzJYJ1lnZSnh84ndpySKcZIZbYEk2QZxkhltgSTZBqEF4YbLhpD02R-QD0VvBiMrNHp0JmID_73pKS5TjahN&_nc_ohc=AhMIPhWDsjEAX8LoKKw&_nc_ad=z-m&_nc_cid=0&_nc_ht=scontent.fhan2-3.fna&oh=00_AfCM4K6AedgJFkwmc5Busso270WCtPR5cf66fKIiYBk8YA&oe=655E3F8E"},
	}

	c.JSON(http.StatusOK, gin.H{
		"Groups": groups,
	})
}
