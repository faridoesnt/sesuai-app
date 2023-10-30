package handlers

import (
	"Sesuai/internal/api/helpers"
	"Sesuai/internal/api/models/response"
	"Sesuai/pkg/ahttp"
	"fmt"
	"github.com/kataras/iris/v12"
)

func GetHoroscope(c iris.Context) {
	headers := helpers.GetHeaders(c)

	horoscopes, err := app.Services.Horoscope.GetHoroscopes()
	if err != nil {
		HttpError(c, headers, fmt.Errorf(err.Error()), ahttp.ErrFailure(err.Error()))
		return
	}

	data := make(map[string]interface{})
	data["horoscope_list"] = []response.Horoscope{}

	if len(horoscopes) > 0 {
		data["horoscope_list"] = horoscopes
	}

	HttpSuccess(c, headers, data)
	return
}
