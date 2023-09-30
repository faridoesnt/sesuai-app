package handlers

import (
	"Sesuai/internal/api/helpers"
	"github.com/kataras/iris/v12"
)

func BloodType(c iris.Context) {
	headers := helpers.GetHeaders(c)

	bloodType := app.Services.BloodType.GetBloodType()

	data := make(map[string]interface{})
	data["blood_type_list"] = bloodType

	HttpSuccess(c, headers, data)
	return
}
