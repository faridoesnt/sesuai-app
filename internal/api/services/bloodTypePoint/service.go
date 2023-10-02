package bloodTypePoint

import (
	"Sesuai/internal/api/constracts"
	"Sesuai/internal/api/entities"
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

func (s Service) GetBloodTypePoint(categoryId string) (bloodTypePoint []entities.BloodTypePoint, err error) {
	bloodTypePoint, err = s.repo.FindBloodTypePoint(categoryId)

	return
}
