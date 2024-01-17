package handlers

import (
	"Sesuai/internal/api/helpers"
	"Sesuai/internal/api/models/response"
	"github.com/kataras/iris/v12"
)

func Horoscope(c iris.Context) {
	headers := helpers.GetHeaders(c)

	var horoscopes []response.Horoscope

	shios, _ := app.Services.Shio.GetShio()

	if len(shios) > 0 {
		for _, shio := range shios {
			horoscope := response.Horoscope{
				Id:   shio.Id,
				Name: shio.Name,
			}

			horoscopes = append(horoscopes, horoscope)
		}
	}

	data := make(map[string]interface{})
	data["horoscope_list"] = horoscopes

	HttpSuccess(c, headers, data)
	return
}
