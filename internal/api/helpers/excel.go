package helpers

import "github.com/xuri/excelize/v2"

type CellValue struct {
	Value      interface{}
	SheetName  string
	Cell       string
	CellMerge  string
	Style      int // need to be defined first with excelize.File.NewStyle()
	CellCol    string
	CellRow    string
	CellWidth  float64
	CellHeight float64
}

func SetCellValue(f *excelize.File, c CellValue) *excelize.File {
	// apply style
	if c.Style > 0 {
		f.SetCellStyle(c.SheetName, c.Cell, c.Cell, c.Style)
	}

	// merging cells
	if c.CellMerge != "" {
		f.MergeCell(c.SheetName, c.Cell, c.CellMerge)
	}

	// set value
	f.SetCellValue(c.SheetName, c.Cell, c.Value)

	return f
}
