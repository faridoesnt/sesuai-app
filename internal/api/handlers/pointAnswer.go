package handlers

import (
	"Sesuai/internal/api/constants"
	"Sesuai/internal/api/entities"
	"Sesuai/internal/api/helpers"
	"Sesuai/pkg/ahttp"
	"fmt"
	"github.com/kataras/iris/v12"
)

func GetPointAnswer(c iris.Context) {
	headers := helpers.GetHeaders(c)

	adminId := c.Values().GetString(constants.AuthUserId)

	hasAccess, err := app.Services.AccessMenu.IsAdminHasAccessMenu(adminId, constants.Point)
	if err != nil {
		HttpError(c, headers, fmt.Errorf(err.Error()), ahttp.ErrFailure(err.Error()))
		return
	}

	if !hasAccess {
		HttpError(c, headers, fmt.Errorf("admin doesn't have access"), ahttp.ErrFailure("admin_doesn't_have_access"))
		return
	}

	pointAnswer, err := app.Services.PointAnswer.GetPointAnswer()
	if err != nil {
		HttpError(c, headers, fmt.Errorf(err.Error()), ahttp.ErrFailure(err.Error()))
		return
	}

	data := make(map[string]interface{})
	data["point_answer_list"] = []entities.PointAnswer{}

	if len(pointAnswer) > 0 {
		data["point_answer_list"] = pointAnswer
	}

	HttpSuccess(c, headers, data)
	return
}

func UpdatePointAnswer(c iris.Context) {
	headers := helpers.GetHeaders(c)

	adminId := c.Values().GetString(constants.AuthUserId)

	hasAccess, err := app.Services.AccessMenu.IsAdminHasAccessMenu(adminId, constants.Point)
	if err != nil {
		HttpError(c, headers, fmt.Errorf(err.Error()), ahttp.ErrFailure(err.Error()))
		return
	}

	if !hasAccess {
		HttpError(c, headers, fmt.Errorf("admin doesn't have access"), ahttp.ErrFailure("admin_doesn't_have_access"))
		return
	}

	params := entities.RequestPointAnswer{}

	err = c.ReadJSON(&params)
	if err != nil {
		HttpError(c, headers, fmt.Errorf(err.Error()), ahttp.ErrFailure(err.Error()))
		return
	}

	if len(params.Point) == 0 {
		HttpError(c, headers, fmt.Errorf("point cant empty"), ahttp.ErrFailure("point_cant_empty"))
		return
	}

	if len(params.PointAnswerId) == 0 {
		HttpError(c, headers, fmt.Errorf("point answer id cant empty"), ahttp.ErrFailure("point_answer_id_cant_empty"))
		return
	}

	if len(params.Point) < len(params.PointAnswerId) || len(params.Point) > len(params.PointAnswerId) {
		HttpError(c, headers, fmt.Errorf("length point and point answer not same"), ahttp.ErrFailure("length_point_and_point_answer_not_same"))
		return
	}

	err = app.Services.PointAnswer.UpdatePointAnswer(params)
	if err != nil {
		HttpError(c, headers, fmt.Errorf("error update point answer"), ahttp.ErrFailure(err.Error()))
		return
	}

	HttpSuccess(c, headers, nil)
	return
}
