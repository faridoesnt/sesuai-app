package adminPhoneNumber

import (
	"Sesuai/internal/api/constracts"
	"Sesuai/internal/api/entities"
)

type Service struct {
	app  constracts.App
	repo constracts.AdminPhoneNumberRepository
}

func Init(a *constracts.App) (svc constracts.AdminPhoneNumberService) {
	r := initRepository(a.Datasources.WriterDB, a.Datasources.ReaderDB)

	svc = &Service{
		app:  *a,
		repo: r,
	}

	return
}

func (s Service) GetAdminPhoneNumber() (adminPhoneNumber entities.AdminPhoneNumber, err error) {
	adminPhoneNumber, err = s.repo.FindAdminPhoneNumber()

	return
}
