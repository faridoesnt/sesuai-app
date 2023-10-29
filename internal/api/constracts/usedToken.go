package constracts

type UsedTokenRepository interface {
	InsertUsedToken(tokenId, userId string) (err error)
}

type UsedTokenService interface {
	InsertUsedToken(tokenId, userId string) (err error)
}
