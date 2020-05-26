package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rafaeleyng/my-remote/business"
)

type serviceStatus struct {
	Service string `json:"service"`
	Status  string `json:"status"`
}

func handleGet(c *gin.Context) {
	c.JSON(http.StatusOK, business.Response{
		Data: gin.H{
			"app": "my-remote",
		},
	})
}

func handleGetHealthcheck(c *gin.Context) {
	c.JSON(http.StatusOK, business.Response{
		Data: gin.H{
			"services": []serviceStatus{
				serviceStatus{
					Service: "app",
					Status:  "working",
				},
			},
		},
	})
}

func SetupAPIRootRoutes(router gin.IRouter) {
	router.GET("/", handleGet)
	router.GET("/healthcheck", handleGetHealthcheck)
}
