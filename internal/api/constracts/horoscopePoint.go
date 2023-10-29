package constracts

import (
	"Sesuai/internal/api/entities"
	"Sesuai/internal/api/models/response"
)

type HoroscopePointRepository interface {
	FindHoroscopePoint(elementId string) (horoscopePoint []entities.HoroscopePoint, err error)
	UpdateHoroscopePoint(params entities.RequestHoroscopePoint) (err error)
	FindHoroscopePointByIdAndElementId(horoscopeId, elementId string) (horoscopePoint entities.HoroscopePoint, err error)
}

type HoroscopePointService interface {
	GetHoroscopePoint(elementId string) (horoscopePoint []response.HoroscopePoint, err error)
	UpdateHoroscopePoint(params entities.RequestHoroscopePoint) (err error)
	GetPointHoroscopeByIdAndElementId(horoscopeId, elementId string) (pointHoroscope float64, err error)
}
