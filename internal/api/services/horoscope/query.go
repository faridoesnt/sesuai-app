package horoscope

const (
	findHoroscopes = `
		SELECT
			id_horoscope,
			name
		FROM
		    horoscope
	`

	findHoroscopeByName = `
		SELECT
			id_horoscope,
			name
		FROM
		    horoscope
		WHERE
		    name = ?
	`
)
