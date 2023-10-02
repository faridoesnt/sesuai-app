package constracts

import "Sesuai/internal/api/entities"

type PointAnswerRepository interface {
	FindPointAnswer() (pointAnswer []entities.PointAnswer, err error)
	UpdatePointAnswer(params entities.RequestPointAnswer) (err error)
}

type PointAnswerService interface {
	GetPointAnswer() (pointAnswer []entities.PointAnswer, err error)
	UpdatePointAnswer(params entities.RequestPointAnswer) (err error)
}
