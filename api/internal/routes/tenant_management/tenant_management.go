package tenant_management

import (
	"github.com/ForgeResponse/ForgeResponse/v2/internal/database"
	"github.com/ForgeResponse/ForgeResponse/v2/internal/middleware"
	"github.com/ForgeResponse/ForgeResponse/v2/internal/utils"
	"github.com/ForgeResponse/ForgeResponse/v2/pkg/model/auditmodel"
	"github.com/ForgeResponse/ForgeResponse/v2/pkg/model/tenantmodel"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GenerateTenantManagementRoutes(router *gin.Engine) {
	tenantManagementRoutes := router.Group("/v1/tenants")
	tenantManagementRoutes.Use(middleware.ValidateTransactionId())
	{
		tenantManagementRoutes.DELETE("/:id", middleware.ValidateAuth(), deleteTenant)
		tenantManagementRoutes.GET("/:id", middleware.ValidateAuth(), getTenant)
		tenantManagementRoutes.POST("", middleware.ValidateHmacAuthWithTimestampValidation(), createTenant)
		tenantManagementRoutes.PUT("/:id", middleware.ValidateHmacAuthWithoutTimestampValidation(), updateTenant)
	}
}

func createTenant(c *gin.Context) {
	var json tenantmodel.CreateTenantRequest
	transactionId, _ := utils.GetTransactionIdHeader(c)
	if err := c.ShouldBindJSON(&json); err != nil {
		utilErr := utils.GenerateError(auditmodel.TenantCreated, utils.InvalidRequestBody, http.StatusBadRequest, transactionId.TransactionId, err)
		utils.WriteErrorLog(utilErr)
		c.AbortWithStatusJSON(utilErr.ExternalError.Code, utilErr.ExternalError)
		return
	}
	result, err := database.Client.CreateTenant(tenantmodel.Tenant{
		Name: json.Name,
	})
	if err != nil {
		utilErr := utils.GenerateError(auditmodel.TenantCreated, utils.GenericInternalServerErrorMessage, http.StatusInternalServerError, transactionId.TransactionId, err)
		utils.WriteErrorLog(utilErr)
		c.AbortWithStatusJSON(utilErr.ExternalError.Code, utilErr.ExternalError)
		return
	}
	c.JSON(http.StatusOK, result)
}

func deleteTenant(c *gin.Context) {
	transactionId, _ := utils.GetTransactionIdHeader(c)
	id := c.Param("id")
	err := database.Client.DeleteTenant(id)
	if err == utils.ErrTenantNotFound {
		utilErr := utils.GenerateError(auditmodel.TenantDeleted, utils.TenantNotFound, http.StatusNotFound, transactionId.TransactionId, nil)
		utils.WriteErrorLog(utilErr)
		c.AbortWithStatusJSON(utilErr.ExternalError.Code, utilErr.ExternalError)
		return
	} else if err != nil {
		utilErr := utils.GenerateError(auditmodel.TenantDeleted, utils.GenericInternalServerErrorMessage, http.StatusInternalServerError, transactionId.TransactionId, err)
		utils.WriteErrorLog(utilErr)
		c.AbortWithStatusJSON(utilErr.ExternalError.Code, utilErr.ExternalError)
		return
	}
	c.Status(http.StatusNoContent)
}

func getTenant(c *gin.Context) {
	transactionId, _ := utils.GetTransactionIdHeader(c)
	id := c.Param("id")
	result, err := database.Client.GetTenant(id)
	if err == utils.ErrTenantNotFound {
		utilErr := utils.GenerateError(auditmodel.TenantRetrieved, utils.TenantNotFound, http.StatusNotFound, transactionId.TransactionId, nil)
		utils.WriteErrorLog(utilErr)
		c.AbortWithStatusJSON(utilErr.ExternalError.Code, utilErr.ExternalError)
		return
	} else if err != nil {
		utilErr := utils.GenerateError(auditmodel.TenantRetrieved, utils.GenericInternalServerErrorMessage, http.StatusInternalServerError, transactionId.TransactionId, err)
		utils.WriteErrorLog(utilErr)
		c.AbortWithStatusJSON(utilErr.ExternalError.Code, utilErr.ExternalError)
		return
	}
	c.JSON(http.StatusOK, result)
}

func updateTenant(c *gin.Context) {
	var json tenantmodel.UpdateTenantRequest
	transactionId, _ := utils.GetTransactionIdHeader(c)
	id := c.Param("id")
	if err := c.ShouldBindJSON(&json); err != nil {
		utilErr := utils.GenerateError(auditmodel.TenantUpdated, utils.InvalidRequestBody, http.StatusBadRequest, transactionId.TransactionId, err)
		utils.WriteErrorLog(utilErr)
		c.AbortWithStatusJSON(utilErr.ExternalError.Code, utilErr.ExternalError)
		return
	}
	result, err := database.Client.UpdateTenant(id, json)
	if err != nil {
		utilErr := utils.GenerateError(auditmodel.TenantUpdated, utils.GenericInternalServerErrorMessage, http.StatusInternalServerError, transactionId.TransactionId, err)
		utils.WriteErrorLog(utilErr)
		c.AbortWithStatusJSON(utilErr.ExternalError.Code, utilErr.ExternalError)
		return
	}
	c.JSON(http.StatusOK, result)
}
