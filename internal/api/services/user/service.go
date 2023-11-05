package user

import (
	"Sesuai/internal/api/constracts"
	"Sesuai/internal/api/entities"
)

type Service struct {
	app  constracts.App
	repo constracts.UserRepository
}

func Init(a *constracts.App) (svc constracts.UserService) {
	r := initRepository(a.Datasources.WriterDB, a.Datasources.ReaderDB)

	svc = &Service{
		app:  *a,
		repo: r,
	}

	return
}

func (s Service) GetUserByEmail(email string) (user entities.User, err error) {
	user, err = s.repo.FindUserByEmail(email)

	return
}

func (s Service) RefreshToken(email, token string) (err error) {
	err = s.repo.RefreshToken(email, token)

	return
}

func (s Service) InsertUser(user entities.RequestRegister) (err error) {
	err = s.repo.InsertUser(user)

	return
}

func (s Service) IsExistPhoneNumber(phoneNumber string) bool {
	total, _ := s.repo.CountPhoneNumber(phoneNumber)

	if total > 0 {
		return true
	}

	return false
}

func (s Service) GetUserLoggedIn(userId, token string) (user entities.User, err error) {
	user, err = s.repo.FindUserLoggedIn(userId, token)

	return
}

func (s Service) GetProfileUser(userId string) (profileUser entities.User, err error) {
	profileUser, err = s.repo.FindProfileUser(userId)

	return
}

func (s Service) UpdateProfileUser(params entities.UpdateProfile) (err error) {
	err = s.repo.UpdateProfileUser(params)

	return
}

func (s Service) IsEmailAlreadyUsed(email, userId string) (isExist bool, err error) {
	count, err := s.repo.CountEmailAlreadyUsed(email, userId)

	return count > 0, err
}

func (s Service) IsPhoneNumberAlreadyUsed(phoneNumber, userId string) (isExist bool, err error) {
	count, err := s.repo.CountPhoneNumberAlreadyUsed(phoneNumber, userId)

	return count > 0, err
}
