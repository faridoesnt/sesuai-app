package handlers

import (
	"Sesuai/internal/api/entities"
	"Sesuai/internal/api/helpers"
	"Sesuai/pkg/ahttp"
	"fmt"
	"github.com/kataras/iris/v12"
)

func GetShioPoint(c iris.Context) {
	headers := helpers.GetHeaders(c)

	categoryId := c.Params().GetString("categoryId")

	shioPoint, err := app.Services.ShioPoint.GetShioPoint(categoryId)
	if err != nil {
		HttpError(c, headers, fmt.Errorf(err.Error()), ahttp.ErrFailure(err.Error()))
		return
	}

	data := make(map[string]interface{})
	data["shio_point_list"] = []entities.ShioPoint{}

	if len(shioPoint) > 0 {
		data["shio_point_list"] = shioPoint
	}

	HttpSuccess(c, headers, data)
	return
}
