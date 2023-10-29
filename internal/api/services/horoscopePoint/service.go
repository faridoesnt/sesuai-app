package horoscopePoint

import (
	"Sesuai/internal/api/constracts"
	"Sesuai/internal/api/entities"
	"Sesuai/internal/api/models/response"
	"strconv"
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

func (s Service) GetHoroscopePoint(elementId string) (horoscopePoint []response.HoroscopePoint, err error) {
	horoscopesPoint, err := s.repo.FindHoroscopePoint(elementId)

	if len(horoscopesPoint) > 0 {
		for _, horoscope := range horoscopesPoint {
			horoscopePoint = append(horoscopePoint, response.HoroscopePoint{
				HoroscopeId:   horoscope.HoroscopeId,
				ElementId:     horoscope.ElementId,
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

func (s Service) GetPointHoroscopeByIdAndElementId(horoscopeId, elementId string) (pointHoroscope float64, err error) {
	horoscopePoint, err := s.repo.FindHoroscopePointByIdAndElementId(horoscopeId, elementId)

	pointHoroscope, _ = strconv.ParseFloat(horoscopePoint.Point, 64)

	return
}
