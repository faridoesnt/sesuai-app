package pointAnswer

import (
	"Sesuai/internal/api/constracts"
	"Sesuai/internal/api/entities"
)

type Service struct {
	app  constracts.App
	repo constracts.PointAnswerRepository
}

func Init(a *constracts.App) (svc constracts.PointAnswerService) {
	r := initRepository(a.Datasources.WriterDB, a.Datasources.ReaderDB)

	svc = &Service{
		app:  *a,
		repo: r,
	}

	return
}

func (s Service) GetPointAnswer() (pointAnswer []entities.PointAnswer, err error) {
	pointAnswer, err = s.repo.FindPointAnswer()

	return
}

func (s Service) UpdatePointAnswer(params entities.RequestPointAnswer) (err error) {
	err = s.repo.UpdatePointAnswer(params)

	return
}
