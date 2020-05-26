package controllers

import (
	"github.com/gin-gonic/gin"
)

func SetupV1Routes(router gin.IRouter) {
	inputsGroup := router.Group("/inputs")

	SetupInputsRoutes(inputsGroup)
}
