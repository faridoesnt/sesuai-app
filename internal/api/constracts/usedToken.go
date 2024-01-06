package constracts

import "Sesuai/internal/api/entities"

type UsedTokenRepository interface {
	InsertUsedToken(tokenId, submissionId string) (err error)
	CountSubmissionToken(token, submissionId string) (total int64, err error)
	FindUsedTokenByUserId(userId string) (tokenResults []entities.TokenResult, err error)
}

type UsedTokenService interface {
	InsertUsedToken(tokenId, submissionId string) (err error)
	IsSubmissionToken(token, submissionId string) (isSubmissionToken bool, err error)
	GetUsedTokenByUserId(userId string) (tokenResults []entities.TokenResult, err error)
}
