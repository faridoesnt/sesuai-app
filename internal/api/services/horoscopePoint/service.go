package horoscopePoint

import (
	"Sesuai/internal/api/constracts"
	"Sesuai/internal/api/entities"
	"Sesuai/internal/api/models/response"
)

type Service struct {
	app  constracts.App
	repo constracts.HoroscopePointRepository
}

func Init(a *constracts.App) (svc constracts.HoroscopePointService) {
	r := initRepository(a.Datasources.WriterDB, a.Datasources.ReaderDB)

	svc = &Service{
		app:  *a,
		repo: r,
	}

	return
}

func (s Service) GetHoroscopePoint(categoryId string) (horoscopePoint []response.HoroscopePoint, err error) {
	horoscopesPoint, err := s.repo.FindHoroscopePoint(categoryId)

	if len(horoscopesPoint) > 0 {
		for _, horoscope := range horoscopesPoint {
			horoscopePoint = append(horoscopePoint, response.HoroscopePoint{
				HoroscopeId:   horoscope.HoroscopeId,
				CategoryId:    horoscope.CategoryId,
				HoroscopeName: horoscope.HoroscopeName,
				Point:         horoscope.Point,
			})
		}
	} else {
		horoscopePoint = []response.HoroscopePoint{}
	}

	return
}

func (s Service) UpdateHoroscopePoint(params entities.RequestHoroscopePoint) (err error) {
	err = s.repo.UpdateHoroscopePoint(params)

	return
}
