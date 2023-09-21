package response

type Auth struct {
	Token       string `json:"token"`
	Id          string `json:"id"`
	FullName    string `json:"full_name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	DateBirth   string `json:"date_birth"`
	TimeBirth   string `json:"time_birth"`
	BloodType   string `json:"blood_type"`
	Shio        string `json:"shio"`
	Horoscope   string `json:"horoscope"`
	Sex         string `json:"sex"`
	Language    string `json:"language"`
}
