package constracts

import "Sesuai/internal/api/entities"

type LoggingRepository interface {
	InsertMobileLogAdmin(mobileLog entities.MobileLogAdmin) (err error)
	InsertMobileLogUser(mobileLog entities.MobileLogUser) (err error)
}

type LoggingService interface {
	InsertMobileLogAdmin(mobileLog entities.MobileLogAdmin) (err error)
	InsertMobileLogUser(mobileLog entities.MobileLogUser) (err error)
}
