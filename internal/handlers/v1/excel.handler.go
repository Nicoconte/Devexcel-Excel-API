package handlers

import (
	"devexcel-excel-api/internal/services"
	"devexcel-excel-api/internal/types"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func GenerateExcelHandler(ctx *gin.Context) {
	excelParam := &types.ExcelParams{}

	err := ctx.Bind(&excelParam)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	excelParam.Filename = strings.ReplaceAll(excelParam.Filename, " ", "_")

	target, err := services.GenerateExcel(*excelParam)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	filename := fmt.Sprintf("%s.xlsx", excelParam.Filename)

	fmt.Println("Salida ", target)

	ctx.Header("Content-Description", "File Transfer")
	ctx.Header("Content-Transfer-Encoding", "binary")
	ctx.Header("Content-Disposition", "attachment; filename="+filename)
	ctx.Header("Content-Type", "application/octet-stream")
	ctx.File(target)

	os.Remove(target)

	return
}
