package handlers

import (
	"Sesuai/internal/api/helpers"
	"Sesuai/pkg/ahttp"
	"github.com/kataras/iris/v12"
)

func GetSubmissions(c iris.Context) {
	headers := helpers.GetHeaders(c)

	submissions := app.Services.Submission.GetSubmissions()

	HttpSuccess(c, headers, submissions)
	return
}

func GetResultSubmission(c iris.Context) {
	headers := helpers.GetHeaders(c)

	submissionId := c.Params().GetString("submissionId")

	if submissionId == "" {
		HttpError(c, headers, ahttp.Error{Message: "Submission Id Is Empty"}, ahttp.ErrFailure("submission_id_is_empty"))
		return
	}

	resultSubmission, err := app.Services.Submission.GetResultSubmission(submissionId)
	if err != nil {
		HttpError(c, headers, ahttp.Error{Message: "Error while get result submissions"}, ahttp.ErrFailure(err.Error()))
		return
	}

	HttpSuccess(c, headers, resultSubmission)
	return
}
