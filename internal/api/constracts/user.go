package constracts

import "Sesuai/internal/api/entities"

type UserRepository interface {
	FindUserByEmail(email string) (user entities.User, err error)
	RefreshToken(email, token string) (err error)
	InsertUser(user entities.RequestRegister) (err error)
	CountPhoneNumber(phoneNumber string) (total int64, err error)
	FindUserLoggedIn(userId, token string) (user entities.User, err error)
}

type UserService interface {
	GetUserByEmail(email string) (user entities.User, err error)
	RefreshToken(email, token string) (err error)
	InsertUser(user entities.RequestRegister) (err error)
	IsExistPhoneNumber(phoneNumber string) bool
	GetUserLoggedIn(userId, token string) (user entities.User, err error)
}
