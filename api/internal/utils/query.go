package utils

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func QueryLimit(c *gin.Context) int64 {
	limit := c.DefaultQuery("limit", "0")
	limitInt64, err := strconv.ParseInt(limit, 10, 64)
	if err != nil {
		return 0
	}
	return limitInt64
}
