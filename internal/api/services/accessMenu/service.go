package accessMenu

import (
	"Sesuai/internal/api/constracts"
	"Sesuai/internal/api/entities"
)

type Service struct {
	app  constracts.App
	repo constracts.AccessMenuRepository
}

func Init(a *constracts.App) (svc constracts.AccessMenuService) {
	r := initRepository(a.Datasources.WriterDB, a.Datasources.ReaderDB)

	svc = &Service{
		app:  *a,
		repo: r,
	}

	return
}

func (s Service) GetAccessMenuByAdminId(adminId string) (accessMenus []entities.AccessMenu, err error) {
	accessMenus, err = s.repo.FindAccessMenuByAdminId(adminId)

	return
}
