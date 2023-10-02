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

func UpdateShioPoint(c iris.Context) {
	headers := helpers.GetHeaders(c)

	params := entities.RequestShioPoint{}

	err := c.ReadJSON(&params)
	if err != nil {
		HttpError(c, headers, fmt.Errorf(err.Error()), ahttp.ErrFailure(err.Error()))
		return
	}

	if len(params.Point) == 0 {
		HttpError(c, headers, fmt.Errorf("point cant empty"), ahttp.ErrFailure("point_cant_empty"))
		return
	}

	if len(params.ShioId) == 0 {
		HttpError(c, headers, fmt.Errorf("shio id cant empty"), ahttp.ErrFailure("shio_id_cant_empty"))
		return
	}

	if params.CategoryId == "" {
		HttpError(c, headers, fmt.Errorf("category id cant empty"), ahttp.ErrFailure("category_id_cant_empty"))
		return
	}

	if len(params.Point) < len(params.ShioId) || len(params.Point) > len(params.ShioId) {
		HttpError(c, headers, fmt.Errorf("length point and shio not same"), ahttp.ErrFailure("length_point_and_shio_not_same"))
		return
	}

	err = app.Services.ShioPoint.UpdateShioPoint(params)
	if err != nil {
		HttpError(c, headers, fmt.Errorf("error update shio point"), ahttp.ErrFailure(err.Error()))
		return
	}

	HttpSuccess(c, headers, nil)
	return
}
