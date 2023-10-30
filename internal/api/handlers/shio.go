package handlers

import (
	"Sesuai/internal/api/helpers"
	"Sesuai/internal/api/models/response"
	"Sesuai/pkg/ahttp"
	"fmt"
	"github.com/kataras/iris/v12"
)

func GetShio(c iris.Context) {
	headers := helpers.GetHeaders(c)

	shios, err := app.Services.Shio.GetShio()
	if err != nil {
		HttpError(c, headers, fmt.Errorf(err.Error()), ahttp.ErrFailure(err.Error()))
		return
	}

	data := make(map[string]interface{})
	data["shio_list"] = []response.Shio{}

	if len(shios) > 0 {
		data["shio_list"] = shios
	}

	HttpSuccess(c, headers, data)
	return
}
