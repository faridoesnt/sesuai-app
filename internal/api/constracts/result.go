package constracts

import (
	"Sesuai/internal/api/entities"
	"Sesuai/internal/api/models/response"
)

type ResultRepository interface {
	FindResults(userId string) (results []entities.Submission, err error)
	FindResultBySubmissionId(userId, submissionId string) (results []entities.Result, err error)
	FindAllResult(userId string) (results []entities.Result, err error)
}

type ResultService interface {
	GetResults(userId string) (listResult []response.Result, err error)
	GetResultBySubmissionId(userId, submissionId string) (results []entities.Result, err error)
	GetAllResult(userId string) (results []entities.Result, err error)
}
