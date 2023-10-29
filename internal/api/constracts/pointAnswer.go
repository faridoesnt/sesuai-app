package constracts

import "Sesuai/internal/api/entities"

type PointAnswerRepository interface {
	FindPointAnswer() (pointAnswer []entities.PointAnswer, err error)
	UpdatePointAnswer(params entities.RequestPointAnswer) (err error)
	FindPointAnswerById(answerId string) (pointAnswer entities.PointAnswer, err error)
}

type PointAnswerService interface {
	GetPointAnswer() (pointAnswer []entities.PointAnswer, err error)
	UpdatePointAnswer(params entities.RequestPointAnswer) (err error)
	GetPointAnswerById(answerId string) (pointAnswer entities.PointAnswer, err error)
}
