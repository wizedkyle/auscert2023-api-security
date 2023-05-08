package investigation_management

import (
	"github.com/ForgeResponse/ForgeResponse/v2/internal/database"
	"github.com/ForgeResponse/ForgeResponse/v2/internal/middleware"
	"github.com/ForgeResponse/ForgeResponse/v2/internal/utils"
	"github.com/ForgeResponse/ForgeResponse/v2/pkg/model/auditmodel"
	"github.com/ForgeResponse/ForgeResponse/v2/pkg/model/authmodel"
	"github.com/ForgeResponse/ForgeResponse/v2/pkg/model/investigationmodel"
	"github.com/ForgeResponse/ForgeResponse/v2/pkg/model/investigationtemplatemodel"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func GenerateInvestigationManagementRoutes(router *gin.Engine) {
	investigationManagementRoutes := router.Group("/v1/investigations")
	investigationManagementRoutes.Use(middleware.ValidateAuth(), middleware.ValidateTransactionId())
	{
		// Investigation routes
		investigationManagementRoutes.DELETE("/:id", middleware.ValidateScopes(authmodel.DeleteInvestigations), deleteInvestigation)
		investigationManagementRoutes.GET("", middleware.ValidateScopes(authmodel.ReadInvestigations), listInvestigations)
		investigationManagementRoutes.GET("/:id", middleware.ValidateScopes(authmodel.ReadInvestigations), getInvestigation)
		investigationManagementRoutes.POST("", middleware.ValidateScopes(authmodel.WriteInvestigations), createInvestigation)
		investigationManagementRoutes.PUT("/:id", middleware.ValidateScopes(authmodel.ModifyInvestigations), updateInvestigation)

		// Template routes
		investigationManagementRoutes.DELETE("/template/:id", middleware.ValidateScopes(authmodel.DeleteInvestigationTemplates), deleteInvestigationTemplate)
		investigationManagementRoutes.GET("/template", middleware.ValidateScopes(authmodel.ReadInvestigationTemplates), listInvestigationTemplates)
		investigationManagementRoutes.GET("/template/:id", middleware.ValidateScopes(authmodel.ReadInvestigationTemplates), getInvestigationTemplate)
		investigationManagementRoutes.POST("/template", middleware.ValidateScopes(authmodel.WriteInvestigationTemplates), createInvestigationTemplate)
		investigationManagementRoutes.PUT("/template/:id", middleware.ValidateScopes(authmodel.ModifyInvestigationTemplates), updateInvestigationTemplate)
	}
}

func createInvestigation(c *gin.Context) {
	var json investigationmodel.CreateInvestigationRequest
	transactionId, _ := utils.GetTransactionIdHeader(c)
	tenantId := c.GetString("tenantId")
	if err := c.ShouldBindJSON(&json); err != nil {
		utilErr := utils.GenerateError(auditmodel.InvestigationCreated, utils.InvalidRequestBody, http.StatusBadRequest, transactionId.TransactionId, err)
		utils.WriteErrorLog(utilErr)
		c.AbortWithStatusJSON(utilErr.ExternalError.Code, utilErr.ExternalError)
		return
	}
	if json.TemplateId != "" {
		templateResult, err := database.Client.GetInvestigationTemplate(json.TemplateId, tenantId)
		if err == utils.ErrInvestigationTemplateNotFound {
			utilErr := utils.GenerateError(auditmodel.InvestigationCreated, utils.InvestigationTemplateNotFound, http.StatusNotFound, transactionId.TransactionId, nil)
			utils.WriteErrorLog(utilErr)
			c.AbortWithStatusJSON(utilErr.ExternalError.Code, utilErr.ExternalError)
			return
		} else if err != nil {
			utilErr := utils.GenerateError(auditmodel.InvestigationRetrieved, utils.GenericInternalServerErrorMessage, http.StatusInternalServerError, transactionId.TransactionId, err)
			utils.WriteErrorLog(utilErr)
			c.AbortWithStatusJSON(utilErr.ExternalError.Code, utilErr.ExternalError)
			return
		}
		result, err := database.Client.CreateInvestigation(tenantId, investigationmodel.Investigation{
			AssignedTo:      json.AssignedTo,
			CreatedAt:       time.Now(),
			Description:     templateResult.Description,
			InvestigationId: "",
			Severity:        templateResult.Severity,
			Status:          templateResult.Status,
			Tags:            templateResult.Tags,
			TenantId:        tenantId,
			Title:           templateResult.TitlePrefix + " " + json.Title,
			Tlp:             templateResult.Tlp,
			Version:         1,
		})
		if err != nil {
			utilErr := utils.GenerateError(auditmodel.InvestigationCreated, utils.GenericInternalServerErrorMessage, http.StatusInternalServerError, transactionId.TransactionId, err)
			utils.WriteErrorLog(utilErr)
			c.AbortWithStatusJSON(utilErr.ExternalError.Code, utilErr.ExternalError)
			return
		}
		c.JSON(http.StatusOK, result)
	} else {
		result, err := database.Client.CreateInvestigation(tenantId, investigationmodel.Investigation{
			AssignedTo:      json.AssignedTo,
			CreatedAt:       time.Now(),
			Description:     json.Description,
			InvestigationId: "",
			Severity:        json.Severity,
			Status:          json.Status,
			Tags:            json.Tags,
			TenantId:        tenantId,
			Title:           json.Title,
			Tlp:             json.Tlp,
			Version:         1,
		})
		if err != nil {
			utilErr := utils.GenerateError(auditmodel.InvestigationCreated, utils.GenericInternalServerErrorMessage, http.StatusInternalServerError, transactionId.TransactionId, err)
			utils.WriteErrorLog(utilErr)
			c.AbortWithStatusJSON(utilErr.ExternalError.Code, utilErr.ExternalError)
			return
		}
		c.JSON(http.StatusOK, result)
	}
}

func deleteInvestigation(c *gin.Context) {
	transactionId, _ := utils.GetTransactionIdHeader(c)
	tenantId := c.GetString("tenantId")
	id := c.Param("id")
	err := database.Client.DeleteInvestigation(id, tenantId)
	if err == utils.ErrInvestigationNotFound {
		utilErr := utils.GenerateError(auditmodel.InvestigationDeleted, utils.InvestigationNotFound, http.StatusNotFound, transactionId.TransactionId, nil)
		utils.WriteErrorLog(utilErr)
		c.AbortWithStatusJSON(utilErr.ExternalError.Code, utilErr.ExternalError)
		return
	} else if err != nil {
		utilErr := utils.GenerateError(auditmodel.InvestigationDeleted, utils.GenericInternalServerErrorMessage, http.StatusInternalServerError, transactionId.TransactionId, err)
		utils.WriteErrorLog(utilErr)
		c.AbortWithStatusJSON(utilErr.ExternalError.Code, utilErr.ExternalError)
		return
	}
	c.Status(http.StatusNoContent)
}

func getInvestigation(c *gin.Context) {
	transactionId, _ := utils.GetTransactionIdHeader(c)
	tenantId := c.GetString("tenantId")
	id := c.Param("id")
	result, err := database.Client.GetInvestigation(id, tenantId)
	if err == utils.ErrInvestigationNotFound {
		utilErr := utils.GenerateError(auditmodel.InvestigationRetrieved, utils.InvestigationNotFound, http.StatusNotFound, transactionId.TransactionId, nil)
		utils.WriteErrorLog(utilErr)
		c.AbortWithStatusJSON(utilErr.ExternalError.Code, utilErr.ExternalError)
		return
	} else if err != nil {
		utilErr := utils.GenerateError(auditmodel.InvestigationRetrieved, utils.GenericInternalServerErrorMessage, http.StatusInternalServerError, transactionId.TransactionId, err)
		utils.WriteErrorLog(utilErr)
		c.AbortWithStatusJSON(utilErr.ExternalError.Code, utilErr.ExternalError)
		return
	}
	c.JSON(http.StatusOK, result)
}

func listInvestigations(c *gin.Context) {
	transactionId, _ := utils.GetTransactionIdHeader(c)
	tenantId := c.GetString("tenantId")
	limit := utils.QueryLimit(c)
	results, err := database.Client.ListInvestigations(limit, tenantId)
	if err == utils.ErrInvestigationsNotFound {
		utilErr := utils.GenerateError(auditmodel.InvestigationRetrieved, utils.InvestigationsNotFound, http.StatusNotFound, transactionId.TransactionId, nil)
		utils.WriteErrorLog(utilErr)
		c.AbortWithStatusJSON(utilErr.ExternalError.Code, utilErr.ExternalError)
		return
	} else if err != nil {
		utilErr := utils.GenerateError(auditmodel.InvestigationRetrieved, utils.GenericInternalServerErrorMessage, http.StatusInternalServerError, transactionId.TransactionId, err)
		utils.WriteErrorLog(utilErr)
		c.AbortWithStatusJSON(utilErr.ExternalError.Code, utilErr.ExternalError)
		return
	}
	c.JSON(http.StatusOK, results)
}

func updateInvestigation(c *gin.Context) {
	var json investigationmodel.UpdateInvestigationRequest
	transactionId, _ := utils.GetTransactionIdHeader(c)
	tenantId := c.GetString("tenantId")
	id := c.Param("id")
	if err := c.ShouldBindJSON(&json); err != nil {
		utilErr := utils.GenerateError(auditmodel.InvestigationUpdated, utils.InvalidRequestBody, http.StatusBadRequest, transactionId.TransactionId, err)
		utils.WriteErrorLog(utilErr)
		c.AbortWithStatusJSON(utilErr.ExternalError.Code, utilErr.ExternalError)
		return
	}
	result, err := database.Client.UpdateInvestigation(id, tenantId, json)
	if err == utils.ErrInvestigationNotFound {
		utilErr := utils.GenerateError(auditmodel.InvestigationUpdated, utils.InvestigationNotFound, http.StatusNotFound, transactionId.TransactionId, err)
		utils.WriteErrorLog(utilErr)
		c.AbortWithStatusJSON(utilErr.ExternalError.Code, utilErr.ExternalError)
		return
	} else if err != nil {
		utilErr := utils.GenerateError(auditmodel.InvestigationUpdated, utils.GenericInternalServerErrorMessage, http.StatusInternalServerError, transactionId.TransactionId, err)
		utils.WriteErrorLog(utilErr)
		c.AbortWithStatusJSON(utilErr.ExternalError.Code, utilErr.ExternalError)
		return
	}
	c.JSON(http.StatusOK, result)
}

func createInvestigationTemplate(c *gin.Context) {
	var json investigationtemplatemodel.UpdateInvestigationRequest
	transactionId, _ := utils.GetTransactionIdHeader(c)
	tenantId := c.GetString("tenantId")
	if err := c.ShouldBindJSON(&json); err != nil {
		utilErr := utils.GenerateError(auditmodel.InvestigationTemplateCreated, utils.InvalidRequestBody, http.StatusBadRequest, transactionId.TransactionId, err)
		utils.WriteErrorLog(utilErr)
		c.AbortWithStatusJSON(utilErr.ExternalError.Code, utilErr.ExternalError)
		return
	}
	result, err := database.Client.CreateInvestigationTemplate(tenantId, investigationtemplatemodel.InvestigationTemplate{
		TenantId:    tenantId,
		CreatedAt:   time.Now(),
		Description: json.Description,
		TitlePrefix: json.TitlePrefix,
		Severity:    json.Severity,
		Status:      json.Status,
		Tags:        json.Tags,
		Tlp:         json.Tlp,
	})
	if err != nil {
		utilErr := utils.GenerateError(auditmodel.InvestigationTemplateCreated, utils.GenericInternalServerErrorMessage, http.StatusInternalServerError, transactionId.TransactionId, err)
		utils.WriteErrorLog(utilErr)
		c.AbortWithStatusJSON(utilErr.ExternalError.Code, utilErr.ExternalError)
		return
	}
	c.JSON(http.StatusOK, result)
}

func deleteInvestigationTemplate(c *gin.Context) {
	transactionId, _ := utils.GetTransactionIdHeader(c)
	tenantId := c.GetString("tenantId")
	id := c.Param("id")
	err := database.Client.DeleteInvestigationTemplate(id, tenantId)
	if err == utils.ErrInvestigationTemplateNotFound {
		utilErr := utils.GenerateError(auditmodel.InvestigationTemplateDeleted, utils.InvestigationTemplateNotFound, http.StatusNotFound, transactionId.TransactionId, nil)
		utils.WriteErrorLog(utilErr)
		c.AbortWithStatusJSON(utilErr.ExternalError.Code, utilErr.ExternalError)
		return
	} else if err != nil {
		utilErr := utils.GenerateError(auditmodel.InvestigationTemplateDeleted, utils.GenericInternalServerErrorMessage, http.StatusInternalServerError, transactionId.TransactionId, err)
		utils.WriteErrorLog(utilErr)
		c.AbortWithStatusJSON(utilErr.ExternalError.Code, utilErr.ExternalError)
		return
	}
	c.Status(http.StatusNoContent)
}

func getInvestigationTemplate(c *gin.Context) {
	transactionId, _ := utils.GetTransactionIdHeader(c)
	tenantId := c.GetString("tenantId")
	id := c.Param("id")
	result, err := database.Client.GetInvestigationTemplate(id, tenantId)
	if err == utils.ErrInvestigationTemplateNotFound {
		utilErr := utils.GenerateError(auditmodel.InvestigationTemplateRetrieved, utils.InvestigationTemplateNotFound, http.StatusNotFound, transactionId.TransactionId, nil)
		utils.WriteErrorLog(utilErr)
		c.AbortWithStatusJSON(utilErr.ExternalError.Code, utilErr.ExternalError)
		return
	} else if err != nil {
		utilErr := utils.GenerateError(auditmodel.InvestigationTemplateRetrieved, utils.GenericInternalServerErrorMessage, http.StatusInternalServerError, transactionId.TransactionId, err)
		utils.WriteErrorLog(utilErr)
		c.AbortWithStatusJSON(utilErr.ExternalError.Code, utilErr.ExternalError)
		return
	}
	c.JSON(http.StatusOK, result)
}

func listInvestigationTemplates(c *gin.Context) {
	transactionId, _ := utils.GetTransactionIdHeader(c)
	tenantId := c.GetString("tenantId")
	limit := utils.QueryLimit(c)
	results, err := database.Client.ListInvestigationTemplates(limit, tenantId)
	if err == utils.ErrInvestigationTemplatesNotFound {
		utilErr := utils.GenerateError(auditmodel.InvestigationTemplateRetrieved, utils.InvestigationTemplatesNotFound, http.StatusNotFound, transactionId.TransactionId, nil)
		utils.WriteErrorLog(utilErr)
		c.AbortWithStatusJSON(utilErr.ExternalError.Code, utilErr.ExternalError)
		return
	} else if err != nil {
		utilErr := utils.GenerateError(auditmodel.InvestigationTemplateRetrieved, utils.GenericInternalServerErrorMessage, http.StatusInternalServerError, transactionId.TransactionId, err)
		utils.WriteErrorLog(utilErr)
		c.AbortWithStatusJSON(utilErr.ExternalError.Code, utilErr.ExternalError)
		return
	}
	c.JSON(http.StatusOK, results)
}

func updateInvestigationTemplate(c *gin.Context) {
	var json investigationtemplatemodel.UpdateInvestigationRequest
	transactionId, _ := utils.GetTransactionIdHeader(c)
	tenantId := c.GetString("tenantId")
	id := c.Param("id")
	if err := c.ShouldBindJSON(&json); err != nil {
		utilErr := utils.GenerateError(auditmodel.InvestigationTemplateUpdated, utils.InvalidRequestBody, http.StatusBadRequest, transactionId.TransactionId, err)
		utils.WriteErrorLog(utilErr)
		c.AbortWithStatusJSON(utilErr.ExternalError.Code, utilErr.ExternalError)
		return
	}
	result, err := database.Client.UpdateInvestigationTemplate(id, tenantId, json)
	if err == utils.ErrInvestigationTemplateNotFound {
		utilErr := utils.GenerateError(auditmodel.InvestigationTemplateUpdated, utils.InvestigationTemplateNotFound, http.StatusNotFound, transactionId.TransactionId, err)
		utils.WriteErrorLog(utilErr)
		c.AbortWithStatusJSON(utilErr.ExternalError.Code, utilErr.ExternalError)
		return
	} else if err != nil {
		utilErr := utils.GenerateError(auditmodel.InvestigationTemplateUpdated, utils.GenericInternalServerErrorMessage, http.StatusInternalServerError, transactionId.TransactionId, err)
		utils.WriteErrorLog(utilErr)
		c.AbortWithStatusJSON(utilErr.ExternalError.Code, utilErr.ExternalError)
		return
	}
	c.JSON(http.StatusOK, result)
}
