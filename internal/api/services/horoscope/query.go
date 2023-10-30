package horoscope

const (
	findHoroscopes = `
		SELECT
			id_horoscope,
			name
		FROM
		    horoscope
		ORDER BY name ASC
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

	findHoroscopeUser = `
		SELECT
			h.id_horoscope,
			h.name
		FROM
		    horoscope as h
		LEFT JOIN user as u
			ON h.id_horoscope = u.id_horoscope
		WHERE
		    u.id_user = ?
	`
)
