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
