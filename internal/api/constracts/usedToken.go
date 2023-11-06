package constracts

type UsedTokenRepository interface {
	InsertUsedToken(tokenId, userId string) (err error)
	CountUserToken(token, userId string) (total int64, err error)
}

type UsedTokenService interface {
	InsertUsedToken(tokenId, userId string) (err error)
	IsUserToken(token, userId string) (isUserToken bool, err error)
}
