package constracts

import (
	"Sesuai/internal/api/entities"
	"github.com/xuri/excelize/v2"
)

type RecapSubmissionsRepository interface {
	FindRecapUser(params entities.RequestRecapSubmissions) (recapUser []entities.RecapUser, err error)
	CountRecapSubmissionsUser(userId string) (recapSubmissions entities.RecapSubmissions, err error)
}

type RecapSubmissionsService interface {
	GetRecapSubmissions(params entities.RequestRecapSubmissions) (resultRecapSubmissions []entities.ResultRecapSubmissions, err error)
	GenerateExcel(data []entities.ResultRecapSubmissions) (f *excelize.File, err error)
}
