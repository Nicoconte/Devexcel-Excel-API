package handlers

import (
	"devexcel-excel-api/internal/services"
	"devexcel-excel-api/internal/types"
	"devexcel-excel-api/internal/utils"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func GenerateExcelHandler(ctx *gin.Context) {
	excel := &types.Excel{}

	err := ctx.Bind(&excel)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Cannot bind excel struct. Error: %s", err.Error())})
		return
	}

	if len(excel.Spreadsheets) <= 0 || excel.Spreadsheets == nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": errors.New("Spreadsheets cannot be empty")})
		return
	}

	if excel.Filename == "" {
		excel.Filename = utils.NewGuid()
	} else {
		excel.Filename = strings.ReplaceAll(excel.Filename, " ", "_")
	}

	target, err := services.BuildExcel(*excel)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Cannot generate excel. Error: %s", err.Error())})
		return
	}

	filename := fmt.Sprintf("%s.xlsx", excel.Filename)

	ctx.Header("Content-Description", "File Transfer")
	ctx.Header("Content-Transfer-Encoding", "binary")
	ctx.Header("Content-Disposition", "attachment; filename="+filename)
	ctx.Header("Content-Type", "application/octet-stream")
	ctx.File(target)

	utils.DeleteFileFromStorage(target)
}
