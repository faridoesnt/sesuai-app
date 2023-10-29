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
