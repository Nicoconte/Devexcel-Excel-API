package routes

import (
	"devexcel-excel-api/internal/handlers/v1"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RoutesHandler() *gin.Engine {
	r := gin.Default()

	r.GET("api/v1/test", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"test": "hola mundo"})
		return
	})

	r.POST("api/v1/excel", handlers.GenerateExcelHandler)

	return r
}
