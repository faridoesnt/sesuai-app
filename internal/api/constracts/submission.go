package constracts

import (
	"Sesuai/internal/api/entities"
	"Sesuai/internal/api/models/response"
)

type SubmissionRepository interface {
	FindSubmissions() (submissions []entities.Submission, err error)
	FindResultSubmission(submissionId string) (resultSubmission []entities.ResultSubmission, err error)
}

type SubmissionService interface {
	GetSubmissions() (submissions []response.Submission)
	GetResultSubmission(submissionId string) (resultSubmission []response.ResultSubmission, err error)
}
