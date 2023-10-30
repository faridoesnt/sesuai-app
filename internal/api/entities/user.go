package entities

type User struct {
	UserId      string `db:"user_id" json:"-"`
	FullName    string `db:"full_name" json:"full_name"`
	Email       string `db:"email" json:"email"`
	Password    string `db:"password" json:"-"`
	PhoneNumber string `db:"phone_number" json:"phone_number"`
	DateBirth   string `db:"date_birth" json:"birth_date"`
	BirthTime   string `db:"birth_time" json:"birth_time"`
	BloodType   string `db:"blood_type" json:"blood_type"`
	Shio        string `db:"shio" json:"shio"`
	Horoscope   string `db:"horoscope" json:"horoscope"`
	Sex         string `db:"sex" json:"gender"`
	Language    string `db:"language" json:"-"`
}
