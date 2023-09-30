package constracts

import (
	"Sesuai/internal/api/entities"
	"Sesuai/internal/api/models/response"
)

type HoroscopePointRepository interface {
	FindHoroscopePoint(categoryId string) (horoscopePoint []entities.HoroscopePoint, err error)
	UpdateHoroscopePoint(params entities.RequestHoroscopePoint) (err error)
}

type HoroscopePointService interface {
	GetHoroscopePoint(categoryId string) (horoscopePoint []response.HoroscopePoint, err error)
	UpdateHoroscopePoint(params entities.RequestHoroscopePoint) (err error)
}
