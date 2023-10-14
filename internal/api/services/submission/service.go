package submission

import (
	"Sesuai/internal/api/constracts"
	"Sesuai/internal/api/models/response"
)

type Service struct {
	app  constracts.App
	repo constracts.SubmissionRepository
}

func Init(a *constracts.App) (svc constracts.SubmissionService) {
	r := initRepository(a.Datasources.WriterDB, a.Datasources.ReaderDB)

	svc = &Service{
		app:  *a,
		repo: r,
	}

	return
}

func (s Service) GetSubmissions() (submissions []response.Submission) {
	listSubmissions, _ := s.repo.FindSubmissions()

	if len(listSubmissions) > 0 {
		for _, submission := range listSubmissions {
			submissions = append(submissions, response.Submission{
				SubmissionId:  submission.SubmissionId,
				Name:          submission.Name,
				Email:         submission.Email,
				Timer:         submission.Timer,
				TotalQuestion: submission.TotalSubmission + " / " + submission.TotalQuestion,
				CreatedAt:     submission.CreatedAt,
			})
		}
	} else {
		submissions = []response.Submission{}
	}

	return
}

func (s Service) GetResultSubmission(submissionId string) (resultSubmission []response.ResultSubmission, err error) {
	results, err := s.repo.FindResultSubmission(submissionId)

	if len(resultSubmission) > 0 {
		for _, result := range results {
			resultSubmission = append(resultSubmission, response.ResultSubmission{
				ElementId:   result.ElementId,
				ElementName: result.ElementName,
				Point:       result.Point,
			})
		}
	} else {
		resultSubmission = []response.ResultSubmission{}
	}

	return
}
