package services

import (
	"devexcel-excel-api/internal/types"
	"devexcel-excel-api/internal/utils"
	"errors"
	"fmt"

	"github.com/xuri/excelize/v2"
)

var file *excelize.File = nil
var currSpreadsheetName string = ""

func BuildExcel(excel types.Excel) (string, error) {
	spreadsheetSeen := make(map[string]bool)

	file = excelize.NewFile()

	for _, spreadSheet := range excel.Spreadsheets {

		if spreadsheetSeen[spreadSheet.Name] {
			return "", errors.New(fmt.Sprintf("Invalid Spreadsheet name %s. It's already exists", spreadSheet.Name))
		}

		currSpreadsheetName = spreadSheet.Name

		file.NewSheet(currSpreadsheetName)

		buildColumns(spreadSheet.Columns)

		buildCells(spreadSheet.Cells)

		//Spreadsheet processed => FLAG
		spreadsheetSeen[spreadSheet.Name] = true
	}

	file.DeleteSheet("Sheet1")

	outputpath := fmt.Sprintf("%s/%s.xlsx", utils.GetStoragePath(), excel.Filename)

	if err := file.SaveAs(outputpath); err != nil {
		fmt.Println("Err Excel: ", err.Error())
		return "", err
	}

	return outputpath, nil
}

func buildCells(cells []types.ExcelCell) error {
	for _, cell := range cells {

		style, err := file.NewStyle(cell.Style)
		if err != nil {
			return err
		}
		file.SetCellStyle(currSpreadsheetName, cell.Axis, cell.Axis, style)

		if cell.Comment != (types.ExcelCellComment{}) {
			cParsed := fmt.Sprintf(`{"author":"%s ", "text":" %s"}`, cell.Comment.Content.Author, cell.Comment.Content.Text)
			file.AddComment(currSpreadsheetName, cell.Axis, cParsed)
		}
	}

	return nil
}

func buildColumns(columns []types.ExcelColumn) error {
	for index, column := range columns {

		axis := utils.Alphabet[index]

		startIndex := setColumnTitle(axis, column.Title)

		err := setColumnWidth(axis, column.Width)
		if err != nil {
			return err
		}

		err = setColumnValues(column.Values, axis, startIndex)
		if err != nil {
			return err
		}

		startIndex = 0
	}

	return nil
}

func setColumnTitle(axis string, title string) int {
	if title != "" {
		file.SetCellValue(currSpreadsheetName, fmt.Sprintf("%s1", axis), title)
		return 2
	}

	return 1
}

func setColumnValues(values []any, axis string, startIndex int) error {
	if len(values) == 0 {
		return nil
	}

	rows := startIndex

	for _, v := range values {
		err := file.SetCellValue(currSpreadsheetName, fmt.Sprintf("%s%d", axis, rows), v)

		if err != nil {
			return err
		}
		rows++
	}

	return nil
}

func setColumnWidth(axis string, width float64) error {
	if width == 0.0 {
		width = 20.0
	}

	err := file.SetColWidth(currSpreadsheetName, axis, axis, width)

	if err != nil {
		return err
	}

	return nil

}
