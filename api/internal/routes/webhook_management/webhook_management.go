package webhook_management

import (
	"github.com/ForgeResponse/ForgeResponse/v2/internal/database"
	"github.com/ForgeResponse/ForgeResponse/v2/internal/middleware"
	"github.com/ForgeResponse/ForgeResponse/v2/internal/utils"
	"github.com/ForgeResponse/ForgeResponse/v2/pkg/model/auditmodel"
	"github.com/ForgeResponse/ForgeResponse/v2/pkg/model/authmodel"
	"github.com/ForgeResponse/ForgeResponse/v2/pkg/model/webhookmodel"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/url"
)

var encryptionKey = []byte("passphrasewhichneedstobe32bytes!")

func GenerateWebhookManagementRoutes(router *gin.Engine) {
	webhookManagementRoutes := router.Group("/v1/webhooks")
	webhookManagementRoutes.Use(middleware.ValidateAuth(), middleware.ValidateTransactionId())
	{
		webhookManagementRoutes.DELETE("/:id", middleware.ValidateScopes(authmodel.DeleteWebhooks), deleteWebhook)
		webhookManagementRoutes.GET("", middleware.ValidateScopes(authmodel.ReadWebhooks), listWebhooks)
		webhookManagementRoutes.GET("/:id", middleware.ValidateScopes(authmodel.ReadWebhooks), getWebhook)
		webhookManagementRoutes.POST("", middleware.ValidateScopes(authmodel.WriteWebhooks), createWebhook)
		webhookManagementRoutes.POST("/:id/rotate", middleware.ValidateScopes(authmodel.ModifyWebhooks), rotateWebhookSecret)
		webhookManagementRoutes.PUT("/:id", middleware.ValidateScopes(authmodel.ModifyWebhooks), updateWebhook)
	}
}

func createWebhook(c *gin.Context) {
	var json webhookmodel.CreateWebhookRequest
	transactionId, _ := utils.GetTransactionIdHeader(c)
	tenantId := c.GetString("tenantId")
	if err := c.ShouldBindJSON(&json); err != nil {
		utilErr := utils.GenerateError(auditmodel.WebhookCreated, utils.InvalidRequestBody, http.StatusBadRequest, transactionId.TransactionId, err)
		utils.WriteErrorLog(utilErr)
		c.AbortWithStatusJSON(utilErr.ExternalError.Code, utilErr.ExternalError)
		return
	}
	_, err := url.ParseRequestURI(json.Url)
	if err != nil {
		utilErr := utils.GenerateError(auditmodel.WebhookCreated, utils.InvalidRequestBody, http.StatusBadRequest, transactionId.TransactionId, err)
		utils.WriteErrorLog(utilErr)
		c.AbortWithStatusJSON(utilErr.ExternalError.Code, utilErr.ExternalError)
		return
	}
	if len(json.Events) > len(auditmodel.AuditScopes) {
		utilErr := utils.GenerateError(auditmodel.WebhookCreated, utils.InvalidNumberOfEvents, http.StatusBadRequest, transactionId.TransactionId, nil)
		utils.WriteErrorLog(utilErr)
		c.AbortWithStatusJSON(utilErr.ExternalError.Code, utilErr.ExternalError)
		return
	}
	for _, event := range json.Events {
		if !utils.ValidateEventScopes(event) {
			utilErr := utils.GenerateError(auditmodel.WebhookCreated, event+utils.InvalidEvent, http.StatusBadRequest, transactionId.TransactionId, nil)
			utils.WriteErrorLog(utilErr)
			c.AbortWithStatusJSON(utilErr.ExternalError.Code, utilErr.ExternalError)
			return
		}
	}
	secret, err := utils.RandomString(44)
	if err != nil {
		utilErr := utils.GenerateError(auditmodel.WebhookCreated, utils.GenericInternalServerErrorMessage, http.StatusInternalServerError, transactionId.TransactionId, err)
		utils.WriteErrorLog(utilErr)
		c.AbortWithStatusJSON(utilErr.ExternalError.Code, utilErr.ExternalError)
		return
	}
	encryptedSecret, err := utils.EncryptData([]byte(secret), encryptionKey)
	if err != nil {
		utilErr := utils.GenerateError(auditmodel.WebhookCreated, utils.GenericInternalServerErrorMessage, http.StatusInternalServerError, transactionId.TransactionId, err)
		utils.WriteErrorLog(utilErr)
		c.AbortWithStatusJSON(utilErr.ExternalError.Code, utilErr.ExternalError)
		return
	}
	encryptedUrl, err := utils.EncryptData([]byte(json.Url), encryptionKey)
	if err != nil {
		utilErr := utils.GenerateError(auditmodel.WebhookCreated, utils.GenericInternalServerErrorMessage, http.StatusInternalServerError, transactionId.TransactionId, err)
		utils.WriteErrorLog(utilErr)
		c.AbortWithStatusJSON(utilErr.ExternalError.Code, utilErr.ExternalError)
		return
	}
	result, err := database.Client.CreateWebhook(tenantId, webhookmodel.Webhook{
		Algorithm:   "sha256",
		Description: json.Description,
		Events:      json.Events,
		Secret:      encryptedSecret,
		TenantId:    tenantId,
		Url:         encryptedUrl,
		Version:     1,
	})
	if err != nil {
		utilErr := utils.GenerateError(auditmodel.WebhookCreated, utils.GenericInternalServerErrorMessage, http.StatusInternalServerError, transactionId.TransactionId, err)
		utils.WriteErrorLog(utilErr)
		c.AbortWithStatusJSON(utilErr.ExternalError.Code, utilErr.ExternalError)
		return
	}
	c.JSON(http.StatusOK, webhookmodel.WebhookResponse{
		Id:          result.Id,
		Algorithm:   "sha256",
		Description: result.Description,
		Events:      result.Events,
		Secret:      secret,
		Url:         json.Url,
	})
}

func deleteWebhook(c *gin.Context) {
	transactionId, _ := utils.GetTransactionIdHeader(c)
	tenantId := c.GetString("tenantId")
	id := c.Param("id")
	err := database.Client.DeleteWebhook(id, tenantId)
	if err == utils.ErrWebhookNotFound {
		utilErr := utils.GenerateError(auditmodel.WebhookDeleted, utils.WebhookNotFound, http.StatusNotFound, transactionId.TransactionId, err)
		utils.WriteErrorLog(utilErr)
		c.AbortWithStatusJSON(utilErr.ExternalError.Code, utilErr.ExternalError)
		return
	} else if err != nil {
		utilErr := utils.GenerateError(auditmodel.WebhookDeleted, utils.GenericInternalServerErrorMessage, http.StatusInternalServerError, transactionId.TransactionId, err)
		utils.WriteErrorLog(utilErr)
		c.AbortWithStatusJSON(utilErr.ExternalError.Code, utilErr.ExternalError)
		return
	}
	c.Status(http.StatusNoContent)
}

func getWebhook(c *gin.Context) {
	transactionId, _ := utils.GetTransactionIdHeader(c)
	tenantId := c.GetString("tenantId")
	id := c.Param("id")
	result, err := database.Client.GetWebhook(id, tenantId)
	if err == utils.ErrWebhookNotFound {
		utilErr := utils.GenerateError(auditmodel.WebhookRetrieved, utils.WebhookNotFound, http.StatusNotFound, transactionId.TransactionId, err)
		utils.WriteErrorLog(utilErr)
		c.AbortWithStatusJSON(utilErr.ExternalError.Code, utilErr.ExternalError)
		return
	} else if err != nil {
		utilErr := utils.GenerateError(auditmodel.WebhookRetrieved, utils.GenericInternalServerErrorMessage, http.StatusInternalServerError, transactionId.TransactionId, err)
		utils.WriteErrorLog(utilErr)
		c.AbortWithStatusJSON(utilErr.ExternalError.Code, utilErr.ExternalError)
		return
	}
	decryptedSecret, err := utils.DecryptData(result.Secret, encryptionKey)
	if err != nil {
		utilErr := utils.GenerateError(auditmodel.WebhookRetrieved, utils.GenericInternalServerErrorMessage, http.StatusInternalServerError, transactionId.TransactionId, err)
		utils.WriteErrorLog(utilErr)
		c.AbortWithStatusJSON(utilErr.ExternalError.Code, utilErr.ExternalError)
		return
	}
	decryptedUrl, err := utils.DecryptData(result.Url, encryptionKey)
	if err != nil {
		utilErr := utils.GenerateError(auditmodel.WebhookRetrieved, utils.GenericInternalServerErrorMessage, http.StatusInternalServerError, transactionId.TransactionId, err)
		utils.WriteErrorLog(utilErr)
		c.AbortWithStatusJSON(utilErr.ExternalError.Code, utilErr.ExternalError)
		return
	}
	result.Secret = string(decryptedSecret)
	result.Url = string(decryptedUrl)
	c.JSON(http.StatusOK, result)
}

func listWebhooks(c *gin.Context) {
	transactionId, _ := utils.GetTransactionIdHeader(c)
	tenantId := c.GetString("tenantId")
	limit := utils.QueryLimit(c)
	results, err := database.Client.ListWebhooks(limit, tenantId)
	if err == utils.ErrWebhooksNotFound {
		utilErr := utils.GenerateError(auditmodel.WebhookRetrieved, utils.WebhookNotFound, http.StatusNotFound, transactionId.TransactionId, err)
		utils.WriteErrorLog(utilErr)
		c.AbortWithStatusJSON(utilErr.ExternalError.Code, utilErr.ExternalError)
		return
	} else if err != nil {
		utilErr := utils.GenerateError(auditmodel.WebhookRetrieved, utils.GenericInternalServerErrorMessage, http.StatusInternalServerError, transactionId.TransactionId, err)
		utils.WriteErrorLog(utilErr)
		c.AbortWithStatusJSON(utilErr.ExternalError.Code, utilErr.ExternalError)
		return
	}
	for i, result := range results {
		decryptedSecret, err := utils.DecryptData(result.Secret, encryptionKey)
		if err != nil {
			utilErr := utils.GenerateError(auditmodel.WebhookRetrieved, utils.GenericInternalServerErrorMessage, http.StatusInternalServerError, transactionId.TransactionId, err)
			utils.WriteErrorLog(utilErr)
			c.AbortWithStatusJSON(utilErr.ExternalError.Code, utilErr.ExternalError)
			return
		}
		decryptedUrl, err := utils.DecryptData(result.Url, encryptionKey)
		if err != nil {
			utilErr := utils.GenerateError(auditmodel.WebhookRetrieved, utils.GenericInternalServerErrorMessage, http.StatusInternalServerError, transactionId.TransactionId, err)
			utils.WriteErrorLog(utilErr)
			c.AbortWithStatusJSON(utilErr.ExternalError.Code, utilErr.ExternalError)
			return
		}
		results[i].Secret = string(decryptedSecret)
		results[i].Url = string(decryptedUrl)
	}
	c.JSON(http.StatusOK, results)
}

func rotateWebhookSecret(c *gin.Context) {
	transactionId, _ := utils.GetTransactionIdHeader(c)
	tenantId := c.GetString("tenantId")
	id := c.Param("id")
	_, err := database.Client.GetWebhook(id, tenantId)
	if err == utils.ErrWebhookNotFound {
		utilErr := utils.GenerateError(auditmodel.WebhookRotated, utils.WebhookNotFound, http.StatusNotFound, transactionId.TransactionId, err)
		utils.WriteErrorLog(utilErr)
		c.AbortWithStatusJSON(utilErr.ExternalError.Code, utilErr.ExternalError)
		return
	} else if err != nil {
		utilErr := utils.GenerateError(auditmodel.WebhookRotated, utils.GenericInternalServerErrorMessage, http.StatusInternalServerError, transactionId.TransactionId, err)
		utils.WriteErrorLog(utilErr)
		c.AbortWithStatusJSON(utilErr.ExternalError.Code, utilErr.ExternalError)
		return
	}
	secret, err := utils.RandomString(44)
	if err != nil {
		utilErr := utils.GenerateError(auditmodel.WebhookRotated, utils.GenericInternalServerErrorMessage, http.StatusInternalServerError, transactionId.TransactionId, err)
		utils.WriteErrorLog(utilErr)
		c.AbortWithStatusJSON(utilErr.ExternalError.Code, utilErr.ExternalError)
		return
	}
	encryptedSecret, err := utils.EncryptData([]byte(secret), encryptionKey)
	if err != nil {
		utilErr := utils.GenerateError(auditmodel.WebhookRotated, utils.GenericInternalServerErrorMessage, http.StatusInternalServerError, transactionId.TransactionId, err)
		utils.WriteErrorLog(utilErr)
		c.AbortWithStatusJSON(utilErr.ExternalError.Code, utilErr.ExternalError)
		return
	}
	result, err := database.Client.RotateWebhookSecret(id, tenantId, string(encryptedSecret))
	if err == utils.ErrWebhookNotFound {
		utilErr := utils.GenerateError(auditmodel.WebhookRotated, utils.WebhookNotFound, http.StatusNotFound, transactionId.TransactionId, err)
		utils.WriteErrorLog(utilErr)
		c.AbortWithStatusJSON(utilErr.ExternalError.Code, utilErr.ExternalError)
		return
	} else if err != nil {
		utilErr := utils.GenerateError(auditmodel.WebhookRotated, utils.GenericInternalServerErrorMessage, http.StatusInternalServerError, transactionId.TransactionId, err)
		utils.WriteErrorLog(utilErr)
		c.AbortWithStatusJSON(utilErr.ExternalError.Code, utilErr.ExternalError)
		return
	}
	decryptedUrl, err := utils.DecryptData(result.Url, encryptionKey)
	if err != nil {
		utilErr := utils.GenerateError(auditmodel.WebhookUpdated, utils.GenericInternalServerErrorMessage, http.StatusInternalServerError, transactionId.TransactionId, err)
		utils.WriteErrorLog(utilErr)
		c.AbortWithStatusJSON(utilErr.ExternalError.Code, utilErr.ExternalError)
		return
	}
	result.Secret = secret
	result.Url = string(decryptedUrl)
	c.JSON(http.StatusOK, result)
}

func updateWebhook(c *gin.Context) {
	var json webhookmodel.UpdateWebhookRequest
	transactionId, _ := utils.GetTransactionIdHeader(c)
	tenantId := c.GetString("tenantId")
	id := c.Param("id")
	if err := c.ShouldBindJSON(&json); err != nil {
		utilErr := utils.GenerateError(auditmodel.WebhookUpdated, utils.InvalidRequestBody, http.StatusBadRequest, transactionId.TransactionId, err)
		utils.WriteErrorLog(utilErr)
		c.AbortWithStatusJSON(utilErr.ExternalError.Code, utilErr.ExternalError)
		return
	}
	_, err := url.ParseRequestURI(json.Url)
	if err != nil {
		utilErr := utils.GenerateError(auditmodel.WebhookUpdated, utils.InvalidRequestBody, http.StatusBadRequest, transactionId.TransactionId, err)
		utils.WriteErrorLog(utilErr)
		c.AbortWithStatusJSON(utilErr.ExternalError.Code, utilErr.ExternalError)
		return
	}
	if len(json.Events) > len(auditmodel.AuditScopes) {
		utilErr := utils.GenerateError(auditmodel.WebhookUpdated, utils.InvalidNumberOfEvents, http.StatusBadRequest, transactionId.TransactionId, nil)
		utils.WriteErrorLog(utilErr)
		c.AbortWithStatusJSON(utilErr.ExternalError.Code, utilErr.ExternalError)
		return
	}
	for _, event := range json.Events {
		if !utils.ValidateEventScopes(event) {
			utilErr := utils.GenerateError(auditmodel.WebhookUpdated, event+utils.InvalidEvent, http.StatusBadRequest, transactionId.TransactionId, nil)
			utils.WriteErrorLog(utilErr)
			c.AbortWithStatusJSON(utilErr.ExternalError.Code, utilErr.ExternalError)
			return
		}
	}
	encryptedUrl, err := utils.EncryptData([]byte(json.Url), encryptionKey)
	if err != nil {
		utilErr := utils.GenerateError(auditmodel.WebhookUpdated, utils.GenericInternalServerErrorMessage, http.StatusInternalServerError, transactionId.TransactionId, err)
		utils.WriteErrorLog(utilErr)
		c.AbortWithStatusJSON(utilErr.ExternalError.Code, utilErr.ExternalError)
		return
	}
	json.Url = encryptedUrl
	result, err := database.Client.UpdateWebhook(id, tenantId, json)
	if err == utils.ErrWebhookNotFound {
		utilErr := utils.GenerateError(auditmodel.WebhookUpdated, utils.WebhookNotFound, http.StatusNotFound, transactionId.TransactionId, err)
		utils.WriteErrorLog(utilErr)
		c.AbortWithStatusJSON(utilErr.ExternalError.Code, utilErr.ExternalError)
		return
	} else if err != nil {
		utilErr := utils.GenerateError(auditmodel.WebhookUpdated, utils.GenericInternalServerErrorMessage, http.StatusInternalServerError, transactionId.TransactionId, err)
		utils.WriteErrorLog(utilErr)
		c.AbortWithStatusJSON(utilErr.ExternalError.Code, utilErr.ExternalError)
		return
	}
	decryptedSecret, err := utils.DecryptData(result.Secret, encryptionKey)
	if err != nil {
		utilErr := utils.GenerateError(auditmodel.WebhookUpdated, utils.GenericInternalServerErrorMessage, http.StatusInternalServerError, transactionId.TransactionId, err)
		utils.WriteErrorLog(utilErr)
		c.AbortWithStatusJSON(utilErr.ExternalError.Code, utilErr.ExternalError)
		return
	}
	decryptedUrl, err := utils.DecryptData(result.Url, encryptionKey)
	if err != nil {
		utilErr := utils.GenerateError(auditmodel.WebhookUpdated, utils.GenericInternalServerErrorMessage, http.StatusInternalServerError, transactionId.TransactionId, err)
		utils.WriteErrorLog(utilErr)
		c.AbortWithStatusJSON(utilErr.ExternalError.Code, utilErr.ExternalError)
		return
	}
	result.Secret = string(decryptedSecret)
	result.Url = string(decryptedUrl)
	c.JSON(http.StatusOK, result)
}
