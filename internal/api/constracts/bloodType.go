package constracts

import (
	"Sesuai/internal/api/entities"
	"Sesuai/internal/api/models/response"
)

type BloodTypeRepository interface {
	FindBloodType() (bloodType []entities.BloodType, err error)
	FindBloodTypeUser(userId string) (bloodType entities.BloodType, err error)
	CountBloodTypeById(bloodTypeId string) (count int64, err error)
}

type BloodTypeService interface {
	GetBloodType() (listBloodType []response.BloodType)
	GetBloodTypeUser(userId string) (bloodType entities.BloodType, err error)
	IsBloodTypeExist(bloodTypeId string) bool
}
