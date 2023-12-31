package constracts

import "Sesuai/internal/api/entities"

type ResultRepository interface {
	FindResultBySubmissionId(userId, submissionId string) (results []entities.Result, err error)
	FindAllResult(userId string) (results []entities.Result, err error)
}

type ResultService interface {
	GetResultBySubmissionId(userId, submissionId string) (results []entities.Result, err error)
	GetAllResult(userId string) (results []entities.Result, err error)
}
