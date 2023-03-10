package server

import (
	"health-check/controllers"
	"health-check/middlerware"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(middlerware.TimeoutMiddleware())
	router.Use(middlerware.CORSMiddleware())

	health := new(controllers.HealthController)

	router.GET("v1/health", health.V1Status)

	router.GET("v2/health", health.V2Status)

	return router

}
