package services

import (
	"fmt"

	"github.com/tranthaison1231/meta-clone/api/db"
	h "github.com/tranthaison1231/meta-clone/api/helpers"
	"github.com/tranthaison1231/meta-clone/api/models"
)

func CreateNewConnection(userID string, connectionID string) (*models.Connection, error) {
	fmt.Println("==================== | CREATE NEW CONNECTION | USER ID ====================", userID)
	fmt.Println("==================== | CREATE NEW CONNECTION | CONNECTION ID ====================", connectionID)

	var existingConnection models.Connection

	query := db.DB.Where("connection_id = ? AND user_id = ?", connectionID, userID).First(&existingConnection)

	fmt.Println("==================== | CREATE NEW CONNECTION | EXISTING CONNECTION ====================", existingConnection)
	fmt.Println("==================== | CREATE NEW CONNECTION | QUERY ROWS AFFECTED ====================", query.RowsAffected)

	if query.RowsAffected != 0 {
		return &existingConnection, nil
	}

	newConnection := models.Connection{
		ConnectionID: connectionID,
		UserID:       userID,
	}

	fmt.Println("==================== | CREATE NEW CONNECTION | NEW CONNECTION ====================", newConnection)

	err := db.DB.Save(&newConnection).Error

	fmt.Println("==================== | CREATE NEW CONNECTION | CREATE NEW CONNECTION ERROR ====================", err)

	if err != nil {
		return nil, err
	}

	return &newConnection, nil
}

func GetOnlineUserIds(request models.BasePaginationRequest) (*models.BasePaginationResponse[string], error) {
	var userIds []string

	query := db.DB.Raw("SELECT DISTINCT(user_id) from connections")

	pagination, err := h.Paginate(&userIds, query, &request)

	if err != nil {
		return nil, err
	}

	return &models.BasePaginationResponse[string]{
		Items:       userIds,
		CurrentPage: pagination.CurrentPage,
		Count:       pagination.Count,
		TotalPages:  pagination.TotalPages,
	}, nil
}

func RemoveConnection(connectionID string) error {
	err := db.DB.Where("connection_id = ?", connectionID).Delete(models.Connection{}).Error

	return err
}
