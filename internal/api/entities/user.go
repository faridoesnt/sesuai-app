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
	Shio        int    `db:"shio" json:"shio"`
	Horoscope   string `db:"horoscope" json:"horoscope"`
	Sex         string `db:"sex" json:"gender"`
	Language    string `db:"language" json:"-"`
}

type UpdateProfile struct {
	UserId      string `db:"id_user" json:"-"`
	Email       string `db:"email" json:"email"`
	FullName    string `db:"full_name" json:"full_name"`
	PhoneNumber string `db:"phone_number" json:"phone_number"`
	DateBirth   string `db:"date_birth" json:"birth_date"`
	BirthTime   string `db:"birth_time" json:"birth_time"`
	BloodType   string `db:"id_blood_type" json:"id_blood_type"`
	Shio        string `db:"id_shio" json:"-"`
	Horoscope   string `db:"id_horoscope" json:"-"`
	Sex         string `db:"sex" json:"gender"`
}

type ChangePassword struct {
	CurrentPassword   string `db:"current_password" json:"current_password"`
	NewPassword       string `db:"new_password" json:"new_password"`
	RepeatNewPassword string `db:"-" json:"repeat_new_password"`
}
