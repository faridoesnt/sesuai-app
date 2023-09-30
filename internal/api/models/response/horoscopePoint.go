package response

type HoroscopePoint struct {
	HoroscopeId   string `json:"horoscope_id"`
	CategoryId    string `json:"category_id"`
	HoroscopeName string `json:"horoscope_name"`
	Point         string `json:"point"`
}
