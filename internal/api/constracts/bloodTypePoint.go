package constracts

import "Sesuai/internal/api/entities"

type BloodTypePointRepository interface {
	FindBloodTypePoint(elementId string) (bloodTypePoint []entities.BloodTypePoint, err error)
	UpdateBloodTypePoint(params entities.RequestBloodTypePoint) (err error)
	FindBloodTypePointByIdAndElementId(bloodTypeId, elementId string) (bloodTypePoint entities.BloodTypePoint, err error)
}

type BloodTypePointService interface {
	GetBloodTypePoint(elementId string) (bloodTypePoint []entities.BloodTypePoint, err error)
	UpdateBloodTypePoint(params entities.RequestBloodTypePoint) (err error)
	GetPointBloodTypeByIdAndElementId(bloodTypeId, elementId string) (pointBloodType float64, err error)
}
