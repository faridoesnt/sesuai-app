package questionTest

import (
	"Sesuai/internal/api/constracts"
	"Sesuai/internal/api/entities"
)

type Service struct {
	app  constracts.App
	repo constracts.QuestionTestRepository
}

func Init(a *constracts.App) (svc constracts.QuestionTestService) {
	r := initRepository(a.Datasources.WriterDB, a.Datasources.ReaderDB)

	svc = &Service{
		app:  *a,
		repo: r,
	}

	return
}

func (s Service) GetQuestionsTest() (questionsTest []entities.QuestionTest, err error) {
	questionsTest, err = s.repo.FindQuestionsTest()

	return
}

func (s Service) SubmitQuestionTest(params entities.SubmitQuestionTest, userId string, totalPointQuestionsByElement map[string]float64) (err error) {
	err = s.repo.SubmitQuestionTest(params, userId, totalPointQuestionsByElement)

	return
}

func (s Service) CheckQuestionTestUser(userId string) (isExist bool, err error) {
	count, err := s.repo.CountQuestionTestUser(userId)

	return count > 0, err
}
