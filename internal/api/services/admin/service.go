package admin

import (
	"Sesuai/internal/api/constracts"
	"Sesuai/internal/api/entities"
)

type Service struct {
	app  constracts.App
	repo constracts.AdminRepository
}

func Init(a *constracts.App) (svc constracts.AdminService) {
	r := initRepository(a.Datasources.WriterDB, a.Datasources.ReaderDB)

	svc = &Service{
		app:  *a,
		repo: r,
	}

	return
}

func (s Service) GetAdmins() (admins []entities.AdminList, err error) {
	admins, err = s.repo.FindAdmins()

	if len(admins) == 0 {
		return []entities.AdminList{}, err
	}

	return
}

func (s Service) GetAdminById(adminId string) (admin entities.AdminList, err error) {
	admin, err = s.repo.FindAdminById(adminId)

	return
}

func (s Service) GetAdminByEmail(email string) (admin entities.Admin, err error) {
	admin, err = s.repo.FindAdminByEmail(email)

	return
}

func (s Service) RefreshToken(email, token string) (err error) {
	err = s.repo.RefreshToken(email, token)

	return
}

func (s Service) GetAdminLoggedIn(adminId, token string) (admin entities.Admin, err error) {
	admin, err = s.repo.FindAdminLoggedIn(adminId, token)

	return
}

func (s Service) IsEmailExist(email string) bool {
	total, _ := s.repo.CountEmail(email)

	if total > 0 {
		return true
	}

	return false
}

func (s Service) IsPhoneNumberExist(phoneNumber string) bool {
	total, _ := s.repo.CountPhoneNumber(phoneNumber)

	if total > 0 {
		return true
	}

	return false
}

func (s Service) InsertAdmin(params entities.RequestAdmin) (err error) {
	err = s.repo.InsertAdmin(params)

	return
}

func (s Service) UpdateAdmin(adminId string, params entities.RequestAdmin) (err error) {
	err = s.repo.UpdateAdmin(adminId, params)

	return
}

func (s Service) IsAdminExist(adminId string) bool {
	total, _ := s.repo.CountAdmin(adminId)

	if total > 0 {
		return true
	}

	return false
}

func (s Service) IsAdminWithTokenExist(adminId, token string) bool {
	total, _ := s.repo.CountAdminWithToken(adminId, token)

	if total > 0 {
		return true
	}

	return false
}

func (s Service) DeleteAdmin(adminId string) (err error) {
	err = s.repo.DeleteAdmin(adminId)

	return
}
