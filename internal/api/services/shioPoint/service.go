package shioPoint

import (
	"Sesuai/internal/api/constracts"
	"Sesuai/internal/api/entities"
	"strconv"
)

type Service struct {
	app  constracts.App
	repo constracts.ShioPointRepository
}

func Init(a *constracts.App) (svc constracts.ShioPointService) {
	r := initRepository(a.Datasources.WriterDB, a.Datasources.ReaderDB)

	svc = &Service{
		app:  *a,
		repo: r,
	}

	return
}

func (s Service) GetShioPoint(elementId string) (shioPoint []entities.ShioPoint, err error) {
	shioPoint, err = s.repo.FindShioPoint(elementId)

	return
}

func (s Service) UpdateShioPoint(params entities.RequestShioPoint) (err error) {
	err = s.repo.UpdateShioPoint(params)

	return
}

func (s Service) GetPointShioByIdAndElementId(shioId, elementId string) (pointShio float64, err error) {
	shioPoint, err := s.repo.FindShioPointByIdAndElementId(shioId, elementId)

	pointShio, _ = strconv.ParseFloat(shioPoint.Point, 64)

	return
}
