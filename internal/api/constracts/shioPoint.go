package constracts

import "Sesuai/internal/api/entities"

type ShioPointRepository interface {
	FindShioPoint(elementId string) (shioPoint []entities.ShioPoint, err error)
	UpdateShioPoint(params entities.RequestShioPoint) (err error)
	FindShioPointByIdAndElementId(shioId, elementId string) (shioPoint entities.ShioPoint, err error)
}

type ShioPointService interface {
	GetShioPoint(elementId string) (shioPoint []entities.ShioPoint, err error)
	UpdateShioPoint(params entities.RequestShioPoint) (err error)
	GetPointShioByIdAndElementId(shioId, elementId string) (pointShio float64, err error)
}
