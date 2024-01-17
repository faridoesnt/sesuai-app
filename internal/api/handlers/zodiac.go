package handlers

import (
	"Sesuai/internal/api/helpers"
	"Sesuai/internal/api/models/response"
	"github.com/kataras/iris/v12"
)

func Zodiac(c iris.Context) {
	headers := helpers.GetHeaders(c)

	var zodiacs []response.Zodiac

	horoscopes, _ := app.Services.Horoscope.GetHoroscopes()

	if len(horoscopes) > 0 {
		for _, horoscope := range horoscopes {
			zodiac := response.Zodiac{
				Id:   horoscope.Id,
				Name: horoscope.Name,
			}

			zodiacs = append(zodiacs, zodiac)
		}
	}

	data := make(map[string]interface{})
	data["zodiac_list"] = zodiacs

	HttpSuccess(c, headers, data)
	return
}
