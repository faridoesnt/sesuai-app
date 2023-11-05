package horoscope

import (
	"Sesuai/internal/api/constracts"
	"Sesuai/internal/api/entities"
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

func (s Service) GetHoroscopes() (horoscopes []response.Horoscope, err error) {
	horoscope, err := s.repo.FindHoroscopes()

	if len(horoscope) > 0 {
		for _, val := range horoscope {
			horoscopes = append(horoscopes, response.Horoscope{
				Id:   val.Id,
				Name: val.Name,
			})
		}
	} else {
		horoscopes = []response.Horoscope{}
	}

	return
}

func (s Service) GetHoroscopeByName(horoscopeName string) (horoscope response.Horoscope, err error) {
	data, err := s.repo.FindHoroscopeByName(horoscopeName)

	horoscope.Id = data.Id
	horoscope.Name = data.Name

	return
}

func (s Service) GetHoroscopeUser(userId string) (horoscope entities.Horoscope, err error) {
	horoscope, err = s.repo.FindHoroscopeUser(userId)

	return
}

func (s Service) IsHoroscopeExist(horoscopeId string) bool {
	count, _ := s.repo.CountHoroscopeById(horoscopeId)

	return count > 0
}
