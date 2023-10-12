package constracts

import "Sesuai/internal/api/entities"

type MenuRepository interface {
	FindMenuIdByName(name string) (menu entities.Menu, err error)
}

type MenuService interface {
	GetMenuIdByName(name string) (id string, err error)
}
