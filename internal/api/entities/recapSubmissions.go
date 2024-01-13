package entities

type RequestRecapSubmissions struct {
	Horoscope string `db:"horoscope" json:"zodiac"`
	Shio      string `db:"shio" json:"horoscope"`
	BloodType string `db:"blood_type" json:"blood_type"`
	Gender    string `db:"gender" json:"gender"`
}

type RecapUser struct {
	UserId    string `db:"id_user"`
	Name      string `db:"name"`
	BirthDate string `db:"birth_date"`
	Gender    string `db:"gender"`
	Horoscope string `db:"horoscope"`
	Shio      string `db:"shio"`
	BloodType string `db:"blood_type"`
}

type RecapSubmissions struct {
	TotalSubmissions       int `db:"total_submissions"`
	TotalUnlockSubmissions int `db:"total_unlock_submissions"`
}

type ResultRecapSubmissions struct {
	UserId                 string `json:"id_user"`
	Name                   string `json:"name"`
	BirthDate              string `json:"birth_date"`
	Gender                 string `json:"gender"`
	Horoscope              string `json:"horoscope"`
	Shio                   string `json:"shio"`
	BloodType              string `json:"blood_type"`
	TotalSubmissions       int    `json:"total_submissions"`
	TotalUnlockSubmissions int    `json:"total_unlock_submissions"`
}
