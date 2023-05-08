package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func GenerateRouter() *gin.Engine {
	router := gin.Default()
	router.SetTrustedProxies(nil)
	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods: []string{
			"DELETE",
			"GET",
			"POST",
			"PUT",
		},
		AllowHeaders: []string{
			"Origin",
			"Authorization",
		},
		AllowCredentials: true,
		ExposeHeaders: []string{
			"Content-Length",
		},
		MaxAge: 12 * time.Hour,
	}))
	return router
}
