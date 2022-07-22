package routes

import (
	"devexcel-excel-api/internal/handlers/v1"

	"github.com/gin-gonic/gin"
)

func RoutesHandler() *gin.Engine {
	r := gin.Default()

	r.POST("api/v1/excel", handlers.GenerateExcelHandler)

	return r
}
