package handlers

import (
	"Sesuai/internal/api/helpers"
	"github.com/kataras/iris/v12"
)

func BloodType(c iris.Context) {
	headers := helpers.GetHeaders(c)

	bloodType := app.Services.BloodType.GetBloodType()

	HttpSuccess(c, headers, bloodType)
	return
}
