package logging

import (
	"Sesuai/internal/api/constracts"
	"Sesuai/internal/api/entities"
)

type Service struct {
	app  constracts.App
	repo constracts.LoggingRepository
}

func Init(a *constracts.App) (svc constracts.LoggingService) {
	r := initRepository(a.Datasources.WriterDB, a.Datasources.ReaderDB)

	svc = &Service{
		app:  *a,
		repo: r,
	}

	return
}

func (s Service) InsertMobileLogAdmin(mobileLog entities.MobileLogAdmin) (err error) {
	err = s.repo.InsertMobileLogAdmin(mobileLog)

	return
}

func (s Service) InsertMobileLogUser(mobileLog entities.MobileLogUser) (err error) {
	err = s.repo.InsertMobileLogUser(mobileLog)

	return
}
