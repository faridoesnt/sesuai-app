package constracts

import (
	"Sesuai/internal/api/entities"
	"Sesuai/internal/api/models/response"
)

type GenerateTokenRepository interface {
	FindGenerateToken(adminId string) (tokens []entities.GenerateToken, err error)
	InsertNewToken(adminId, token string) (err error)
	ToggleInactiveToken(tokenId string) (err error)
	FindGenerateTokenByToken(params string) (token entities.GenerateToken, err error)
}

type GenerateTokenService interface {
	GetGenerateToken(adminId string) (listToken []response.GenerateToken)
	InsertNewToken(adminId, token string) (err error)
	ToggleInactiveToken(tokenId string) (err error)
	GetGenerateTokenByToken(params string) (token entities.GenerateToken, err error)
}
