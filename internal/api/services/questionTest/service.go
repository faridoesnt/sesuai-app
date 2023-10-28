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
