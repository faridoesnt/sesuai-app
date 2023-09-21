package constracts

import (
	"Sesuai/internal/api/entities"
	"Sesuai/internal/api/models/response"
)

type GenerateTokenRepository interface {
	FindGenerateToken(adminId string) (tokens []entities.GenerateToken, err error)
	InsertNewToken(adminId, token string) (err error)
	UpdateToken(tokenId string) (err error)
}

type GenerateTokenService interface {
	GetGenerateToken(adminId string) (listToken []response.GenerateToken)
	InsertNewToken(adminId, token string) (err error)
	UpdateToken(tokenId string) (err error)
}
