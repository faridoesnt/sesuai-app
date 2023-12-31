package generateToken

import (
	"Sesuai/internal/api/constracts"
	"Sesuai/internal/api/entities"
	"Sesuai/internal/api/models/response"
)

type Service struct {
	app  constracts.App
	repo constracts.GenerateTokenRepository
}

func Init(a *constracts.App) (svc constracts.GenerateTokenService) {
	r := initRepository(a.Datasources.WriterDB, a.Datasources.ReaderDB)

	svc = &Service{
		app:  *a,
		repo: r,
	}

	return
}

func (s Service) GetGenerateToken(adminId string) (listToken []response.GenerateToken) {
	tokens, _ := s.repo.FindGenerateToken(adminId)

	if len(tokens) > 0 {
		for _, token := range tokens {
			listToken = append(listToken, response.GenerateToken{
				Id:        token.Id,
				Admin:     token.Admin,
				Token:     token.Token,
				Status:    token.Status,
				CreatedAt: token.CreatedAt,
			})
		}
	} else {
		listToken = []response.GenerateToken{}
	}

	return
}

func (s Service) InsertNewToken(adminId, token string) (err error) {
	err = s.repo.InsertNewToken(adminId, token)

	return
}

func (s Service) ToggleInactiveToken(tokenId string) (err error) {
	err = s.repo.ToggleInactiveToken(tokenId)

	return
}

func (s Service) GetGenerateTokenByToken(params string) (token entities.GenerateToken, err error) {
	token, err = s.repo.FindGenerateTokenByToken(params)

	return
}
