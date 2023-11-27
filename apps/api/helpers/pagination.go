package h

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
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

type PaginateRequestResult struct {
	CurrentPage int
	Count       int
	TotalPages  int
}

func ConstructPaginateRequest(c *gin.Context) *models.BasePaginationRequest {
	limit, _ := strconv.ParseInt(c.Request.URL.Query().Get("limit"), 10, 64)
	page, _ := strconv.ParseInt(c.Request.URL.Query().Get("page"), 10, 64)
	orderBy := c.Request.URL.Query().Get("orderBy")

	return &models.BasePaginationRequest{
		Page:    int(page),
		Limit:   int(limit),
		OrderBy: orderBy,
	}
}

func Paginate(model interface{}, query *gorm.DB, request *models.BasePaginationRequest) *PaginateRequestResult {
	setRequestDefaultValue(request)

	page := request.Page
	limit := request.Limit
	orderBy := request.OrderBy

	var totalItems int64
	query.Count(&totalItems)

	fmt.Println("totalItems", totalItems)
	fmt.Println("errors", query.Error)

	totalPages := int(math.Ceil(float64(totalItems) / float64(limit)))

	result := query.Offset((page - 1) * limit).Limit(limit)

	if orderBy != "" {
		orderByField := strings.Split(request.OrderBy, ":")[0]
		order := strings.Split(request.OrderBy, ":")[0]

		result = query.Order(orderByField + " " + order)
	}

	result.Find(model)

	return &PaginateRequestResult{
		CurrentPage: page,
		Count:       int(result.RowsAffected),
		TotalPages:  int(totalPages),
	}
}
