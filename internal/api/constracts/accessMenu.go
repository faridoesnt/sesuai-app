package constracts

import "Sesuai/internal/api/entities"

type AccessMenuRepository interface {
	FindAccessMenuByAdminId(adminId string) (accessMenus []entities.AccessMenu, err error)
}

type AccessMenuService interface {
	GetAccessMenuByAdminId(adminId string) (accessMenus []entities.AccessMenu, err error)
}
