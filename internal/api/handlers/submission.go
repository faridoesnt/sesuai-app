package handlers

import (
	"Sesuai/internal/api/constants"
	"Sesuai/internal/api/helpers"
	"Sesuai/pkg/ahttp"
	"fmt"
	"github.com/kataras/iris/v12"
)

func GetSubmissions(c iris.Context) {
	headers := helpers.GetHeaders(c)

	adminId := c.Values().GetString(constants.AuthUserId)

	hasAccess, err := app.Services.AccessMenu.IsAdminHasAccessMenu(adminId, constants.Submition)
	if err != nil {
		HttpError(c, headers, fmt.Errorf(err.Error()), ahttp.ErrFailure(err.Error()))
		return
	}

	if !hasAccess {
		HttpError(c, headers, fmt.Errorf("admin doesn't have access"), ahttp.ErrFailure("admin_doesn't_have_access"))
		return
	}

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
