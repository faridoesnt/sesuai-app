package constracts

import "Sesuai/internal/api/entities"

type UserRepository interface {
	FindUserByEmail(email string) (user entities.User, err error)
	RefreshToken(email, token string) (err error)
	InsertUser(user entities.RequestRegister) (err error)
	CountPhoneNumber(phoneNumber string) (total int64, err error)
	FindUserLoggedIn(userId, token string) (user entities.User, err error)
	FindProfileUser(userId string) (profileUser entities.User, err error)
	UpdateProfileUser(params entities.UpdateProfile) (err error)
	CountEmailAlreadyUsed(email, userId string) (total int64, err error)
	CountPhoneNumberAlreadyUsed(phoneNumber, userId string) (total int64, err error)
	FindUserById(userId string) (user entities.User, err error)
	ChangePassword(userId, newPassword string) (err error)
}

type UserService interface {
	GetUserByEmail(email string) (user entities.User, err error)
	RefreshToken(email, token string) (err error)
	InsertUser(user entities.RequestRegister) (err error)
	IsExistPhoneNumber(phoneNumber string) bool
	GetUserLoggedIn(userId, token string) (user entities.User, err error)
	GetProfileUser(userId string) (profileUser entities.User, err error)
	UpdateProfileUser(params entities.UpdateProfile) (err error)
	IsEmailAlreadyUsed(email, userId string) (isExist bool, err error)
	IsPhoneNumberAlreadyUsed(phoneNumber, userId string) (isExist bool, err error)
	GetUserById(userId string) (user entities.User, err error)
	ChangePassword(userId, newPassword string) (err error)
}
