package recapSubmissions

import (
	"Sesuai/internal/api/constants"
	"Sesuai/internal/api/constracts"
	"Sesuai/internal/api/entities"
	"Sesuai/internal/api/helpers"
	"github.com/xuri/excelize/v2"
	"log"
	"os"
	"sync"
)

type Service struct {
	app  constracts.App
	repo constracts.RecapSubmissionsRepository
}

func Init(a *constracts.App) (svc constracts.RecapSubmissionsService) {
	r := initRepository(a.Datasources.WriterDB, a.Datasources.ReaderDB)

	svc = &Service{
		app:  *a,
		repo: r,
	}

	return
}

func SetTemplateLoc(templateLoc string) string {
	currentDir, _ := os.Getwd()
	templateLoc = currentDir + "/internal/api" + templateLoc

	return templateLoc
}

func (s Service) GetRecapSubmissions(params entities.RequestRecapSubmissions) (resultRecapSubmissions []entities.ResultRecapSubmissions, err error) {
	recapUser, err := s.repo.FindRecapUser(params)

	if len(recapUser) > 0 {
		for _, value := range recapUser {
			var summaries []entities.SummariesSubmission

			totalSubmissions := 0
			totalUnlockSubmissions := 0

			recapSubmissions, _ := s.repo.CountRecapSubmissionsUser(value.UserId)

			if recapSubmissions.TotalSubmissions != 0 {
				totalSubmissions = recapSubmissions.TotalSubmissions

				summaries, err = s.GetSummariesSubmissionByUserId(value.UserId)
			}

			if recapSubmissions.TotalUnlockSubmissions != 0 {
				totalUnlockSubmissions = recapSubmissions.TotalUnlockSubmissions
			}

			resultRecapSubmissions = append(resultRecapSubmissions, entities.ResultRecapSubmissions{
				UserId:                 value.UserId,
				Name:                   value.Name,
				BirthDate:              value.BirthDate,
				Gender:                 value.Gender,
				Horoscope:              value.Horoscope,
				Shio:                   value.Shio,
				BloodType:              value.BloodType,
				TotalSubmissions:       totalSubmissions,
				TotalUnlockSubmissions: totalUnlockSubmissions,
				Summaries:              summaries,
			})
		}
	}

	return
}

func (s Service) GetSummariesSubmissionByUserId(userId string) (summaries []entities.SummariesSubmission, err error) {
	var summariesPoint []entities.SummariesPointSubmission

	summaries, err = s.repo.FindSummariesSubmissionByUserId(userId)

	if len(summaries) > 0 {
		for index, val := range summaries {

			summariesPoint, err = s.repo.FindSummariesPointSubmissionBySubmissionId(val.SubmissionId)

			if len(summariesPoint) > 0 {
				summariesPoint = helpers.FormattedPointSummaries(summariesPoint)
			}

			summaries[index].Points = summariesPoint
		}
	}

	return
}

func (s Service) GenerateExcel(data []entities.ResultRecapSubmissions) (f *excelize.File, err error) {
	f, err = excelize.OpenFile(SetTemplateLoc(constants.RecapSubmissions))
	if err != nil {
		log.Println(err.Error())
		return
	}

	defer func() {
		if err := f.Close(); err != nil {
			log.Println(err.Error())
		}
	}()

	styleBody, _ := f.NewStyle(constants.ExcelStyleBody(true))

	startRow := 3
	number := 1

	startRowSheet2 := 3
	numberSheet2 := 1

	for _, val := range data {
		f.SetActiveSheet(0)
		sheetName := f.GetSheetName(0)

		var wg1 sync.WaitGroup

		wg1.Add(1)

		go func() {
			defer wg1.Done()

			cell := ""

			// no
			cell, _ = excelize.CoordinatesToCellName(1, startRow)
			f = helpers.SetCellValue(f, helpers.CellValue{Value: number, SheetName: sheetName, Cell: cell, Style: styleBody})

			// name
			cell, _ = excelize.CoordinatesToCellName(2, startRow)
			f = helpers.SetCellValue(f, helpers.CellValue{Value: val.Name, SheetName: sheetName, Cell: cell, Style: styleBody})

			// birth date
			cell, _ = excelize.CoordinatesToCellName(3, startRow)
			f = helpers.SetCellValue(f, helpers.CellValue{Value: val.BirthDate, SheetName: sheetName, Cell: cell, Style: styleBody})

			// gender
			cell, _ = excelize.CoordinatesToCellName(4, startRow)
			f = helpers.SetCellValue(f, helpers.CellValue{Value: val.Gender, SheetName: sheetName, Cell: cell, Style: styleBody})

			// horoscope
			cell, _ = excelize.CoordinatesToCellName(5, startRow)
			f = helpers.SetCellValue(f, helpers.CellValue{Value: val.Shio, SheetName: sheetName, Cell: cell, Style: styleBody})

			// zodiac
			cell, _ = excelize.CoordinatesToCellName(6, startRow)
			f = helpers.SetCellValue(f, helpers.CellValue{Value: val.Horoscope, SheetName: sheetName, Cell: cell, Style: styleBody})

			// blood type
			cell, _ = excelize.CoordinatesToCellName(7, startRow)
			f = helpers.SetCellValue(f, helpers.CellValue{Value: val.BloodType, SheetName: sheetName, Cell: cell, Style: styleBody})

			// total submissions
			cell, _ = excelize.CoordinatesToCellName(8, startRow)
			f = helpers.SetCellValue(f, helpers.CellValue{Value: val.TotalSubmissions, SheetName: sheetName, Cell: cell, Style: styleBody})

			// total unlock submissions
			cell, _ = excelize.CoordinatesToCellName(9, startRow)
			f = helpers.SetCellValue(f, helpers.CellValue{Value: val.TotalUnlockSubmissions, SheetName: sheetName, Cell: cell, Style: styleBody})

			if len(val.Summaries) > 0 {
				// Sheet 2
				f.SetActiveSheet(1)
				sheetName = f.GetSheetName(1)

				for _, _v := range val.Summaries {
					// no
					cell, _ = excelize.CoordinatesToCellName(1, startRowSheet2)
					f = helpers.SetCellValue(f, helpers.CellValue{Value: numberSheet2, SheetName: sheetName, Cell: cell, Style: styleBody})

					// name
					cell, _ = excelize.CoordinatesToCellName(2, startRowSheet2)
					f = helpers.SetCellValue(f, helpers.CellValue{Value: val.Name, SheetName: sheetName, Cell: cell, Style: styleBody})

					// token
					cell, _ = excelize.CoordinatesToCellName(3, startRowSheet2)
					f = helpers.SetCellValue(f, helpers.CellValue{Value: _v.Token, SheetName: sheetName, Cell: cell, Style: styleBody})

					for index, __v := range _v.Points {
						cell, _ = excelize.CoordinatesToCellName(4+index, startRowSheet2)
						f = helpers.SetCellValue(f, helpers.CellValue{Value: __v.Point, SheetName: sheetName, Cell: cell, Style: styleBody})
					}

					startRowSheet2++
					numberSheet2++
				}

			}

			startRow++
			number++
		}()

		wg1.Wait()
	}

	return
}
