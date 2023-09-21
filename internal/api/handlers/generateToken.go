package handlers

import (
	"Sesuai/internal/api/constants"
	"Sesuai/internal/api/entities"
	"Sesuai/internal/api/helpers"
	"Sesuai/pkg/ahttp"
	"Sesuai/pkg/autils"
	"errors"
	"github.com/kataras/iris/v12"
)

func GetGenerateToken(c iris.Context) {
	headers := helpers.GetHeaders(c)

	adminId := c.Values().GetString(constants.AuthUserId)

	tokens := app.Services.GenerateToken.GetGenerateToken(adminId)

	HttpSuccess(c, headers, tokens)
	return
}

func GenerateNewToken(c iris.Context) {
	headers := helpers.GetHeaders(c)

	adminId := c.Values().GetString(constants.AuthUserId)

	newToken := autils.RandomString(5)

	err := app.Services.GenerateToken.InsertNewToken(adminId, newToken)
	if err != nil {
		HttpError(c, headers, errors.New("error generate new token"), ahttp.ErrFailure(err.Error()))
		return
	}

	HttpSuccess(c, headers, nil)
	return
}

func UseToken(c iris.Context) {
	headers := helpers.GetHeaders(c)

	params := entities.UseToken{}

	err := c.ReadJSON(&params)
	if err != nil {
		HttpError(c, headers, err, ahttp.ErrInvalid(err.Error()))
		return
	}

	err = app.Services.GenerateToken.UpdateToken(params.Id)
	if err != nil {
		HttpError(c, headers, errors.New("error use token"), ahttp.ErrFailure(err.Error()))
		return
	}

	HttpSuccess(c, headers, nil)
	return
}
