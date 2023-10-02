package constracts

import "Sesuai/internal/api/entities"

type BloodTypePointRepository interface {
	FindBloodTypePoint(categoryId string) (bloodTypePoint []entities.BloodTypePoint, err error)
	UpdateBloodTypePoint(params entities.RequestBloodTypePoint) (err error)
}

type BloodTypePointService interface {
	GetBloodTypePoint(categoryId string) (bloodTypePoint []entities.BloodTypePoint, err error)
	UpdateBloodTypePoint(params entities.RequestBloodTypePoint) (err error)
}
