package constracts

import (
	"Sesuai/internal/api/entities"
	"Sesuai/internal/api/models/response"
)

type HoroscopeRepository interface {
	FindHoroscopes() (horoscopes []entities.Horoscope, err error)
	FindHoroscopeByName(horoscopeName string) (horoscope entities.Horoscope, err error)
	FindHoroscopeUser(userId string) (horoscope entities.Horoscope, err error)
	CountHoroscopeById(horoscopeId string) (count int64, err error)
}

type HoroscopeService interface {
	GetHoroscopes() (horoscopes []response.Horoscope, err error)
	GetHoroscopeByName(horoscopeName string) (horoscope response.Horoscope, err error)
	GetHoroscopeUser(userId string) (horoscope entities.Horoscope, err error)
	IsHoroscopeExist(horoscopeId string) bool
}
