package autils

import "math/rand"

func RandomString(length int) string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	const (
		letterIdxBits = 6
		letterIdxMask = 1<<letterIdxBits - 1
		letterIdxMax  = 63 / letterIdxBits
	)

	src := rand.NewSource(GetDateTime(nil).UnixNano())
	b := make([]byte, length)

	for i, cache, remain := length-1, src.Int63(), letterIdxMax; i >= 0; {

		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}

		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}

		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}
