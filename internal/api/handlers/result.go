package handlers

import (
	"Sesuai/internal/api/constants"
	"Sesuai/internal/api/entities"
	"Sesuai/internal/api/helpers"
	"Sesuai/pkg/ahttp"
	"fmt"
	"github.com/kataras/iris/v12"
)

func GetResult(c iris.Context) {
	headers := helpers.GetHeaders(c)

	userId := c.Values().GetString(constants.AuthUserId)

	result, err := app.Services.Result.GetResult(userId)
	if err != nil {
		HttpError(c, headers, fmt.Errorf(err.Error()), ahttp.ErrFailure(err.Error()))
		return
	}

	data := make(map[string]interface{})
	data["result_list"] = []entities.Result{}

	if len(result) > 0 {
		data["result_list"] = result
	}

	HttpSuccess(c, headers, data)
	return
}

func GetAllResult(c iris.Context) {
	headers := helpers.GetHeaders(c)

	userId := c.Values().GetString(constants.AuthUserId)

	params := entities.RequestAllResult{}

	err := c.ReadJSON(&params)
	if err != nil {
		HttpError(c, headers, fmt.Errorf(err.Error()), ahttp.ErrFailure(err.Error()))
		return
	}

	token, err := app.Services.GenerateToken.GetGenerateTokenByToken(params.Token)
	if err != nil {
		HttpError(c, headers, fmt.Errorf("Token Not Found"), ahttp.ErrFailure("token_not_found"))
		return
	}

	if token.Status == "non active" {
		isUserToken, err := app.Services.UsedToken.IsUserToken(params.Token, userId)
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

	data := make(map[string]interface{})
	data["result_list"] = []entities.Result{}

	if len(allResult) > 0 {
		data["result_list"] = allResult
	}

	HttpSuccess(c, headers, data)
	return
}
