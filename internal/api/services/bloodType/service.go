package bloodType

import (
	"Sesuai/internal/api/constracts"
	"Sesuai/internal/api/entities"
	"Sesuai/internal/api/models/response"
)

type Service struct {
	app  constracts.App
	repo constracts.BloodTypeRepository
}

func Init(a *constracts.App) (svc constracts.BloodTypeService) {
	r := initRepository(a.Datasources.WriterDB, a.Datasources.ReaderDB)

	svc = &Service{
		app:  *a,
		repo: r,
	}

	return
}

func (s Service) GetBloodType() (listBloodType []response.BloodType) {
	bloodType, _ := s.repo.FindBloodType()

	if len(bloodType) > 0 {
		for _, val := range bloodType {
			listBloodType = append(listBloodType, response.BloodType{
				Id:   val.Id,
				Name: val.Name,
			})
		}
	} else {
		listBloodType = []response.BloodType{}
	}

	return
}

func (s Service) GetBloodTypeUser(userId string) (bloodType entities.BloodType, err error) {
	bloodType, err = s.repo.FindBloodTypeUser(userId)

	return
}

func (s Service) IsBloodTypeExist(bloodTypeId string) bool {
	count, _ := s.repo.CountBloodTypeById(bloodTypeId)

	return count > 0
}
