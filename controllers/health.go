package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type HealthController struct{}

func (h HealthController) V1Status(c *gin.Context) {
	fmt.Println("Request Processing")
	fmt.Println(c.Request.Header.Get("Origin"))
	time.Sleep(1 * time.Second)
	c.String(http.StatusOK, "Version V1 Working!")

}

func (h HealthController) V2Status(c *gin.Context) {
	c.String(http.StatusOK, "Version V2 Working!")

}
