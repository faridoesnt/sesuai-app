package constracts

type UsedTokenRepository interface {
	InsertUsedToken(tokenId, submissionId string) (err error)
	CountSubmissionToken(token, submissionId string) (total int64, err error)
	FindUsedTokenByUserId(userId string) (token string, err error)
}

type UsedTokenService interface {
	InsertUsedToken(tokenId, submissionId string) (err error)
	IsSubmissionToken(token, submissionId string) (isSubmissionToken bool, err error)
	GetUsedTokenByUserId(userId string) (token string, err error)
}
