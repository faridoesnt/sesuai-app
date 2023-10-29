package constracts

import (
	"Sesuai/internal/api/entities"
	"Sesuai/internal/api/models/response"
)

type ShioRepository interface {
	FindShio() (shio []entities.Shio, err error)
	FindShioUser(userId string) (shio entities.Shio, err error)
}

type ShioService interface {
	GetShio() (shio []response.Shio, err error)
	GetShioUser(userId string) (shio entities.Shio, err error)
}
