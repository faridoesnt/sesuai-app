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

func GetHoroscopePoint(c iris.Context) {
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

	elementId := c.Params().GetString("elementId")

	horoscopePoint, err := app.Services.HoroscopePoint.GetHoroscopePoint(elementId)
	if err != nil {
		HttpError(c, headers, fmt.Errorf(err.Error()), ahttp.ErrFailure(err.Error()))
		return
	}

	data := make(map[string]interface{})
	data["horoscope_point_list"] = []response.HoroscopePoint{}

	if len(horoscopePoint) > 0 {
		data["horoscope_point_list"] = horoscopePoint
	}

	HttpSuccess(c, headers, data)
	return
}

func UpdateHoroscopePoint(c iris.Context) {
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

	params := entities.RequestHoroscopePoint{}

	err = c.ReadJSON(&params)
	if err != nil {
		HttpError(c, headers, fmt.Errorf(err.Error()), ahttp.ErrFailure(err.Error()))
		return
	}

	if len(params.Point) == 0 {
		HttpError(c, headers, fmt.Errorf("point cant empty"), ahttp.ErrFailure("point_cant_empty"))
		return
	}

	if len(params.HoroscopeId) == 0 {
		HttpError(c, headers, fmt.Errorf("horoscope id cant empty"), ahttp.ErrFailure("horoscope_id_cant_empty"))
		return
	}

	if params.ElementId == "" {
		HttpError(c, headers, fmt.Errorf("element id cant empty"), ahttp.ErrFailure("element_id_cant_empty"))
		return
	}

	if len(params.Point) < len(params.HoroscopeId) || len(params.Point) > len(params.HoroscopeId) {
		HttpError(c, headers, fmt.Errorf("length point and horoscope not same"), ahttp.ErrFailure("length_point_and_horoscope_not_same"))
		return
	}

	err = app.Services.HoroscopePoint.UpdateHoroscopePoint(params)
	if err != nil {
		HttpError(c, headers, fmt.Errorf("error update horoscope point"), ahttp.ErrFailure(err.Error()))
		return
	}

	HttpSuccess(c, headers, nil)
	return
}
