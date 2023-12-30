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
	Shio        int    `json:"shio"`
	ShioSupport string `json:"shio_support"`
	Horoscope   string `json:"horoscope"`
	Sex         string `json:"sex"`
	Language    string `json:"language"`
	Type        string `json:"type"`
	TokenResult string `json:"token_result"`
}
