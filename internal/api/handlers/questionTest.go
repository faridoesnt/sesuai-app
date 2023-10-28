package handlers

import (
	"Sesuai/internal/api/entities"
	"Sesuai/internal/api/helpers"
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

	data := make(map[string]interface{})
	data["questions_list"] = []entities.QuestionTest{}
	data["total"] = 0

	if len(questionsTest) > 0 {
		data["questions_list"] = questionsTest
		data["total"] = len(questionsTest)
	}

	HttpSuccess(c, headers, data)
	return
}
