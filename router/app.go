package router

import (
	"github.com/gin-gonic/gin"
	"kevinsheeran/hello_go/service"
)

func Router() *gin.Engine {

	router := gin.Default()

	router.GET("/index", service.GetIndex)
	return router
}
