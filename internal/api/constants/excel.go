package constants

import "github.com/xuri/excelize/v2"

const (
	ExcelGeneralBorderColor = "80807f"
)

const (
	RecapSubmissions = "/files/recap_submissions.xlsx"
)

func ExcelStyleBody(isCenter bool) *excelize.Style {
	alignment := &excelize.Alignment{}
	if isCenter {
		alignment.Horizontal = "center"
	}
	alignment.WrapText = true

	border := []excelize.Border{
		{Type: "top", Color: ExcelGeneralBorderColor, Style: 1},
		{Type: "left", Color: ExcelGeneralBorderColor, Style: 1},
		{Type: "bottom", Color: ExcelGeneralBorderColor, Style: 1},
		{Type: "right", Color: ExcelGeneralBorderColor, Style: 1},
	}

	return &excelize.Style{
		Alignment: alignment,
		Border:    border,
	}
}
