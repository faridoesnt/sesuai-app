package entities

type AdminPhoneNumber struct {
	Id          string `db:"id"`
	PhoneNumber string `db:"phone_number"`
}
