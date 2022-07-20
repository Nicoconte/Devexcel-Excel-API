package types

import "github.com/xuri/excelize/v2"

type Excel struct {
	Filename     string
	Spreadsheets []ExcelSpreadsheet
}

type ExcelSpreadsheet struct {
	Name    string
	Rows    []ExcelRow
	Columns []ExcelColumn
	Cells   []ExcelCell
}

type ExcelColumn struct {
	Title  string
	Values []any
	Width  float64
}

type ExcelCell struct {
	Axis    string
	Style   *excelize.Style
	Comment ExcelCellComment
}

type ExcelRow struct {
	Index  int
	Height float64
}

type ExcelCellComment struct {
	Content ExcelCellCommentContent
}

type ExcelCellCommentContent struct {
	Author string
	Text   string
}
