package constracts

import "Sesuai/internal/api/entities"

type AdminRepository interface {
	FindAdmins() (admins []entities.AdminList, err error)
	FindAdminById(adminId string) (admin entities.AdminList, err error)
	FindAdminByEmail(email string) (admin entities.Admin, err error)
	RefreshToken(email, token string) (err error)
	FindAdminLoggedIn(adminId, token string) (admin entities.Admin, err error)
	CountEmail(email string) (total int64, err error)
	CountPhoneNumber(phoneNumber string) (total int64, err error)
	InsertAdmin(params entities.RequestAdmin) (err error)
	UpdateAdmin(adminId string, params entities.RequestAdmin) (err error)
}

type AdminService interface {
	GetAdmins() (admins []entities.AdminList, err error)
	GetAdminById(adminId string) (admin entities.AdminList, err error)
	GetAdminByEmail(email string) (admin entities.Admin, err error)
	RefreshToken(email, token string) (err error)
	GetAdminLoggedIn(adminId, token string) (admin entities.Admin, err error)
	IsEmailExist(email string) bool
	IsPhoneNumberExist(phoneNumber string) bool
	InsertAdmin(params entities.RequestAdmin) (err error)
	UpdateAdmin(adminId string, params entities.RequestAdmin) (err error)
}
