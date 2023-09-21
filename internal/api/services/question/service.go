package question

import (
	"Sesuai/internal/api/constracts"
	"Sesuai/internal/api/entities"
)

type Service struct {
	app  constracts.App
	repo constracts.QuestionRepository
}

func Init(a *constracts.App) (svc constracts.QuestionService) {
	r := initRepository(a.Datasources.WriterDB, a.Datasources.ReaderDB)

	svc = &Service{
		app:  *a,
		repo: r,
	}

	return
}

func (s Service) GetQuestions(category string) (questions []entities.Question) {
	questions, _ = s.repo.FindQuestions(category)

	return
}

func (s Service) GetQuestion(categoryId string) (question entities.Question) {
	question, _ = s.repo.FindQuestion(categoryId)

	return
}

func (s Service) InsertQuestion(params entities.RequestQuestion) (err error) {
	err = s.repo.InsertQuestion(params)

	return
}

func (s Service) IsExistQuestion(questionId string) bool {
	_, err := s.repo.FindQuestion(questionId)
	if err != nil {
		return false
	}

	return true
}

func (s Service) DeleteQuestion(questionId string) (err error) {
	err = s.repo.DeleteQuestion(questionId)

	return
}
