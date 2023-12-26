package constracts

import (
	"Sesuai/internal/api/entities"
	"Sesuai/internal/api/models/response"
)

type SubmissionRepository interface {
	FindSubmissions() (submissions []entities.Submission, err error)
	FindSubmissionsByEmailUser(emailUser string) (submissions []entities.Submission, err error)
	FindSubmissionsByFullName(fullName string) (submissions []entities.Submission, err error)
	FindResultSubmission(submissionId string) (resultSubmission []entities.Result, err error)
}

type SubmissionService interface {
	GetSubmissions() (submissions []response.Submission)
	GetSubmissionsByEmailUser(emailUser string) (submissions []response.Submission)
	GetSubmissionsByFullName(fullName string) (submissions []response.Submission)
	GetResultSubmission(submissionId string) (resultSubmission []entities.Result, err error)
}
