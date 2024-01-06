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

func (s Service) InsertUsedToken(tokenId, submissionId string) (err error) {
	err = s.repo.InsertUsedToken(tokenId, submissionId)

	return
}

func (s Service) IsSubmissionToken(token, submissionId string) (isSubmissionToken bool, err error) {
	total, err := s.repo.CountSubmissionToken(token, submissionId)

	return total > 0, err
}

func (s Service) GetUsedTokenByUserId(userId string) (token string, err error) {
	token, err = s.repo.FindUsedTokenByUserId(userId)

	return
}
