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

func (s Service) GetQuestionsByCategoryId(categoryId string) (questions []entities.Question) {
	questions, _ = s.repo.FindQuestionsByCategoryId(categoryId)

	return
}

func (s Service) GetAllQuestionsByCategoryId(categoryId string) (questions []entities.Question, err error) {
	questions, err = s.repo.FindAllQuestionsByCategoryId(categoryId)

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

func (s Service) UpdateQuestion(questionId string, params entities.RequestQuestion) (err error) {
	err = s.repo.UpdateQuestion(questionId, params)

	return
}

func (s Service) DeleteQuestion(questionId string) (err error) {
	err = s.repo.DeleteQuestion(questionId)

	return
}
