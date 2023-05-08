package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type TransactionIdHeader struct {
	TransactionId string `header:"X-Request-Id"`
}

// GetTransactionIdHeader
// Retrieves the X-Request-ID header from the supplied context.
func GetTransactionIdHeader(c *gin.Context) (*TransactionIdHeader, error) {
	var header TransactionIdHeader
	if err := c.ShouldBindHeader(&header); err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &header, nil
}
