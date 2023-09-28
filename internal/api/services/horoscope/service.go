package horoscope

import (
	"Sesuai/internal/api/constracts"
	"Sesuai/internal/api/models/response"
)

type Service struct {
	app  constracts.App
	repo constracts.HoroscopeRepository
}

func Init(a *constracts.App) (svc constracts.HoroscopeService) {
	r := initRepository(a.Datasources.WriterDB, a.Datasources.ReaderDB)

	svc = &Service{
		app:  *a,
		repo: r,
	}

	return
}

func (s Service) GetHoroscopeByName(horoscopeName string) (horoscope response.Horoscope, err error) {
	data, err := s.repo.FindHoroscopeByName(horoscopeName)

	horoscope.Id = data.Id
	horoscope.Name = data.Name

	return
}
