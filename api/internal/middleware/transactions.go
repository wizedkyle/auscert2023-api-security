package middleware

import (
	"github.com/ForgeResponse/ForgeResponse/v2/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

// ValidateTransactionId
// Checks for X-Request-ID on each request, if it is not found the header is added.
func ValidateTransactionId() gin.HandlerFunc {
	return func(c *gin.Context) {
		header, err := utils.GetTransactionIdHeader(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, utils.GenericInternalServerErrorMessage)
			return
		}
		if header.TransactionId != "" {
			c.Next()
		} else {
			id := uuid.NewString()
			c.Header("X-Request-ID", id)
			c.Request.Header.Add("X-Request-ID", id)
			c.Next()
		}
	}
}
