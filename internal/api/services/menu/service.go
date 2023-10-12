package menu

import "Sesuai/internal/api/constracts"

type Service struct {
	app  constracts.App
	repo constracts.MenuRepository
}

func Init(a *constracts.App) (svc constracts.MenuService) {
	r := initRepository(a.Datasources.WriterDB, a.Datasources.ReaderDB)

	svc = &Service{
		app:  *a,
		repo: r,
	}

	return
}

func (s Service) GetMenuIdByName(name string) (id string, err error) {
	menu, err := s.repo.FindMenuIdByName(name)

	return menu.Id, err
}
