package constracts

import "Sesuai/internal/api/entities"

type QuestionRepository interface {
	InsertQuestion(params entities.RequestQuestion) (err error)
	FindQuestionsByElementId(elementId string) (questions []entities.Question, err error)
	FindAllQuestionsByElementId(elementId string) (questions []entities.Question, err error)
	FindQuestion(elementId string) (question entities.Question, err error)
	UpdateQuestion(questionId string, params entities.RequestQuestion) (err error)
	DeleteQuestion(elementId string) (err error)
}

type QuestionService interface {
	GetQuestionsByElementId(elementId string) (questions []entities.Question)
	GetAllQuestionsByElementId(elementId string) (questions []entities.Question, err error)
	GetQuestion(elementId string) (question entities.Question)
	InsertQuestion(params entities.RequestQuestion) (err error)
	IsExistQuestion(questionId string) bool
	UpdateQuestion(questionId string, params entities.RequestQuestion) (err error)
	DeleteQuestion(questionId string) (err error)
}
