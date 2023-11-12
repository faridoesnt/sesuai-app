package handlers

import (
	"Sesuai/internal/api/constants"
	"Sesuai/internal/api/entities"
	"Sesuai/internal/api/helpers"
	"Sesuai/internal/api/models/response"
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

	data := make(map[string]interface{})
	data["submission_list"] = []response.Submission{}

	if len(submissions) > 0 {
		data["submission_list"] = submissions
	}

	HttpSuccess(c, headers, data)
	return
}

func GetResultSubmission(c iris.Context) {
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

	submissionId := c.Params().GetString("submissionId")

	if submissionId == "" {
		HttpError(c, headers, fmt.Errorf("submission id is empty"), ahttp.ErrFailure("submission_id_is_empty"))
		return
	}

	resultSubmission, err := app.Services.Submission.GetResultSubmission(submissionId)
	if err != nil {
		HttpError(c, headers, fmt.Errorf("error while get result submissions"), ahttp.ErrFailure(err.Error()))
		return
	}

	data := make(map[string]interface{})
	data["result_list"] = []entities.Result{}

	if len(resultSubmission) > 0 {
		data["result_list"] = resultSubmission
	}

	HttpSuccess(c, headers, data)
	return
}
