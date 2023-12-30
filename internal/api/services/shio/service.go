package shio

import (
	"Sesuai/internal/api/constracts"
	"Sesuai/internal/api/entities"
	"Sesuai/internal/api/models/response"
)

type Service struct {
	app  constracts.App
	repo constracts.ShioRepository
}

func Init(a *constracts.App) (svc constracts.ShioService) {
	r := initRepository(a.Datasources.WriterDB, a.Datasources.ReaderDB)

	svc = &Service{
		app:  *a,
		repo: r,
	}

	return
}

func (s Service) GetShio() (shio []response.Shio, err error) {
	listShio, err := s.repo.FindShio()

	for _, val := range listShio {
		shio = append(shio, response.Shio{
			Id:   val.Id,
			Name: val.Name,
		})
	}

	return
}

func (s Service) GetShioByShioName(shioName string) (shio entities.Shio, err error) {
	shio, err = s.repo.FindShioByShioName(shioName)

	return
}

func (s Service) GetShioUser(userId string) (shio entities.Shio, err error) {
	shio, err = s.repo.FindShioUser(userId)

	return
}

func (s Service) IsShioExist(shioId string) bool {
	count, _ := s.repo.CountShioById(shioId)

	return count > 0
}
