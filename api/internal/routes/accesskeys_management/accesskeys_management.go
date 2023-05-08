package accesskeys_management

import (
	"github.com/ForgeResponse/ForgeResponse/v2/internal/database"
	"github.com/ForgeResponse/ForgeResponse/v2/internal/middleware"
	"github.com/ForgeResponse/ForgeResponse/v2/internal/utils"
	"github.com/ForgeResponse/ForgeResponse/v2/pkg/model/accesskeymodel"
	"github.com/ForgeResponse/ForgeResponse/v2/pkg/model/auditmodel"
	"github.com/ForgeResponse/ForgeResponse/v2/pkg/model/authmodel"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GenerateAccessKeyManagementRoutes(router *gin.Engine) {
	accessKeyManagementRoutes := router.Group("/v1/accesskeys")
	accessKeyManagementRoutes.Use(middleware.ValidateTransactionId())
	{
		accessKeyManagementRoutes.DELETE("/:id", middleware.ValidateAuth(), middleware.ValidateScopes(authmodel.DeleteAccessKeys), deleteAccessKey)
		accessKeyManagementRoutes.GET("", listAccessKeys)
		accessKeyManagementRoutes.GET("/:id", middleware.ValidateAuth(), middleware.ValidateScopes(authmodel.ReadAccessKeys), getAccessKey)
		accessKeyManagementRoutes.POST("", middleware.ValidateAuth(), middleware.ValidateScopes(authmodel.WriteAccessKeys), createAccessKey)
		accessKeyManagementRoutes.POST("/:id/rotate", middleware.ValidateAuth(), middleware.ValidateScopes(authmodel.ModifyAccessKeys), rotateAccessKey)
		accessKeyManagementRoutes.PUT("/:id", middleware.ValidateAuth(), middleware.ValidateScopes(authmodel.ModifyAccessKeys), updateAccessKey)
	}
}

func createAccessKey(c *gin.Context) {
	var json accesskeymodel.CreateAccessKeyRequest
	transactionId, _ := utils.GetTransactionIdHeader(c)
	tenantId := c.GetString("tenantId")
	if err := c.ShouldBindJSON(&json); err != nil {
		utilErr := utils.GenerateError(auditmodel.AccessKeyCreated, utils.InvalidRequestBody, http.StatusBadRequest, transactionId.TransactionId, err)
		utils.WriteErrorLog(utilErr)
		c.AbortWithStatusJSON(utilErr.ExternalError.Code, utilErr.ExternalError)
		return
	}
	if len(json.Scopes) > len(authmodel.Scopes) {
		utilErr := utils.GenerateError(auditmodel.AccessKeyCreated, utils.InvalidNumberOfScopes, http.StatusBadRequest, transactionId.TransactionId, nil)
		utils.WriteErrorLog(utilErr)
		c.AbortWithStatusJSON(utilErr.ExternalError.Code, utilErr.ExternalError)
		return
	}
	for _, scope := range json.Scopes {
		if !utils.ValidateAvailableScopes(scope) {
			utilErr := utils.GenerateError(auditmodel.AccessKeyCreated, scope+utils.InvalidScope, http.StatusBadRequest, transactionId.TransactionId, nil)
			utils.WriteErrorLog(utilErr)
			c.AbortWithStatusJSON(utilErr.ExternalError.Code, utilErr.ExternalError)
			return
		}
	}
	keyPrefix, plainTextKey, keyHash, err := utils.GenerateAccessKey()
	if err != nil {
		utilErr := utils.GenerateError(auditmodel.AccessKeyCreated, utils.GenericInternalServerErrorMessage, http.StatusInternalServerError, transactionId.TransactionId, err)
		utils.WriteErrorLog(utilErr)
		c.AbortWithStatusJSON(utilErr.ExternalError.Code, utilErr.ExternalError)
		return
	}
	result, err := database.Client.CreateAccessKey(tenantId, accesskeymodel.AccessKey{
		Description: json.Description,
		Expiration:  utils.GenerateExpiration(json.Duration),
		KeyHash:     keyHash,
		KeyPrefix:   keyPrefix,
		TenantId:    tenantId,
		Scopes:      json.Scopes,
		Version:     1,
	})
	if err != nil {
		utilErr := utils.GenerateError(auditmodel.AccessKeyCreated, utils.GenericInternalServerErrorMessage, http.StatusInternalServerError, transactionId.TransactionId, nil)
		utils.WriteErrorLog(utilErr)
		c.AbortWithStatusJSON(utilErr.ExternalError.Code, utilErr.ExternalError)
		return
	}
	c.JSON(http.StatusOK, accesskeymodel.CreateAccessKeyResponse{
		Id:          result.Id,
		Description: result.Description,
		Expiration:  result.Expiration,
		Key:         plainTextKey,
		Scopes:      result.Scopes,
	})
}

func deleteAccessKey(c *gin.Context) {
	transactionId, _ := utils.GetTransactionIdHeader(c)
	tenantId := c.GetString("tenantId")
	id := c.Param("id")
	err := database.Client.DeleteAccessKey(id, tenantId)
	if err == utils.ErrAccessKeyNotFound {
		utilErr := utils.GenerateError(auditmodel.AccessKeyDeleted, utils.AccessKeyNotFound, http.StatusNotFound, transactionId.TransactionId, nil)
		utils.WriteErrorLog(utilErr)
		c.AbortWithStatusJSON(utilErr.ExternalError.Code, utilErr.ExternalError)
		return
	} else if err != nil {
		utilErr := utils.GenerateError(auditmodel.AccessKeyDeleted, utils.GenericInternalServerErrorMessage, http.StatusNotFound, transactionId.TransactionId, err)
		utils.WriteErrorLog(utilErr)
		c.AbortWithStatusJSON(utilErr.ExternalError.Code, utilErr.ExternalError)
		return
	}
	c.Status(http.StatusNoContent)
}

func getAccessKey(c *gin.Context) {
	transactionId, _ := utils.GetTransactionIdHeader(c)
	tenantId := c.GetString("tenantId")
	id := c.Param("id")
	result, err := database.Client.GetAccessKey(id, tenantId)
	if err == utils.ErrAccessKeyNotFound {
		utilErr := utils.GenerateError(auditmodel.AccessKeyRetrieved, utils.AccessKeyNotFound, http.StatusNotFound, transactionId.TransactionId, nil)
		utils.WriteErrorLog(utilErr)
		c.AbortWithStatusJSON(utilErr.ExternalError.Code, utilErr.ExternalError)
		return
	} else if err != nil {
		utilErr := utils.GenerateError(auditmodel.AccessKeyRetrieved, utils.GenericInternalServerErrorMessage, http.StatusInternalServerError, transactionId.TransactionId, err)
		utils.WriteErrorLog(utilErr)
		c.AbortWithStatusJSON(utilErr.ExternalError.Code, utilErr.ExternalError)
		return
	}
	c.JSON(http.StatusOK, result)
}

func listAccessKeys(c *gin.Context) {
	transactionId, _ := utils.GetTransactionIdHeader(c)
	tenantId := c.GetString("tenantId")
	limit := utils.QueryLimit(c)
	results, err := database.Client.ListAccessKeys(limit, tenantId)
	if err == utils.ErrAccessKeysNotFound {
		utilErr := utils.GenerateError(auditmodel.AccessKeyRetrieved, utils.AccessKeysNotFound, http.StatusNotFound, transactionId.TransactionId, nil)
		utils.WriteErrorLog(utilErr)
		c.AbortWithStatusJSON(utilErr.ExternalError.Code, utilErr.ExternalError)
		return
	} else if err != nil {
		utilErr := utils.GenerateError(auditmodel.AccessKeyRetrieved, utils.GenericInternalServerErrorMessage, http.StatusInternalServerError, transactionId.TransactionId, err)
		utils.WriteErrorLog(utilErr)
		c.AbortWithStatusJSON(utilErr.ExternalError.Code, utilErr.ExternalError)
		return
	}
	c.JSON(http.StatusOK, results)
}

func rotateAccessKey(c *gin.Context) {
	var json accesskeymodel.RotateAccessKeyRequest
	transactionId, _ := utils.GetTransactionIdHeader(c)
	tenantId := c.GetString("tenantId")
	id := c.Param("id")
	if err := c.ShouldBindJSON(&json); err != nil {
		utilErr := utils.GenerateError(auditmodel.AccessKeyRotated, utils.InvalidRequestBody, http.StatusBadRequest, transactionId.TransactionId, err)
		utils.WriteErrorLog(utilErr)
		c.AbortWithStatusJSON(utilErr.ExternalError.Code, utilErr.ExternalError)
		return
	}
	_, err := database.Client.GetAccessKey(id, tenantId)
	if err == utils.ErrAccessKeyNotFound {
		utilErr := utils.GenerateError(auditmodel.AccessKeyRotated, utils.AccessKeyNotFound, http.StatusNotFound, transactionId.TransactionId, nil)
		utils.WriteErrorLog(utilErr)
		c.AbortWithStatusJSON(utilErr.ExternalError.Code, utilErr.ExternalError)
		return
	} else if err != nil {
		utilErr := utils.GenerateError(auditmodel.AccessKeyRotated, utils.GenericInternalServerErrorMessage, http.StatusInternalServerError, transactionId.TransactionId, err)
		utils.WriteErrorLog(utilErr)
		c.AbortWithStatusJSON(utilErr.ExternalError.Code, utilErr.ExternalError)
		return
	}
	keyPrefix, plainTextKey, keyHash, err := utils.GenerateAccessKey()
	if err != nil {
		utilErr := utils.GenerateError(auditmodel.AccessKeyRotated, utils.GenericInternalServerErrorMessage, http.StatusInternalServerError, transactionId.TransactionId, err)
		utils.WriteErrorLog(utilErr)
		c.AbortWithStatusJSON(utilErr.ExternalError.Code, utilErr.ExternalError)
		return
	}
	result, err := database.Client.RotateAccessKey(id, tenantId, accesskeymodel.AccessKey{
		Expiration: utils.GenerateExpiration(json.Duration),
		KeyHash:    keyHash,
		KeyPrefix:  keyPrefix,
	})
	if err == utils.ErrAccessKeyNotFound {
		utilErr := utils.GenerateError(auditmodel.AccessKeyRotated, utils.AccessKeyNotFound, http.StatusNotFound, transactionId.TransactionId, err)
		utils.WriteErrorLog(utilErr)
		c.AbortWithStatusJSON(utilErr.ExternalError.Code, utilErr.ExternalError)
		return
	} else if err != nil {
		utilErr := utils.GenerateError(auditmodel.AccessKeyRotated, utils.GenericInternalServerErrorMessage, http.StatusInternalServerError, transactionId.TransactionId, err)
		utils.WriteErrorLog(utilErr)
		c.AbortWithStatusJSON(utilErr.ExternalError.Code, utilErr.ExternalError)
		return
	}
	c.JSON(http.StatusOK, accesskeymodel.RotateAccessKeyResponse{
		Id:         result.Id,
		Expiration: result.Expiration,
		Key:        plainTextKey,
	})
}

func updateAccessKey(c *gin.Context) {
	var json accesskeymodel.UpdateAccessKeyRequest
	transactionId, _ := utils.GetTransactionIdHeader(c)
	tenantId := c.GetString("tenantId")
	id := c.Param("id")
	if err := c.ShouldBindJSON(&json); err != nil {
		utilErr := utils.GenerateError(auditmodel.AccessKeyUpdated, utils.InvalidRequestBody, http.StatusBadRequest, transactionId.TransactionId, err)
		utils.WriteErrorLog(utilErr)
		c.AbortWithStatusJSON(utilErr.ExternalError.Code, utilErr.ExternalError)
		return
	}
	if len(json.Scopes) > len(authmodel.Scopes) {
		utilErr := utils.GenerateError(auditmodel.AccessKeyUpdated, utils.InvalidNumberOfScopes, http.StatusBadRequest, transactionId.TransactionId, nil)
		utils.WriteErrorLog(utilErr)
		c.AbortWithStatusJSON(utilErr.ExternalError.Code, utilErr.ExternalError)
		return
	}
	for _, scope := range json.Scopes {
		if !utils.ValidateAvailableScopes(scope) {
			utilErr := utils.GenerateError(auditmodel.AccessKeyUpdated, scope+utils.InvalidScope, http.StatusBadRequest, transactionId.TransactionId, nil)
			utils.WriteErrorLog(utilErr)
			c.AbortWithStatusJSON(utilErr.ExternalError.Code, utilErr.ExternalError)
			return
		}
	}
	result, err := database.Client.UpdateAccessKey(id, tenantId, json)
	if err == utils.ErrAccessKeyNotFound {
		utilErr := utils.GenerateError(auditmodel.AccessKeyUpdated, utils.AccessKeyNotFound, http.StatusNotFound, transactionId.TransactionId, err)
		utils.WriteErrorLog(utilErr)
		c.AbortWithStatusJSON(utilErr.ExternalError.Code, utilErr.ExternalError)
		return
	} else if err != nil {
		utilErr := utils.GenerateError(auditmodel.AccessKeyUpdated, utils.GenericInternalServerErrorMessage, http.StatusInternalServerError, transactionId.TransactionId, err)
		utils.WriteErrorLog(utilErr)
		c.AbortWithStatusJSON(utilErr.ExternalError.Code, utilErr.ExternalError)
		return
	}
	c.JSON(http.StatusOK, result)
}
