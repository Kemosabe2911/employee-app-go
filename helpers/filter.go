package helpers

import (
	"errors"

	"github.com/Kemosabe2911/employee-app-go/logger"
	"github.com/gin-gonic/gin"
)

type Pagination struct {
	Filter string
	SortBy string
	Order  string
}

func GetFilterValue(c *gin.Context) (string, error) {
	search := c.Query("search")
	if search == "" {
		logger.Info("Search empty")
		return search, errors.New("search empty")
	}
	return search, nil
}

func GetSortingValue(c *gin.Context) (string, string, error) {
	sort_by := c.Query("sort_by")
	if sort_by == "" {
		logger.Info("sort is empty")
		return sort_by, "", errors.New("sort is empty")
	}
	order := c.Query("order")
	if order == "" {
		logger.Info("order is empty")
		return sort_by, "asc", nil
	}
	return sort_by, order, nil
}
