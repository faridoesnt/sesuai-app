package constracts

import "Sesuai/internal/api/entities"

type AdminRepository interface {
	FindAdmins() (admins []entities.AdminList, err error)
	FindAdminByEmail(email string) (admin entities.Admin, err error)
	RefreshToken(email, token string) (err error)
	FindAdminLoggedIn(adminId, token string) (admin entities.Admin, err error)
	CountEmail(email string) (total int64, err error)
}

type AdminService interface {
	GetAdmins() (admins []entities.AdminList, err error)
	GetAdminByEmail(email string) (admin entities.Admin, err error)
	RefreshToken(email, token string) (err error)
	GetAdminLoggedIn(adminId, token string) (admin entities.Admin, err error)
	IsEmailExist(email string) bool
}
