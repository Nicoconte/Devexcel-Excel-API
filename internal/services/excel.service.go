package services

import (
	"devexcel-excel-api/internal/types"
	"devexcel-excel-api/internal/utils"
	"errors"
	"fmt"

	"github.com/xuri/excelize/v2"
)

func GenerateExcel(excelParam types.ExcelParams) (string, error) {
	spreadsheetSeen := make(map[string]bool)

	f := excelize.NewFile()

	for _, spreadSheet := range excelParam.Spreadsheets {

		if spreadsheetSeen[spreadSheet.Name] {
			return "", errors.New(fmt.Sprintf("Invalid Spreadsheet name %s. It's already exists", spreadSheet.Name))
		}

		f.NewSheet(spreadSheet.Name)

		err := setCellStyle(f, spreadSheet.Name, spreadSheet.Cells)
		if err != nil {
			return "", err
		}

		for columnIndex, column := range spreadSheet.Columns {
			startIndex := 1

			//Set column
			err := setColumnWidth(f, spreadSheet.Name, utils.Alphabet[columnIndex], column.Width)
			if err != nil {
				return "", err
			}

			//Set comments if there is
			err = setCellComments(f, spreadSheet.Name, column.Comments, columnIndex)
			if err != nil {
				return "", err
			}

			//Set title if it not empty
			if column.Title != "" {
				f.SetCellValue(spreadSheet.Name, fmt.Sprintf("%s1", utils.Alphabet[columnIndex]), column.Title)
				startIndex++
			}

			//Set values
			err = setCellValues(f, spreadSheet.Name, column.Values, columnIndex, startIndex)
			if err != nil {
				return "", err
			}

			startIndex = 0
		}

		//Spreadsheet processed FLAG
		spreadsheetSeen[spreadSheet.Name] = true
	}

	f.DeleteSheet("Sheet1")

	outputpath := fmt.Sprintf("%s/%s.xlsx", utils.GetStoragePath(), excelParam.Filename)

	if err := f.SaveAs(outputpath); err != nil {
		fmt.Println("Err Excel: ", err.Error())
		return "", err
	}

	return outputpath, nil
}

func setCellStyle(file *excelize.File, spreadsheetname string, cells []types.ExcelCell) error {
	for _, c := range cells {
		style, err := file.NewStyle(c.Style)
		if err != nil {
			return err
		}

		file.SetCellStyle(spreadsheetname, "A1", "B1", style)
	}

	return nil
}

func setCellValues(file *excelize.File, spreadsheetName string, values []any, columnIndex, startIndex int) error {
	if len(values) == 0 {
		return nil
	}

	rows := startIndex

	for _, v := range values {
		err := file.SetCellValue(spreadsheetName, fmt.Sprintf("%s%d", utils.Alphabet[columnIndex], rows), v)

		if err != nil {
			return err
		}
		rows++
	}

	return nil
}

func setCellComments(file *excelize.File, spreadsheetName string, comments []types.ExcelCellComment, columnIndex int) error {

	if len(comments) == 0 {
		return nil
	}

	for _, c := range comments {

		cParsed := fmt.Sprintf(`{"author":"%s ", "text":" %s"}`, c.Content.Author, c.Content.Text)

		err := file.AddComment(spreadsheetName, fmt.Sprintf("%s%d", utils.Alphabet[columnIndex], c.Index), cParsed)

		if err != nil {
			return err
		}
	}

	return nil
}

func setColumnWidth(file *excelize.File, spreadsheetName string, column string, width float64) error {
	if width == 0.0 {
		width = 20.0
	}

	err := file.SetColWidth(spreadsheetName, column, column, width)

	if err != nil {
		return err
	}

	return nil

}

//TODO: Ver de que los estilos los gestione yo desde el backend, definir parametros
//func setCellStyles(file *excelize.File, spreadsheetName string, cellStyles []types.ExcelCellStyle, columnIndex int) {
// for _, s := range cellStyles {
// 	file.SetCellStyle(spreadsheetName, fmt.Sprintf("%s%d", utils.Alphabet[columnIndex], s.Index), style)
// }
//}
