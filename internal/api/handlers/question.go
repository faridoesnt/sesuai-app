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
	ElementId    string         `json:"element_id"`
	ElementName  string         `json:"element_name"`
	ElementImage string         `json:"element_image"`
	QuestionList []QuestionList `json:"question_list"`
}

type ResponseAllQuestions struct {
	QuestionList []QuestionList `json:"question_list"`
}

func GetQuestions(c iris.Context) {
	headers := helpers.GetHeaders(c)

	adminId := c.Values().GetString(constants.AuthUserId)

	hasAccess, err := app.Services.AccessMenu.IsAdminHasAccessMenu(adminId, constants.QuestionList)
	if err != nil {
		HttpError(c, headers, fmt.Errorf(err.Error()), ahttp.ErrFailure(err.Error()))
		return
	}

	if !hasAccess {
		HttpError(c, headers, fmt.Errorf("admin doesn't have access"), ahttp.ErrFailure("admin_doesn't_have_access"))
		return
	}

	elementId := c.FormValue("element_id")

	elements := []response.Element{}
	result := []ResponseQuestions{}

	if elementId == "" {
		elements = app.Services.Element.GetElements()
	} else {
		element := app.Services.Element.GetElementDetail(elementId)

		if element.Id != "" {
			elements = append(elements, response.Element{
				Id:    element.Id,
				Name:  element.Name,
				Photo: element.Photo,
			})
		}
	}

	data := make(map[string]interface{})
	data["questions"] = result

	if len(elements) > 0 {
		for _, element := range elements {
			questions := app.Services.Question.GetQuestionsByElementId(element.Id)
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
					ElementId:    element.Id,
					ElementName:  element.Name,
					ElementImage: element.Photo,
					QuestionList: questionList,
				})
			}
		}

		data["questions"] = result
	}

	HttpSuccess(c, headers, data)
}

func GetAllQuestionsByElementId(c iris.Context) {
	headers := helpers.GetHeaders(c)

	adminId := c.Values().GetString(constants.AuthUserId)

	hasAccess, err := app.Services.AccessMenu.IsAdminHasAccessMenu(adminId, constants.QuestionList)
	if err != nil {
		HttpError(c, headers, fmt.Errorf(err.Error()), ahttp.ErrFailure(err.Error()))
		return
	}

	if !hasAccess {
		HttpError(c, headers, fmt.Errorf("admin doesn't have access"), ahttp.ErrFailure("admin_doesn't_have_access"))
		return
	}

	elementId := c.Params().GetString("elementId")

	if elementId == "" {
		HttpError(c, headers, ahttp.Error{Message: "Element Id Is Empty"}, ahttp.ErrFailure("element_id_is_empty"))
		return
	}

	if existElement := app.Services.Element.IsExistElement(elementId); !existElement {
		HttpError(c, headers, ahttp.Error{Message: "Element Not Found"}, ahttp.ErrFailure("element_not_found"))
		return
	}

	result := ResponseAllQuestions{}

	questions, err := app.Services.Question.GetAllQuestionsByElementId(elementId)
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

		result.QuestionList = questionList
	} else {
		HttpError(c, headers, fmt.Errorf("Questions Not Found"), ahttp.ErrFailure("questions_not_found"))
		return
	}

	HttpSuccess(c, headers, result)
}

func GetQuestion(c iris.Context) {
	headers := helpers.GetHeaders(c)

	adminId := c.Values().GetString(constants.AuthUserId)

	hasAccess, err := app.Services.AccessMenu.IsAdminHasAccessMenu(adminId, constants.QuestionList)
	if err != nil {
		HttpError(c, headers, fmt.Errorf(err.Error()), ahttp.ErrFailure(err.Error()))
		return
	}

	if !hasAccess {
		HttpError(c, headers, fmt.Errorf("admin doesn't have access"), ahttp.ErrFailure("admin_doesn't_have_access"))
		return
	}

	questionId := c.Params().GetString("questionId")

	if questionId == "" {
		HttpError(c, headers, ahttp.Error{Message: "Question Id Is Empty"}, ahttp.ErrFailure("question_id_is_empty"))
		return
	}

	question := app.Services.Question.GetQuestion(questionId)
	if question.Id == "" {
		HttpError(c, headers, ahttp.Error{Message: "Question Not Found"}, ahttp.ErrNotFound("question_not_found"))
		return
	}

	data := make(map[string]interface{})
	data["id_question"] = question.Id
	data["element_id"] = question.ElementId
	data["element_name"] = question.Element
	data["photo"] = question.Photo
	data["question_ina"] = question.QuestionIna
	data["question_eng"] = question.QuestionEn

	HttpSuccess(c, headers, data)
	return
}

func SaveQuestion(c iris.Context) {
	headers := helpers.GetHeaders(c)

	adminId := c.Values().GetString(constants.AuthUserId)

	hasAccess, err := app.Services.AccessMenu.IsAdminHasAccessMenu(adminId, constants.QuestionList)
	if err != nil {
		HttpError(c, headers, fmt.Errorf(err.Error()), ahttp.ErrFailure(err.Error()))
		return
	}

	if !hasAccess {
		HttpError(c, headers, fmt.Errorf("admin doesn't have access"), ahttp.ErrFailure("admin_doesn't_have_access"))
		return
	}

	params := entities.RequestQuestion{}

	err = c.ReadJSON(&params)
	if err != nil {
		HttpError(c, headers, err, ahttp.ErrInvalid(err.Error()))
		return
	}

	params.AdminId = adminId

	if params.ElementId == "" {
		HttpError(c, headers, ahttp.Error{Message: "Element Id Is Empty"}, ahttp.ErrFailure("element_id_is_empty"))
		return
	}

	if existElement := app.Services.Element.IsExistElement(params.ElementId); !existElement {
		HttpError(c, headers, ahttp.Error{Message: "Element Not Found"}, ahttp.ErrFailure("element_not_found"))
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

	adminId := c.Values().GetString(constants.AuthUserId)

	hasAccess, err := app.Services.AccessMenu.IsAdminHasAccessMenu(adminId, constants.QuestionList)
	if err != nil {
		HttpError(c, headers, fmt.Errorf(err.Error()), ahttp.ErrFailure(err.Error()))
		return
	}

	if !hasAccess {
		HttpError(c, headers, fmt.Errorf("admin doesn't have access"), ahttp.ErrFailure("admin_doesn't_have_access"))
		return
	}

	questionId := c.Params().GetString("questionId")

	params := entities.RequestQuestion{}

	err = c.ReadJSON(&params)
	if err != nil {
		HttpError(c, headers, fmt.Errorf(err.Error()), ahttp.ErrFailure(err.Error()))
		return
	}

	params.AdminId = adminId

	if questionId == "" {
		HttpError(c, headers, ahttp.Error{Message: "Element Id Is Empty"}, ahttp.ErrFailure("element_id_is_empty"))
		return
	}

	if existQuestion := app.Services.Question.IsExistQuestion(questionId); !existQuestion {
		HttpError(c, headers, ahttp.Error{Message: "Question Not Found"}, ahttp.ErrFailure("question_not_found"))
		return
	}

	if params.ElementId == "" {
		HttpError(c, headers, ahttp.Error{Message: "Element Id Is Empty"}, ahttp.ErrFailure("element_id_is_empty"))
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

	adminId := c.Values().GetString(constants.AuthUserId)

	hasAccess, err := app.Services.AccessMenu.IsAdminHasAccessMenu(adminId, constants.QuestionList)
	if err != nil {
		HttpError(c, headers, fmt.Errorf(err.Error()), ahttp.ErrFailure(err.Error()))
		return
	}

	if !hasAccess {
		HttpError(c, headers, fmt.Errorf("admin doesn't have access"), ahttp.ErrFailure("admin_doesn't_have_access"))
		return
	}

	questionId := c.Params().GetString("questionId")

	if questionId == "" {
		HttpError(c, headers, ahttp.Error{Message: "Element Id Is Empty"}, ahttp.ErrFailure("element_id_is_empty"))
		return
	}

	if existQuestion := app.Services.Question.IsExistQuestion(questionId); !existQuestion {
		HttpError(c, headers, ahttp.Error{Message: "Question Not Found"}, ahttp.ErrFailure("question_not_found"))
		return
	}

	err = app.Services.Question.DeleteQuestion(questionId)
	if err != nil {
		HttpError(c, headers, ahttp.Error{Message: "error delete question"}, ahttp.ErrFailure(err.Error()))
		return
	}

	HttpSuccess(c, headers, nil)
	return
}
