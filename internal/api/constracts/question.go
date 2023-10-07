package constracts

import "Sesuai/internal/api/entities"

type QuestionRepository interface {
	InsertQuestion(params entities.RequestQuestion) (err error)
	FindQuestionsByCategoryId(categoryId string) (questions []entities.Question, err error)
	FindQuestion(categoryId string) (question entities.Question, err error)
	UpdateQuestion(questionId string, params entities.RequestQuestion) (err error)
	DeleteQuestion(categoryId string) (err error)
}

type QuestionService interface {
	GetQuestionsByCategoryId(categoryId string) (questions []entities.Question)
	GetQuestion(categoryId string) (question entities.Question)
	InsertQuestion(params entities.RequestQuestion) (err error)
	IsExistQuestion(questionId string) bool
	UpdateQuestion(questionId string, params entities.RequestQuestion) (err error)
	DeleteQuestion(questionId string) (err error)
}
