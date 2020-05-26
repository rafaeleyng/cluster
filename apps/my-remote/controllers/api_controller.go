package controllers

import (
	"github.com/gin-gonic/gin"
)

func SetupAPIRoutes(router gin.IRouter) {
	apiGroup := router.Group("/api")
	v1Group := apiGroup.Group("/v1")
	apiRootGroup := apiGroup.Group("/")

	SetupV1Routes(v1Group)
	SetupAPIRootRoutes(apiRootGroup)
}
