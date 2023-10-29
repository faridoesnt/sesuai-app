package constracts

import "Sesuai/internal/api/entities"

type ResultRepository interface {
	FindResult(userId string) (results []entities.Result, err error)
}

type ResultService interface {
	GetResult(userId string) (results []entities.Result, err error)
}
