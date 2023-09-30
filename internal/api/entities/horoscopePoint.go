package entities

type HoroscopePoint struct {
	HoroscopeId   string `db:"horoscope_id"`
	CategoryId    string `db:"category_id"`
	HoroscopeName string `db:"horoscope_name"`
	Point         string `db:"point"`
}
