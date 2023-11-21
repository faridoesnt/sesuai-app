package constracts

import "Sesuai/internal/api/entities"

type AdminPhoneNumberRepository interface {
	FindAdminPhoneNumber() (adminPhoneNumber entities.AdminPhoneNumber, err error)
}

type AdminPhoneNumberService interface {
	GetAdminPhoneNumber() (adminPhoneNumber entities.AdminPhoneNumber, err error)
}
