package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RoutesHandler() *gin.Engine {
	r := gin.Default()

	r.GET("api/v1/test", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"test": "hola mundo"})
		return
	})

	return r
}
