package types

import "github.com/xuri/excelize/v2"

type ExcelParams struct {
	Filename     string
	Spreadsheets []ExcelSpreadsheet
}

type ExcelSpreadsheet struct {
	Name    string
	Cells   []ExcelCell
	Columns []ExcelColumn
}

type ExcelCell struct {
	Axis  string
	Style *excelize.Style
}

type ExcelColumn struct {
	Title    string
	Comments []ExcelCellComment
	Values   []any
	Width    float64
}

type ExcelCellComment struct {
	Index   int
	Content ExcelCellCommentContent
}

type ExcelCellCommentContent struct {
	Author string
	Text   string
}
