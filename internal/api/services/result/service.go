package result

import (
	"Sesuai/internal/api/constracts"
	"Sesuai/internal/api/entities"
	"Sesuai/internal/api/helpers"
	"Sesuai/internal/api/models/response"
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

func (s Service) GetResults(userId string) (listResult []response.Result, err error) {
	results, err := s.repo.FindResults(userId)

	if len(results) > 0 {
		for _, result := range results {
			listResult = append(listResult, response.Result{
				SubmissionId:  result.SubmissionId,
				Timer:         result.Timer,
				TotalQuestion: result.TotalSubmission + " / " + result.TotalQuestion,
				CreatedAt:     result.CreatedAt,
				Token:         result.Token,
			})
		}
	} else {
		listResult = []response.Result{}
	}

	return
}

func (s Service) GetResultBySubmissionId(userId, submissionId string) (results []entities.Result, err error) {
	results, err = s.repo.FindResultBySubmissionId(userId, submissionId)

	if len(results) > 0 {
		results = helpers.FormattedPoint(results)
	}

	return
}

func (s Service) GetAllResultBySubmissionId(userId, submissionId string) (results []entities.Result, err error) {
	results, err = s.repo.FindAllResultBySubmissionId(userId, submissionId)

	if len(results) > 0 {
		results = helpers.FormattedPoint(results)
	}

	return
}
