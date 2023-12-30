package entities

type RequestCheckEmail struct {
	Email string `json:"email"`
}

type RequestLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Type     string `json:"type"`
}

type RequestRegister struct {
	Email       string  `db:"email" json:"email"`
	Password    string  `db:"password" json:"password"`
	FullName    string  `db:"fullname" json:"full_name"`
	PhoneNumber string  `db:"phone_number" json:"phone_number"`
	BirthDate   string  `db:"date_birth" json:"birth_date"`
	BirthTime   *string `db:"birth_time" json:"birth_time"`
	Gender      string  `db:"sex" json:"gender"`
	BloodType   string  `db:"id_blood_type" json:"id_blood_type"`
	Shio        string  `db:"id_shio"`
	Horoscope   string  `db:"id_horoscope"`
	Token       string  `db:"token"`
	Agreement   bool    `json:"user_agreement"`
}
