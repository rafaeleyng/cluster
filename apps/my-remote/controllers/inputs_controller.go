package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rafaeleyng/my-remote/business"
)

func handleGetMappings(c *gin.Context) {
	c.JSON(http.StatusOK, business.Response{
		Data: business.InputMappings,
	})
}

func handlePut(c *gin.Context) {
	frequency := c.Param("frequency")

	input, ok := business.InputMappings[frequency]
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"data":  nil,
			"error": fmt.Sprintf("Invalid input mapping with frequency '%s'", frequency),
		})
		return
	}

	business.HandleInput(input)
	c.JSON(http.StatusOK, business.Response{
		Data: input,
	})
}

func SetupInputsRoutes(router gin.IRouter) {
	router.GET("/mappings", handleGetMappings)
	router.PUT("/:frequency", handlePut)
}
