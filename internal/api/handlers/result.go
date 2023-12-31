package handlers

import (
	"Sesuai/internal/api/constants"
	"Sesuai/internal/api/entities"
	"Sesuai/internal/api/helpers"
	"Sesuai/pkg/ahttp"
	"fmt"
	"github.com/kataras/iris/v12"
)

func GetResults(c iris.Context) {
	headers := helpers.GetHeaders(c)

	userId := c.Values().GetString(constants.AuthUserId)

	results, err := app.Services.Result.GetResults(userId)
	if err != nil {
		HttpError(c, headers, fmt.Errorf(err.Error()), ahttp.ErrFailure(err.Error()))
		return
	}

	data := make(map[string]interface{})
	data["result_list"] = results

	if len(results) > 0 {
		data["result_list"] = results
	}

	HttpSuccess(c, headers, data)
	return
}

func GetResultBySubmissionId(c iris.Context) {
	headers := helpers.GetHeaders(c)

	userId := c.Values().GetString(constants.AuthUserId)

	submissionId := c.Params().GetString("submissionId")
	if submissionId == "" {
		HttpError(c, headers, fmt.Errorf("Submission Id Empty"), ahttp.ErrFailure("submission_id_empty"))
		return
	}

	result, err := app.Services.Result.GetResultBySubmissionId(userId, submissionId)
	if err != nil {
		HttpError(c, headers, fmt.Errorf(err.Error()), ahttp.ErrFailure(err.Error()))
		return
	}

	adminPhoneNumber, err := app.Services.AdminPhoneNumber.GetAdminPhoneNumber()
	if err != nil {
		HttpError(c, headers, fmt.Errorf(err.Error()), ahttp.ErrFailure(err.Error()))
		return
	}

	for index := range result {
		result[index].Note = "Your " + result[index].ElementName + " element value is " + result[index].Point
	}

	data := make(map[string]interface{})
	data["result_list"] = []entities.Result{}
	data["admin_phone_number"] = ""

	if len(result) > 0 {
		data["result_list"] = result
	}

	if adminPhoneNumber.Id != "" {
		data["admin_phone_number"] = adminPhoneNumber.PhoneNumber
	}

	HttpSuccess(c, headers, data)
	return
}

func GetAllResult(c iris.Context) {
	headers := helpers.GetHeaders(c)

	userId := c.Values().GetString(constants.AuthUserId)

	params := c.FormValue("token")

	if params == "" {
		HttpError(c, headers, fmt.Errorf("Token Can't Empty"), ahttp.ErrFailure("token_can't_empty"))
		return
	}

	token, err := app.Services.GenerateToken.GetGenerateTokenByToken(params)
	if err != nil {
		HttpError(c, headers, fmt.Errorf("Token Not Found"), ahttp.ErrFailure("token_not_found"))
		return
	}

	if token.Status == "non active" {
		isUserToken, err := app.Services.UsedToken.IsUserToken(params, userId)
		if err != nil {
			HttpError(c, headers, fmt.Errorf(err.Error()), ahttp.ErrFailure(err.Error()))
			return
		}

		if !isUserToken {
			HttpError(c, headers, fmt.Errorf("Token Already Used"), ahttp.ErrFailure("token_already_used"))
			return
		}
	} else {
		err = app.Services.GenerateToken.ToggleInactiveToken(token.Id)
		if err != nil {
			HttpError(c, headers, fmt.Errorf(err.Error()), ahttp.ErrFailure(err.Error()))
			return
		}

		err = app.Services.UsedToken.InsertUsedToken(token.Id, userId)
		if err != nil {
			HttpError(c, headers, fmt.Errorf(err.Error()), ahttp.ErrFailure(err.Error()))
			return
		}
	}

	allResult, err := app.Services.Result.GetAllResult(userId)
	if err != nil {
		HttpError(c, headers, fmt.Errorf(err.Error()), ahttp.ErrFailure(err.Error()))
		return
	}

	for index := range allResult {
		allResult[index].Note = "Your " + allResult[index].ElementName + " element value is " + allResult[index].Point
	}

	data := make(map[string]interface{})
	data["result_list"] = []entities.Result{}

	if len(allResult) > 0 {
		data["result_list"] = allResult
	}

	HttpSuccess(c, headers, data)
	return
}
