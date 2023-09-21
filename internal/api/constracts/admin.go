package constracts

import "Sesuai/internal/api/entities"

type AdminRepository interface {
	FindAdminByEmail(email string) (admin entities.Admin, err error)
	RefreshToken(email, token string) (err error)
	FindAdminLoggedIn(adminId, token string) (admin entities.Admin, err error)
}

type AdminService interface {
	GetAdminByEmail(email string) (admin entities.Admin, err error)
	RefreshToken(email, token string) (err error)
	GetAdminLoggedIn(adminId, token string) (admin entities.Admin, err error)
}
