package constracts

import "Sesuai/internal/api/entities"

type AccessMenuRepository interface {
	FindAccessMenuByAdminId(adminId string) (accessMenus []entities.AccessMenu, err error)
	CountAdminAccessMenu(adminId, menu string) (count int64, err error)
}

type AccessMenuService interface {
	GetAccessMenuByAdminId(adminId string) (accessMenus []string, err error)
	IsAdminHasAccessMenu(adminId, menu string) (hasAccess bool, err error)
}
