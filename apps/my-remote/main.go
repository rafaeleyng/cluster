package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rafaeleyng/my-remote/controllers"
)

func setupRouter() *gin.Engine {
	router := gin.Default()
	controllers.SetupAPIRoutes(router)
	return router
}

func main() {
	router := setupRouter()
	router.Run(":9000")
}
