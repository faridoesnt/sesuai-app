package entities

type User struct {
	UserId      string `db:"user_id"`
	FullName    string `db:"full_name"`
	Email       string `db:"email"`
	Password    string `db:"password"`
	PhoneNumber string `db:"phone_number"`
	DateBirth   string `db:"date_birth"`
	BirthTime   string `db:"birth_time"`
	BloodType   string `db:"blood_type"`
	Shio        string `db:"shio"`
	Horoscope   string `db:"horoscope"`
	Sex         string `db:"sex"`
	Language    string `db:"language"`
}
