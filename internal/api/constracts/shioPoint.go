package constracts

import "Sesuai/internal/api/entities"

type ShioPointRepository interface {
	FindShioPoint(categoryId string) (shioPoint []entities.ShioPoint, err error)
}

type ShioPointService interface {
	GetShioPoint(categoryId string) (shioPoint []entities.ShioPoint, err error)
}
