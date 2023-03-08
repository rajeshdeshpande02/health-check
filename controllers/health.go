package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthController struct{}

func (h HealthController) V1Status(c *gin.Context) {
	c.String(http.StatusOK, "Version V1 Working!")

}

func (h HealthController) V2Status(c *gin.Context) {
	c.String(http.StatusOK, "Version V2 Working!")

}
