package h

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/tranthaison1231/meta-clone/api/db"
	"github.com/tranthaison1231/meta-clone/api/models"
	"gorm.io/gorm"
)

const DefaultPage = 1
const DefaultPageSize = 10

func setRequestDefaultValue(request *models.BasePaginationRequest) {
	if request.Page == 0 {
		request.Page = DefaultPage
	}

	if request.Limit == 0 {
		request.Limit = DefaultPageSize
	}
}

type PaginateResult struct {
	CurrentPage int
	Count       int
	TotalPages  int
}

func ConstructPaginateRequest(c *gin.Context) *models.BasePaginationRequest {
	limit, _ := strconv.ParseInt(c.Param("limit"), 10, 64)
	page, _ := strconv.ParseInt(c.Param("page"), 10, 64)
	orderBy := c.Param("order_by")

	return &models.BasePaginationRequest{
		Page:    int(page),
		Limit:   int(limit),
		OrderBy: orderBy,
	}
}

func Paginate(model interface{}, query *gorm.DB, request *models.BasePaginationRequest) *PaginateResult {
	setRequestDefaultValue(request)

	page := request.Page
	limit := request.Limit
	orderBy := request.OrderBy

	fmt.Println("request", request)

	result := query.Offset(page * limit).Limit(limit)

	if orderBy != "" {
		orderByField := strings.Split(request.OrderBy, ":")[0]
		order := strings.Split(request.OrderBy, ":")[0]

		result = query.Order(orderByField + " " + order)
	}

	var totalItems int64
	db.DB.Model(&model).Count(&totalItems)

	totalPages := int(math.Ceil(float64(totalItems) / float64(limit)))

	fmt.Println("totalItems", totalItems)
	fmt.Println("totalPages", totalPages)

	return &PaginateResult{
		CurrentPage: page,
		Count:       int(result.RowsAffected),
		TotalPages:  int(totalPages),
	}
}
