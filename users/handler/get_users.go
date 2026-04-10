package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tristaamne/flowershopbe-v4/common/pagination"
	"github.com/tristaamne/flowershopbe-v4/common/utils"
	"github.com/tristaamne/flowershopbe-v4/users/repository"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetUsersByCondition(coll *mongo.Collection) gin.HandlerFunc {
	return func(c *gin.Context) {
		rawQuery := make(map[string]interface{})
		if err := c.ShouldBindQuery(&rawQuery); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid query parameters"})
			return
		}

		var pg pagination.PaginationQuery
		if err := c.ShouldBindQuery(&pg); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid pagination settings"})
			return
		}

		filter := utils.MapToBSon(rawQuery)

		excludedFields := []string{"page", "limit", "sort_by", "order"}
		for _, field := range excludedFields {
			delete(filter, field)
		}

		opts := pagination.ParsePagingOption(pg)

		userData, err := repository.GetUserByCondition(coll, filter, opts)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"msg":  "success",
			"data": userData,
		})
	}
}
