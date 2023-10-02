package handlers

import (
	"Sesuai/internal/api/entities"
	"Sesuai/internal/api/helpers"
	"Sesuai/pkg/ahttp"
	"fmt"
	"github.com/kataras/iris/v12"
)

func GetBloodTypePoint(c iris.Context) {
	headers := helpers.GetHeaders(c)

	categoryId := c.Params().GetString("categoryId")

	bloodTypePoint, err := app.Services.BloodTypePoint.GetBloodTypePoint(categoryId)
	if err != nil {
		HttpError(c, headers, fmt.Errorf(err.Error()), ahttp.ErrFailure(err.Error()))
		return
	}

	data := make(map[string]interface{})
	data["blood_type_point_list"] = []entities.BloodTypePoint{}

	if len(bloodTypePoint) > 0 {
		data["blood_type_point_list"] = bloodTypePoint
	}

	HttpSuccess(c, headers, data)
	return
}

func UpdateBloodTypePoint(c iris.Context) {
	headers := helpers.GetHeaders(c)

	params := entities.RequestBloodTypePoint{}

	err := c.ReadJSON(&params)
	if err != nil {
		HttpError(c, headers, fmt.Errorf(err.Error()), ahttp.ErrFailure(err.Error()))
		return
	}

	if len(params.Point) == 0 {
		HttpError(c, headers, fmt.Errorf("point cant empty"), ahttp.ErrFailure("point_cant_empty"))
		return
	}

	if len(params.BloodTypeId) == 0 {
		HttpError(c, headers, fmt.Errorf("blood type id cant empty"), ahttp.ErrFailure("blood_type_id_cant_empty"))
		return
	}

	if params.CategoryId == "" {
		HttpError(c, headers, fmt.Errorf("category id cant empty"), ahttp.ErrFailure("category_id_cant_empty"))
		return
	}

	if len(params.Point) < len(params.BloodTypeId) || len(params.Point) > len(params.BloodTypeId) {
		HttpError(c, headers, fmt.Errorf("length point and blood type not same"), ahttp.ErrFailure("length_point_and_blood_type_not_same"))
		return
	}

	err = app.Services.BloodTypePoint.UpdateBloodTypePoint(params)
	if err != nil {
		HttpError(c, headers, fmt.Errorf("error update blood type point"), ahttp.ErrFailure(err.Error()))
		return
	}

	HttpSuccess(c, headers, nil)
	return
}
