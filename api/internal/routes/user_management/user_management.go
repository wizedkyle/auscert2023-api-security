package user_management

import (
	"github.com/ForgeResponse/ForgeResponse/v2/internal/database"
	"github.com/ForgeResponse/ForgeResponse/v2/internal/middleware"
	"github.com/ForgeResponse/ForgeResponse/v2/internal/utils"
	"github.com/ForgeResponse/ForgeResponse/v2/pkg/model/auditmodel"
	"github.com/ForgeResponse/ForgeResponse/v2/pkg/model/authmodel"
	"github.com/ForgeResponse/ForgeResponse/v2/pkg/model/usermodel"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"time"
)

type flagResponse struct {
	Flag string `json:"flag"`
}

func GenerateUserManagementRoutes(router *gin.Engine) {
	userManagementRoutes := router.Group("/v1/users")
	userManagementRoutes.Use(middleware.ValidateAuth(), middleware.ValidateTransactionId())
	{
		userManagementRoutes.DELETE("/:id", middleware.ValidateScopes(authmodel.DeleteUsers), deleteUser)
		userManagementRoutes.DELETE("", deleteAllUsers)
		userManagementRoutes.GET("", middleware.ValidateScopes(authmodel.ReadUsers), listUsers)
		userManagementRoutes.GET("/:id", middleware.ValidateScopes(authmodel.ReadUsers), getUser)
		userManagementRoutes.POST("", middleware.ValidateScopes(authmodel.WriteUsers), createUser)
		userManagementRoutes.PUT("/:id", middleware.ValidateScopes(authmodel.ModifyUsers), updateUser)
	}
}

func createUser(c *gin.Context) {
	var json usermodel.CreateUserRequest
	transactionId, _ := utils.GetTransactionIdHeader(c)
	tenantId := c.GetString("tenantId")
	if err := c.ShouldBindJSON(&json); err != nil {
		utilErr := utils.GenerateError(auditmodel.UserCreated, utils.InvalidRequestBody, http.StatusBadRequest, transactionId.TransactionId, err)
		utils.WriteErrorLog(utilErr)
		c.AbortWithStatusJSON(utilErr.ExternalError.Code, utilErr.ExternalError)
		return
	}
	result, err := database.Client.CreateUser(tenantId, usermodel.User{
		TenantId:    tenantId,
		Email:       json.Email,
		FirstName:   json.FirstName,
		LastName:    json.LastName,
		CreatedTime: time.Now(),
		IsActive:    true,
	})
	if err == utils.ErrEmailAlreadyExists {
		utilErr := utils.GenerateError(auditmodel.UserCreated, utils.EmailAlreadyExists, http.StatusUnprocessableEntity, transactionId.TransactionId, nil)
		utils.WriteErrorLog(utilErr)
		c.AbortWithStatusJSON(utilErr.ExternalError.Code, utilErr.ExternalError)
		return
	} else if err != nil {
		utilErr := utils.GenerateError(auditmodel.UserCreated, utils.GenericInternalServerErrorMessage, http.StatusInternalServerError, transactionId.TransactionId, err)
		utils.WriteErrorLog(utilErr)
		c.AbortWithStatusJSON(utilErr.ExternalError.Code, utilErr.ExternalError)
		return
	}
	c.JSON(http.StatusOK, result)
}

func deleteUser(c *gin.Context) {
	transactionId, _ := utils.GetTransactionIdHeader(c)
	tenantId := c.GetString("tenantId")
	id := c.Param("id")
	err := database.Client.DeleteUser(id, tenantId)
	if err == utils.ErrUserNotFound {
		utilErr := utils.GenerateError(auditmodel.UserDeleted, utils.UserNotFound, http.StatusNotFound, transactionId.TransactionId, nil)
		utils.WriteErrorLog(utilErr)
		c.AbortWithStatusJSON(utilErr.ExternalError.Code, utilErr.ExternalError)
		return
	} else if err != nil {
		utilErr := utils.GenerateError(auditmodel.UserDeleted, utils.GenericInternalServerErrorMessage, http.StatusInternalServerError, transactionId.TransactionId, err)
		utils.WriteErrorLog(utilErr)
		c.AbortWithStatusJSON(utilErr.ExternalError.Code, utilErr.ExternalError)
		return
	}
	c.Status(http.StatusNoContent)
}

func deleteAllUsers(c *gin.Context) {
	var response flagResponse
	conferenceName, exists := c.GetQuery("conference_name")
	if conferenceName == "" && !exists {
		c.AbortWithStatusJSON(http.StatusForbidden, "no conference_name query string provided")
		return
	}
	if conferenceName == "AusCERT2023" {
		response.Flag = os.Getenv("FLAG")
		c.JSON(http.StatusOK, response)
		return
	} else {
		c.AbortWithStatusJSON(http.StatusBadRequest, "sorry looks like the conference name is wrong")
		return
	}

}

func getUser(c *gin.Context) {
	transactionId, _ := utils.GetTransactionIdHeader(c)
	tenantId := c.GetString("tenantId")
	id := c.Param("id")
	result, err := database.Client.GetUser(id, tenantId)
	if err == utils.ErrUserNotFound {
		utilErr := utils.GenerateError(auditmodel.UserRetrieved, utils.UserNotFound, http.StatusNotFound, transactionId.TransactionId, nil)
		utils.WriteErrorLog(utilErr)
		c.AbortWithStatusJSON(utilErr.ExternalError.Code, utilErr.ExternalError)
		return
	} else if err != nil {
		utilErr := utils.GenerateError(auditmodel.UserRetrieved, utils.GenericInternalServerErrorMessage, http.StatusInternalServerError, transactionId.TransactionId, err)
		utils.WriteErrorLog(utilErr)
		c.AbortWithStatusJSON(utilErr.ExternalError.Code, utilErr.ExternalError)
		return
	}
	c.JSON(http.StatusOK, result)
}

func listUsers(c *gin.Context) {
	transactionId, _ := utils.GetTransactionIdHeader(c)
	tenantId := c.GetString("tenantId")
	limit := utils.QueryLimit(c)
	results, err := database.Client.ListUsers(limit, tenantId)
	if err == utils.ErrUsersNotFound {
		utilErr := utils.GenerateError(auditmodel.UserRetrieved, utils.UsersNotFound, http.StatusNotFound, transactionId.TransactionId, nil)
		utils.WriteErrorLog(utilErr)
		c.AbortWithStatusJSON(utilErr.ExternalError.Code, utilErr.ExternalError)
		return
	} else if err != nil {
		utilErr := utils.GenerateError(auditmodel.UserRetrieved, utils.GenericInternalServerErrorMessage, http.StatusInternalServerError, transactionId.TransactionId, err)
		utils.WriteErrorLog(utilErr)
		c.AbortWithStatusJSON(utilErr.ExternalError.Code, utilErr.ExternalError)
		return
	}
	c.JSON(http.StatusOK, results)
}

func updateUser(c *gin.Context) {
	var json usermodel.UpdateUserRequest
	transactionId, _ := utils.GetTransactionIdHeader(c)
	tenantId := c.GetString("tenantId")
	id := c.Param("id")
	if err := c.ShouldBindJSON(&json); err != nil {
		utilErr := utils.GenerateError(auditmodel.UserUpdated, utils.InvalidRequestBody, http.StatusBadRequest, transactionId.TransactionId, err)
		utils.WriteErrorLog(utilErr)
		c.AbortWithStatusJSON(utilErr.ExternalError.Code, utilErr.ExternalError)
		return
	}
	result, err := database.Client.UpdateUser(id, tenantId, json)
	if err == utils.ErrUserNotFound {
		utilErr := utils.GenerateError(auditmodel.UserUpdated, utils.UserNotFound, http.StatusNotFound, transactionId.TransactionId, err)
		utils.WriteErrorLog(utilErr)
		c.AbortWithStatusJSON(utilErr.ExternalError.Code, utilErr.ExternalError)
		return
	} else if err != nil {
		utilErr := utils.GenerateError(auditmodel.UserUpdated, utils.GenericInternalServerErrorMessage, http.StatusInternalServerError, transactionId.TransactionId, err)
		utils.WriteErrorLog(utilErr)
		c.AbortWithStatusJSON(utilErr.ExternalError.Code, utilErr.ExternalError)
		return
	}
	c.JSON(http.StatusOK, result)
}
