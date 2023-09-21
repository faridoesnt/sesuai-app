package handlers

import (
	"Sesuai/internal/api/constants"
	"Sesuai/internal/api/entities"
	"Sesuai/internal/api/helpers"
	"Sesuai/pkg/ahttp"
	"Sesuai/pkg/autils"
	"errors"
	"github.com/kataras/iris/v12"
)

func GetQuestions(c iris.Context) {
	headers := helpers.GetHeaders(c)

	category := c.FormValue("category")

	data := make(map[string]interface{})

	setCategoryData := func(category string) {

		questions := app.Services.Question.GetQuestions(category)

		if len(questions) > 0 {
			categoryData := make([]map[string]interface{}, len(questions))
			for i, question := range questions {
				categoryData[i] = map[string]interface{}{
					"id":           question.Id,
					"question_ina": question.QuestionIna,
					"question_eng": question.QuestionEn,
				}
			}

			result := map[string]interface{}{
				"photo": questions[0].Photo,
				"data":  categoryData,
			}

			data[autils.UncapitalizedInitialLetter(category)] = result
		}
	}

	if category == "" {
		allCategories := app.Services.Category.GetCategory()

		for _, category := range allCategories {
			setCategoryData(category.Name)
		}
	} else {
		setCategoryData(category)
	}

	HttpSuccess(c, headers, data)
	return
}

func GetQuestion(c iris.Context) {
	headers := helpers.GetHeaders(c)

	questionId := c.Params().GetString("questionId")

	if questionId == "" {
		HttpError(c, headers, ahttp.Error{Message: "Question Id Is Empty"}, ahttp.ErrFailure("Question Id Is Empty"))
		return
	}

	question := app.Services.Question.GetQuestion(questionId)
	if question.Id == "" {
		HttpError(c, headers, ahttp.Error{Message: "Question Not Found"}, ahttp.ErrNotFound("Question Not Found"))
		return
	}

	data := make(map[string]interface{})
	data["id_question"] = question.Id
	data["id_category"] = question.CategoryId
	data["category"] = question.Category
	data["photo"] = question.Photo
	data["question_ina"] = question.QuestionIna
	data["question_eng"] = question.QuestionEn

	HttpSuccess(c, headers, data)
	return
}

func SaveQuestion(c iris.Context) {
	headers := helpers.GetHeaders(c)

	adminId := c.Values().GetString(constants.AuthUserId)

	params := entities.RequestQuestion{}

	err := c.ReadJSON(&params)
	if err != nil {
		HttpError(c, headers, err, ahttp.ErrInvalid(err.Error()))
		return
	}

	params.AdminId = adminId

	if existCategory := app.Services.Category.IsExistCategory(params.CategoryId); !existCategory {
		HttpError(c, headers, ahttp.Error{Message: "Category Not Found"}, ahttp.ErrFailure("Category Not Found"))
		return
	}

	err = app.Services.Question.InsertQuestion(params)
	if err != nil {
		HttpError(c, headers, errors.New("error save question"), ahttp.ErrFailure(err.Error()))
		return
	}

	HttpSuccess(c, headers, nil)
	return
}

func DeleteQuestion(c iris.Context) {
	headers := helpers.GetHeaders(c)

	questionId := c.Params().GetString("questionId")

	if questionId == "" {
		HttpError(c, headers, ahttp.Error{Message: "Question Id Is Empty"}, ahttp.ErrFailure("Question Id Is Empty"))
		return
	}

	if existQuestion := app.Services.Question.IsExistQuestion(questionId); !existQuestion {
		HttpError(c, headers, ahttp.Error{Message: "Question Not Found"}, ahttp.ErrFailure("Question Not Found"))
		return
	}

	err := app.Services.Question.DeleteQuestion(questionId)
	if err != nil {
		HttpError(c, headers, ahttp.Error{Message: "error delete question"}, ahttp.ErrFailure(err.Error()))
		return
	}

	HttpSuccess(c, headers, nil)
	return
}
