package handlers

import (
	"Sesuai/internal/api/entities"
	"Sesuai/internal/api/helpers"
	"Sesuai/internal/api/models/response"
	"Sesuai/pkg/ahttp"
	"fmt"
	"github.com/kataras/iris/v12"
)

func GetQuestionsTest(c iris.Context) {
	headers := helpers.GetHeaders(c)

	questionsTest, err := app.Services.QuestionTest.GetQuestionsTest()
	if err != nil {
		HttpError(c, headers, fmt.Errorf("Questions Not Found"), ahttp.ErrFailure("questions_not_found"))
		return
	}

	answers, err := app.Services.PointAnswer.GetPointAnswer()
	if err != nil {
		HttpError(c, headers, fmt.Errorf("Answers Not Found"), ahttp.ErrFailure("answers_not_found"))
		return
	}

	result := response.QuestionTest{}
	result.Questions = []entities.QuestionTest{}
	result.Answers = []entities.PointAnswer{}
	result.Total = 0

	if len(questionsTest) > 0 {
		result.Questions = questionsTest
		result.Answers = answers
		result.Total = len(questionsTest)
	}

	HttpSuccess(c, headers, result)
	return
}
