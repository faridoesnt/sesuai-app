package constracts

import "Sesuai/internal/api/entities"

type QuestionRepository interface {
	InsertQuestion(params entities.RequestQuestion) (err error)
	FindQuestions(category string) (questions []entities.Question, err error)
	FindQuestion(categoryId string) (question entities.Question, err error)
	DeleteQuestion(categoryId string) (err error)
}

type QuestionService interface {
	GetQuestions(category string) (questions []entities.Question)
	GetQuestion(categoryId string) (question entities.Question)
	InsertQuestion(params entities.RequestQuestion) (err error)
	IsExistQuestion(questionId string) bool
	DeleteQuestion(questionId string) (err error)
}
