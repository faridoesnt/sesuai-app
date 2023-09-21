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

func (s Service) GetAdminByEmail(email string) (admin entities.Admin, err error) {
	admin, err = s.repo.FindAdminByEmail(email)

	return
}

func (s Service) RefreshToken(email, token string) (err error) {
	err = s.repo.RefreshToken(email, token)

	return
}

func (s Service) GetAdminLoggedIn(adminId, token string) (admin entities.Admin, err error) {
	admin, _ = s.repo.FindAdminLoggedIn(adminId, token)

	return
}
