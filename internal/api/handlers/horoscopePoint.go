package handlers

import (
	"Sesuai/internal/api/helpers"
	"Sesuai/internal/api/models/response"
	"Sesuai/pkg/ahttp"
	"fmt"
	"github.com/kataras/iris/v12"
)

func GetHoroscopePoint(c iris.Context) {
	headers := helpers.GetHeaders(c)

	categoryId := c.Params().GetString("categoryId")

	horoscopePoint, err := app.Services.HoroscopePoint.GetHoroscopePoint(categoryId)
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
