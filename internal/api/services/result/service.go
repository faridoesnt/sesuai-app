package result

import (
	"Sesuai/internal/api/constracts"
	"Sesuai/internal/api/entities"
	"Sesuai/internal/api/helpers"
)

type Service struct {
	app  constracts.App
	repo constracts.ResultRepository
}

func Init(a *constracts.App) (svc constracts.ResultService) {
	r := initRepository(a.Datasources.WriterDB, a.Datasources.ReaderDB)

	svc = &Service{
		app:  *a,
		repo: r,
	}

	return
}

func (s Service) GetResult(userId string) (results []entities.Result, err error) {
	results, err = s.repo.FindResult(userId)

	if len(results) > 0 {
		results = helpers.FormattedPoint(results)
	}

	return
}

func (s Service) GetAllResult(userId string) (results []entities.Result, err error) {
	results, err = s.repo.FindAllResult(userId)

	if len(results) > 0 {
		results = helpers.FormattedPoint(results)
	}

	return
}
