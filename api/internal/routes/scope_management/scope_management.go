package scope_management

import (
	"github.com/ForgeResponse/ForgeResponse/v2/internal/middleware"
	"github.com/ForgeResponse/ForgeResponse/v2/pkg/model/authmodel"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GenerateScopeManagementRoutes(router *gin.Engine) {
	scopeManagementRoutes := router.Group("/v1/scopes")
	scopeManagementRoutes.Use(middleware.ValidateAuth(), middleware.ValidateTransactionId())
	{
		scopeManagementRoutes.GET("", middleware.ValidateScopes(authmodel.ReadScopes), listScopes)
	}
}

func listScopes(c *gin.Context) {
	c.JSON(http.StatusOK, authmodel.ScopeResponse{Scopes: authmodel.Scopes})
}
