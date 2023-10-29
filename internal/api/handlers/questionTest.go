package handlers

import (
	"Sesuai/internal/api/constants"
	"Sesuai/internal/api/entities"
	"Sesuai/internal/api/helpers"
	"Sesuai/internal/api/models/response"
	"Sesuai/pkg/ahttp"
	"fmt"
	"github.com/kataras/iris/v12"
	"strconv"
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

func SubmitQuestionTest(c iris.Context) {
	headers := helpers.GetHeaders(c)

	userId := c.Values().GetString(constants.AuthUserId)

	params := entities.SubmitQuestionTest{}

	err := c.ReadJSON(&params)
	if err != nil {
		HttpError(c, headers, err, ahttp.ErrInvalid(err.Error()))
		return
	}

	answerQuestionByElement := make(map[string][]string)

	if params.Timer == "" {
		HttpError(c, headers, fmt.Errorf("Timer Can't Empty"), ahttp.ErrFailure("timer_can't_empty"))
		return
	}

	for _, val := range params.Submit {
		if val.QuestionId == "" {
			HttpError(c, headers, fmt.Errorf("Question Id Can't Empty"), ahttp.ErrFailure("question_id_can't_empty"))
			return
		}

		if val.ElementId == "" {
			HttpError(c, headers, fmt.Errorf("Element Id Can't Empty"), ahttp.ErrFailure("element_id_can't_empty"))
			return
		}

		if val.AnswerId == "" {
			HttpError(c, headers, fmt.Errorf("Answer Id Can't Empty"), ahttp.ErrFailure("answer_id_can't_empty"))
			return
		}

		// grouping answer question by element
		answerQuestionByElement[val.ElementId] = append(answerQuestionByElement[val.ElementId], val.AnswerId)
	}

	shioUser, err := app.Services.Shio.GetShioUser(userId)
	if err != nil {
		HttpError(c, headers, fmt.Errorf("User Don't Have Shio"), ahttp.ErrFailure("user_don't_have_shio"))
		return
	}

	horoscopeUser, err := app.Services.Horoscope.GetHoroscopeUser(userId)
	if err != nil {
		HttpError(c, headers, fmt.Errorf("User Don't Have Horoscope"), ahttp.ErrFailure("user_don't_have_horoscope"))
		return
	}

	bloodTypeUser, err := app.Services.BloodType.GetBloodTypeUser(userId)
	if err != nil {
		HttpError(c, headers, fmt.Errorf("User Don't Have Blood Type"), ahttp.ErrFailure("user_don't_have_blood_type"))
		return
	}

	totalPointQuestionsByElement := make(map[string]float64)

	for elementId, answerId := range answerQuestionByElement {
		var totalPoint float64

		// add up point questions by element
		for _, val := range answerId {
			pointAnswer, err := app.Services.PointAnswer.GetPointAnswerById(val)
			if err != nil {
				HttpError(c, headers, fmt.Errorf("Point Answer Not Found"), ahttp.ErrFailure("point_answer_not_found"))
				return
			}

			point, _ := strconv.ParseFloat(pointAnswer.Point, 64)

			totalPoint += point
		}

		// total point divided total question by element
		totalPoint = totalPoint / float64(len(answerId))

		// add up total point with point shio
		pointShio, err := app.Services.ShioPoint.GetPointShioByIdAndElementId(shioUser.Id, elementId)
		if err != nil {
			HttpError(c, headers, fmt.Errorf(err.Error()), ahttp.ErrFailure(err.Error()))
			return
		}

		totalPoint += pointShio

		// add up total point with point horoscope
		pointHoroscope, err := app.Services.HoroscopePoint.GetPointHoroscopeByIdAndElementId(horoscopeUser.Id, elementId)
		if err != nil {
			HttpError(c, headers, fmt.Errorf(err.Error()), ahttp.ErrFailure(err.Error()))
			return
		}

		totalPoint += pointHoroscope

		// add up total point with point blood type
		bloodTypePoint, err := app.Services.BloodTypePoint.GetPointBloodTypeByIdAndElementId(bloodTypeUser.Id, elementId)
		if err != nil {
			HttpError(c, headers, fmt.Errorf(err.Error()), ahttp.ErrFailure(err.Error()))
			return
		}

		totalPoint += bloodTypePoint

		totalPointQuestionsByElement[elementId] = totalPoint
	}

	err = app.Services.QuestionTest.SubmitQuestionTest(params, userId, totalPointQuestionsByElement)
	if err != nil {
		HttpError(c, headers, fmt.Errorf(err.Error()), ahttp.ErrFailure(err.Error()))
		return
	}

	HttpSuccess(c, headers, nil)
	return
}
