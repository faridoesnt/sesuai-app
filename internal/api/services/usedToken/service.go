package usedToken

import "Sesuai/internal/api/constracts"

type Service struct {
	app  constracts.App
	repo constracts.UsedTokenRepository
}

func Init(a *constracts.App) (svc constracts.UsedTokenService) {
	r := initRepository(a.Datasources.WriterDB, a.Datasources.ReaderDB)

	svc = &Service{
		app:  *a,
		repo: r,
	}

	return
}

func (s Service) InsertUsedToken(tokenId, userId string) (err error) {
	err = s.repo.InsertUsedToken(tokenId, userId)

	return
}

func (s Service) IsUserToken(token, userId string) (isUserToken bool, err error) {
	total, err := s.repo.CountUserToken(token, userId)

	return total > 0, err
}

func (s Service) GetUsedTokenByUserId(userId string) (token string, err error) {
	token, err = s.repo.FindUsedTokenByUserId(userId)

	return
}
