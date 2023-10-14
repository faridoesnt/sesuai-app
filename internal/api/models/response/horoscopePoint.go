package response

type HoroscopePoint struct {
	HoroscopeId   string `json:"horoscope_id"`
	ElementId     string `json:"element_id"`
	HoroscopeName string `json:"horoscope_name"`
	Point         string `json:"point"`
}
