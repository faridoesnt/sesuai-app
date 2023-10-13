package handlers

import (
	"Sesuai/internal/api/constants"
	"Sesuai/internal/api/entities"
	"Sesuai/internal/api/helpers"
	"Sesuai/internal/api/models/response"
	"Sesuai/pkg/ahttp"
	"errors"
	"fmt"
	"github.com/kataras/iris/v12"
)

type QuestionList struct {
	QuestionId  string `json:"question_id"`
	QuestionIna string `json:"question_ina"`
	QuestionEng string `json:"question_eng"`
}

type ResponseQuestions struct {
	ElementName  string         `json:"element_name"`
	ElementImage string         `json:"element_image"`
	QuestionList []QuestionList `json:"question_list"`
}

func GetQuestions(c iris.Context) {
	headers := helpers.GetHeaders(c)

	categoryId := c.FormValue("element_id")

	categories := []response.Category{}
	result := []ResponseQuestions{}

	if categoryId == "" {
		categories = app.Services.Category.GetCategory()
	} else {
		category := app.Services.Category.GetCategoryDetail(categoryId)

		categories = append(categories, response.Category{
			Id:    category.Id,
			Name:  category.Name,
			Photo: category.Photo,
		})
	}

	if len(categories) == 0 {
		HttpError(c, headers, fmt.Errorf("no category found"), ahttp.ErrFailure("no_category_found"))
		return
	}

	for _, category := range categories {
		questions := app.Services.Question.GetQuestionsByCategoryId(category.Id)
		questionList := []QuestionList{}

		if len(questions) > 0 {
			for _, question := range questions {
				questionList = append(questionList, QuestionList{
					QuestionId:  question.Id,
					QuestionIna: question.QuestionIna,
					QuestionEng: question.QuestionEn,
				})
			}

			result = append(result, ResponseQuestions{
				ElementName:  category.Name,
				ElementImage: category.Photo,
				QuestionList: questionList,
			})
		}
	}

	data := make(map[string]interface{})
	data["questions"] = result

	HttpSuccess(c, headers, data)
}

func GetAllQuestionsByCategoryId(c iris.Context) {
	headers := helpers.GetHeaders(c)

	categoryId := c.Params().GetString("categoryId")

	if categoryId == "" {
		HttpError(c, headers, ahttp.Error{Message: "Category Id Is Empty"}, ahttp.ErrFailure("category_id_is_empty"))
		return
	}

	category := app.Services.Category.GetCategoryDetail(categoryId)
	result := []ResponseQuestions{}

	questions, err := app.Services.Question.GetAllQuestionsByCategoryId(categoryId)
	if err != nil {
		HttpError(c, headers, fmt.Errorf(err.Error()), ahttp.ErrFailure(err.Error()))
		return
	}

	questionList := []QuestionList{}

	if len(questions) > 0 {
		for _, question := range questions {
			questionList = append(questionList, QuestionList{
				QuestionId:  question.Id,
				QuestionIna: question.QuestionIna,
				QuestionEng: question.QuestionEn,
			})
		}

		result = append(result, ResponseQuestions{
			ElementName:  category.Name,
			ElementImage: category.Photo,
			QuestionList: questionList,
		})
	} else {
		HttpError(c, headers, fmt.Errorf("No Questions Found"), ahttp.ErrFailure("no_questions_found"))
		return
	}

	data := make(map[string]interface{})
	data["questions"] = result

	HttpSuccess(c, headers, data)
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

func UpdateQuestion(c iris.Context) {
	headers := helpers.GetHeaders(c)

	questionId := c.Params().GetString("questionId")

	adminId := c.Values().GetString(constants.AuthUserId)

	params := entities.RequestQuestion{}

	err := c.ReadJSON(&params)
	if err != nil {
		HttpError(c, headers, fmt.Errorf(err.Error()), ahttp.ErrFailure(err.Error()))
		return
	}

	params.AdminId = adminId

	if questionId == "" {
		HttpError(c, headers, ahttp.Error{Message: "Question Id Is Empty"}, ahttp.ErrFailure("question_id_is_empty"))
		return
	}

	if existQuestion := app.Services.Question.IsExistQuestion(questionId); !existQuestion {
		HttpError(c, headers, ahttp.Error{Message: "Question Not Found"}, ahttp.ErrFailure("question_not_found"))
		return
	}

	err = app.Services.Question.UpdateQuestion(questionId, params)
	if err != nil {
		HttpError(c, headers, ahttp.Error{Message: "error update question"}, ahttp.ErrFailure(err.Error()))
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
