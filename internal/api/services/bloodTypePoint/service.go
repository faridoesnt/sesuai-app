package bloodTypePoint

import (
	"Sesuai/internal/api/constracts"
	"Sesuai/internal/api/entities"
	"strconv"
)

type Service struct {
	app  constracts.App
	repo constracts.BloodTypePointRepository
}

func Init(a *constracts.App) (svc constracts.BloodTypePointService) {
	r := initRepository(a.Datasources.WriterDB, a.Datasources.ReaderDB)

	svc = &Service{
		app:  *a,
		repo: r,
	}

	return
}

func (s Service) GetBloodTypePoint(elementId string) (bloodTypePoint []entities.BloodTypePoint, err error) {
	bloodTypePoint, err = s.repo.FindBloodTypePoint(elementId)

	return
}

func (s Service) UpdateBloodTypePoint(params entities.RequestBloodTypePoint) (err error) {
	err = s.repo.UpdateBloodTypePoint(params)

	return
}

func (s Service) GetPointBloodTypeByIdAndElementId(bloodTypeId, elementId string) (pointBloodType float64, err error) {
	bloodTypePoint, err := s.repo.FindBloodTypePointByIdAndElementId(bloodTypeId, elementId)

	pointBloodType, _ = strconv.ParseFloat(bloodTypePoint.Point, 64)

	return
}
