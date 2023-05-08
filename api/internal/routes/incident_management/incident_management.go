package incident_management

import (
	"github.com/ForgeResponse/ForgeResponse/v2/internal/database"
	"github.com/ForgeResponse/ForgeResponse/v2/internal/middleware"
	"github.com/ForgeResponse/ForgeResponse/v2/internal/utils"
	"github.com/ForgeResponse/ForgeResponse/v2/pkg/model/auditmodel"
	"github.com/ForgeResponse/ForgeResponse/v2/pkg/model/authmodel"
	"github.com/ForgeResponse/ForgeResponse/v2/pkg/model/incidentcommentmodel"
	"github.com/ForgeResponse/ForgeResponse/v2/pkg/model/incidentmodel"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func GenerateIncidentManagementRoutes(router *gin.Engine) {
	incidentManagementRoutes := router.Group("/v1/incidents")
	incidentManagementRoutes.Use(middleware.ValidateAuth(), middleware.ValidateTransactionId())
	{
		// Incident Routes
		incidentManagementRoutes.DELETE("/:id", middleware.ValidateScopes(authmodel.DeleteIncident), deleteIncident)
		incidentManagementRoutes.GET("", middleware.ValidateScopes(authmodel.ReadIncident), listIncidents)
		incidentManagementRoutes.GET("/:id", middleware.ValidateScopes(authmodel.ReadIncident), getIncident)
		incidentManagementRoutes.POST("", middleware.ValidateScopes(authmodel.WriteIncident), createIncident)
		incidentManagementRoutes.PUT("/:id", middleware.ValidateScopes(authmodel.ModifyIncident), updateIncident)

		// Incident Comment Routes
		incidentManagementRoutes.DELETE("/:id/comments/:commentId", middleware.ValidateScopes(authmodel.DeleteIncidentComment), deleteIncidentComment)
		incidentManagementRoutes.GET("/:id/comments", middleware.ValidateScopes(authmodel.ReadIncidentComment), listIncidentComments)
		incidentManagementRoutes.GET("/:id/comments/:commentId", middleware.ValidateScopes(authmodel.ReadIncidentComment), getIncidentComment)
		incidentManagementRoutes.POST("/:id/comments", middleware.ValidateScopes(authmodel.WriteIncidentComment), createIncidentComment)
		incidentManagementRoutes.PUT("/:id/comments/:commentId", middleware.ValidateScopes(authmodel.ModifyIncidentComment), updateIncidentComment)
	}
}

func createIncident(c *gin.Context) {
	var json incidentmodel.CreateIncidentRequest
	transactionId, _ := utils.GetTransactionIdHeader(c)
	tenantId := c.GetString("tenantId")
	if err := c.ShouldBindJSON(&json); err != nil {
		utilErr := utils.GenerateError(auditmodel.IncidentCreated, utils.InvalidRequestBody, http.StatusBadRequest, transactionId.TransactionId, err)
		utils.WriteErrorLog(utilErr)
		c.AbortWithStatusJSON(utilErr.ExternalError.Code, utilErr.ExternalError)
		return
	}
	result, err := database.Client.CreateIncident(tenantId, incidentmodel.Incident{
		AssignedTo:  json.AssignedTo,
		Description: json.Description,
		Severity:    json.Severity,
		Status:      json.Status,
		Tags:        json.Tags,
		TenantId:    tenantId,
		Title:       json.Title,
		Tlp:         json.Tlp,
		Version:     1,
	})
	if err != nil {
		utilErr := utils.GenerateError(auditmodel.IncidentCreated, utils.GenericInternalServerErrorMessage, http.StatusInternalServerError, transactionId.TransactionId, err)
		utils.WriteErrorLog(utilErr)
		c.AbortWithStatusJSON(utilErr.ExternalError.Code, utilErr.ExternalError)
		return
	}
	c.JSON(http.StatusOK, result)
}

func deleteIncident(c *gin.Context) {
	transactionId, _ := utils.GetTransactionIdHeader(c)
	tenantId := c.GetString("tenantId")
	id := c.Param("id")
	err := database.Client.DeleteIncident(id, tenantId)
	if err == utils.ErrIncidentNotFound {
		utilErr := utils.GenerateError(auditmodel.IncidentDeleted, utils.IncidentNotFound, http.StatusNotFound, transactionId.TransactionId, nil)
		utils.WriteErrorLog(utilErr)
		c.AbortWithStatusJSON(utilErr.ExternalError.Code, utilErr.ExternalError)
		return
	} else if err != nil {
		utilErr := utils.GenerateError(auditmodel.IncidentDeleted, utils.GenericInternalServerErrorMessage, http.StatusInternalServerError, transactionId.TransactionId, err)
		utils.WriteErrorLog(utilErr)
		c.AbortWithStatusJSON(utilErr.ExternalError.Code, utilErr.ExternalError)
		return
	}
	c.Status(http.StatusNoContent)
}

func getIncident(c *gin.Context) {
	transactionId, _ := utils.GetTransactionIdHeader(c)
	tenantId := c.GetString("tenantId")
	id := c.Param("id")
	result, err := database.Client.GetIncident(id, tenantId)
	if err != nil {
		utilErr := utils.GenerateError(auditmodel.IncidentRetrieved, utils.IncidentNotFound, http.StatusNotFound, transactionId.TransactionId, nil)
		utils.WriteErrorLog(utilErr)
		c.AbortWithStatusJSON(utilErr.ExternalError.Code, utilErr.ExternalError)
		return
	} else if err != nil {
		utilErr := utils.GenerateError(auditmodel.IncidentRetrieved, utils.GenericInternalServerErrorMessage, http.StatusInternalServerError, transactionId.TransactionId, err)
		utils.WriteErrorLog(utilErr)
		c.AbortWithStatusJSON(utilErr.ExternalError.Code, utilErr.ExternalError)
		return
	}
	c.JSON(http.StatusOK, result)
}

func listIncidents(c *gin.Context) {
	transactionId, _ := utils.GetTransactionIdHeader(c)
	tenantId := c.GetString("tenantId")
	limit := utils.QueryLimit(c)
	results, err := database.Client.ListIncidents(limit, tenantId)
	if err == utils.ErrIncidentsNotFound {
		utilErr := utils.GenerateError(auditmodel.IncidentRetrieved, utils.IncidentsNotFound, http.StatusNotFound, transactionId.TransactionId, nil)
		utils.WriteErrorLog(utilErr)
		c.AbortWithStatusJSON(utilErr.ExternalError.Code, utilErr.ExternalError)
		return
	} else if err != nil {
		utilErr := utils.GenerateError(auditmodel.IncidentRetrieved, utils.GenericInternalServerErrorMessage, http.StatusInternalServerError, transactionId.TransactionId, err)
		utils.WriteErrorLog(utilErr)
		c.AbortWithStatusJSON(utilErr.ExternalError.Code, utilErr.ExternalError)
		return
	}
	c.JSON(http.StatusOK, results)
}

func updateIncident(c *gin.Context) {
	var json incidentmodel.UpdateIncidentRequest
	transactionId, _ := utils.GetTransactionIdHeader(c)
	tenantId := c.GetString("tenantId")
	id := c.Param("id")
	if err := c.ShouldBindJSON(&json); err != nil {
		utilErr := utils.GenerateError(auditmodel.IncidentUpdated, utils.InvalidRequestBody, http.StatusBadRequest, transactionId.TransactionId, err)
		utils.WriteErrorLog(utilErr)
		c.AbortWithStatusJSON(utilErr.ExternalError.Code, utilErr.ExternalError)
		return
	}
	result, err := database.Client.UpdateIncident(id, tenantId, json)
	if err == utils.ErrIncidentNotFound {
		utilErr := utils.GenerateError(auditmodel.IncidentUpdated, utils.IncidentNotFound, http.StatusNotFound, transactionId.TransactionId, err)
		utils.WriteErrorLog(utilErr)
		c.AbortWithStatusJSON(utilErr.ExternalError.Code, utilErr.ExternalError)
		return
	} else if err != nil {
		utilErr := utils.GenerateError(auditmodel.IncidentUpdated, utils.GenericInternalServerErrorMessage, http.StatusInternalServerError, transactionId.TransactionId, err)
		utils.WriteErrorLog(utilErr)
		c.AbortWithStatusJSON(utilErr.ExternalError.Code, utilErr.ExternalError)
		return
	}
	c.JSON(http.StatusOK, result)
}

func createIncidentComment(c *gin.Context) {
	var json incidentcommentmodel.CreateIncidentComment
	transactionId, _ := utils.GetTransactionIdHeader(c)
	tenantId := c.GetString("tenantId")
	incidentId := c.Param("id")
	if err := c.ShouldBindJSON(&json); err != nil {
		utilErr := utils.GenerateError(auditmodel.IncidentCommentCreated, utils.InvalidRequestBody, http.StatusBadRequest, transactionId.TransactionId, err)
		utils.WriteErrorLog(utilErr)
		c.AbortWithStatusJSON(utilErr.ExternalError.Code, utilErr.ExternalError)
		return
	}
	result, err := database.Client.CreateIncidentComment(tenantId, incidentcommentmodel.IncidentComment{
		IncidentId: incidentId,
		Comment:    json.Comment,
		CreatedAt:  time.Now(),
		TenantId:   tenantId,
	})
	if err != nil {
		utilErr := utils.GenerateError(auditmodel.IncidentCommentCreated, utils.GenericInternalServerErrorMessage, http.StatusInternalServerError, transactionId.TransactionId, err)
		utils.WriteErrorLog(utilErr)
		c.AbortWithStatusJSON(utilErr.ExternalError.Code, utilErr.ExternalError)
	}
	c.JSON(http.StatusOK, result)
}

func deleteIncidentComment(c *gin.Context) {
	transactionId, _ := utils.GetTransactionIdHeader(c)
	tenantId := c.GetString("tenantId")
	id := c.Param("id")
	commentId := c.Param("commentId")
	err := database.Client.DeleteIncidentComment(commentId, id, tenantId)
	if err == utils.ErrIncidentCommentNotFound {
		utilErr := utils.GenerateError(auditmodel.IncidentCommentDeleted, utils.IncidentCommentNotFound, http.StatusNotFound, transactionId.TransactionId, nil)
		utils.WriteErrorLog(utilErr)
		c.AbortWithStatusJSON(utilErr.ExternalError.Code, utilErr.ExternalError)
		return
	} else if err != nil {
		utilErr := utils.GenerateError(auditmodel.IncidentCommentDeleted, utils.GenericInternalServerErrorMessage, http.StatusInternalServerError, transactionId.TransactionId, err)
		utils.WriteErrorLog(utilErr)
		c.AbortWithStatusJSON(utilErr.ExternalError.Code, utilErr.ExternalError)
		return
	}
	c.Status(http.StatusNoContent)
}

func getIncidentComment(c *gin.Context) {
	transactionId, _ := utils.GetTransactionIdHeader(c)
	tenantId := c.GetString("tenantId")
	id := c.Param("id")
	commentId := c.Param("commentId")
	result, err := database.Client.GetIncidentComment(commentId, id, tenantId)
	if err == utils.ErrIncidentCommentNotFound {
		utilErr := utils.GenerateError(auditmodel.IncidentCommentRetrieved, utils.IncidentCommentNotFound, http.StatusNotFound, transactionId.TransactionId, nil)
		utils.WriteErrorLog(utilErr)
		c.AbortWithStatusJSON(utilErr.ExternalError.Code, utilErr.ExternalError)
		return
	} else if err != nil {
		utilErr := utils.GenerateError(auditmodel.IncidentCommentRetrieved, utils.GenericInternalServerErrorMessage, http.StatusInternalServerError, transactionId.TransactionId, err)
		utils.WriteErrorLog(utilErr)
		c.AbortWithStatusJSON(utilErr.ExternalError.Code, utilErr.ExternalError)
		return
	}
	c.JSON(http.StatusOK, result)
}

func listIncidentComments(c *gin.Context) {
	transactionId, _ := utils.GetTransactionIdHeader(c)
	tenantId := c.GetString("tenantId")
	id := c.Param("id")
	limit := utils.QueryLimit(c)
	results, err := database.Client.ListIncidentComments(limit, id, tenantId)
	if err == utils.ErrIncidentCommentsNotFound {
		utilErr := utils.GenerateError(auditmodel.IncidentCommentRetrieved, utils.IncidentCommentsNotFound, http.StatusNotFound, transactionId.TransactionId, nil)
		utils.WriteErrorLog(utilErr)
		c.AbortWithStatusJSON(utilErr.ExternalError.Code, utilErr.ExternalError)
		return
	} else if err != nil {
		utilErr := utils.GenerateError(auditmodel.IncidentCommentRetrieved, utils.GenericInternalServerErrorMessage, http.StatusInternalServerError, transactionId.TransactionId, err)
		utils.WriteErrorLog(utilErr)
		c.AbortWithStatusJSON(utilErr.ExternalError.Code, utilErr.ExternalError)
		return
	}
	c.JSON(http.StatusOK, results)
}

func updateIncidentComment(c *gin.Context) {
	var json incidentcommentmodel.UpdateIncidentComment
	transactionId, _ := utils.GetTransactionIdHeader(c)
	tenantId := c.GetString("tenantId")
	id := c.Param("id")
	commentId := c.Param("commentId")
	if err := c.ShouldBindJSON(&json); err != nil {
		utilErr := utils.GenerateError(auditmodel.IncidentCommentUpdated, utils.InvalidRequestBody, http.StatusBadRequest, transactionId.TransactionId, err)
		utils.WriteErrorLog(utilErr)
		c.AbortWithStatusJSON(utilErr.ExternalError.Code, utilErr.ExternalError)
		return
	}
	result, err := database.Client.UpdateIncidentComment(commentId, id, tenantId, json)
	if err == utils.ErrIncidentCommentNotFound {
		utilErr := utils.GenerateError(auditmodel.IncidentCommentUpdated, utils.IncidentCommentNotFound, http.StatusNotFound, transactionId.TransactionId, err)
		utils.WriteErrorLog(utilErr)
		c.AbortWithStatusJSON(utilErr.ExternalError.Code, utilErr.ExternalError)
		return
	} else if err != nil {
		utilErr := utils.GenerateError(auditmodel.IncidentCommentUpdated, utils.GenericInternalServerErrorMessage, http.StatusInternalServerError, transactionId.TransactionId, err)
		utils.WriteErrorLog(utilErr)
		c.AbortWithStatusJSON(utilErr.ExternalError.Code, utilErr.ExternalError)
		return
	}
	c.JSON(http.StatusOK, result)
}
