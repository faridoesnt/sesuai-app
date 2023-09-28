package constracts

import (
	"Sesuai/internal/api/entities"
	"Sesuai/internal/api/models/response"
)

type HoroscopeRepository interface {
	FindHoroscopeByName(horoscopeName string) (horoscope entities.Horoscope, err error)
}

type HoroscopeService interface {
	GetHoroscopeByName(horoscopeName string) (horoscope response.Horoscope, err error)
}
