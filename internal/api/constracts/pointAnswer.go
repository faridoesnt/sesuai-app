package constracts

import "Sesuai/internal/api/entities"

type PointAnswerRepository interface {
	FindPointAnswer() (pointAnswer []entities.PointAnswer, err error)
}

type PointAnswerService interface {
	GetPointAnswer() (pointAnswer []entities.PointAnswer, err error)
}
