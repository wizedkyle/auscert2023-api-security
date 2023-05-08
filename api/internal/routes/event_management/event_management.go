package event_management

import (
	"github.com/ForgeResponse/ForgeResponse/v2/internal/middleware"
	"github.com/ForgeResponse/ForgeResponse/v2/pkg/model/auditmodel"
	"github.com/ForgeResponse/ForgeResponse/v2/pkg/model/authmodel"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GenerateEventManagementRoutes(router *gin.Engine) {
	eventManagementRoutes := router.Group("/v1/events")
	eventManagementRoutes.Use(middleware.ValidateQueryStringAuth(), middleware.ValidateTransactionId())
	{
		eventManagementRoutes.GET("", middleware.ValidateScopes(authmodel.ReadEvents), listEvents)
	}
}

func listEvents(c *gin.Context) {
	c.JSON(http.StatusOK, auditmodel.AuditResponse{Events: auditmodel.AuditScopes})
}
