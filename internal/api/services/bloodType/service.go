package bloodType

import (
	"Sesuai/internal/api/constracts"
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
