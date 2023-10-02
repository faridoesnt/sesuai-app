package handlers

import (
	"Sesuai/internal/api/entities"
	"Sesuai/internal/api/helpers"
	"Sesuai/pkg/ahttp"
	"fmt"
	"github.com/kataras/iris/v12"
)

func GetPointAnswer(c iris.Context) {
	headers := helpers.GetHeaders(c)

	pointAnswer, err := app.Services.PointAnswer.GetPointAnswer()
	if err != nil {
		HttpError(c, headers, fmt.Errorf(err.Error()), ahttp.ErrFailure(err.Error()))
		return
	}

	data := make(map[string]interface{})
	data["point_answer_list"] = []entities.PointAnswer{}

	if len(pointAnswer) > 0 {
		data["point_answer_list"] = pointAnswer
	}

	HttpSuccess(c, headers, data)
	return
}
