package entities

type HoroscopePoint struct {
	HoroscopeId   string `db:"horoscope_id"`
	ElementId     string `db:"category_id"`
	HoroscopeName string `db:"horoscope_name"`
	Point         string `db:"point"`
}

type RequestHoroscopePoint struct {
	HoroscopeId []string `json:"horoscope_id"`
	ElementId   string   `json:"element_id"`
	Point       []string `json:"point"`
}
