package horoscope

const (
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
