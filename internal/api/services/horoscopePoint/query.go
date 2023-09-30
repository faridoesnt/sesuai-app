package horoscopePoint

const (
	findHoroscopePoint = `
		SELECT
			h.id_horoscope as horoscope_id,
			c.id_category as category_id,
			h.name as horoscope_name,
			IFNULL(hp.point, "") as point
		FROM 
		    horoscope as h
		LEFT JOIN horoscope_point as hp
			ON hp.id_horoscope = h.id_horoscope
		LEFT JOIN category as c
			ON c.id_category = hp.id_category
		WHERE
		    c.id_category = ?
		ORDER BY h.name ASC
	`

	updateHoroscopePoint = `
		UPDATE horoscope_point SET point = :point WHERE id_horoscope = :id_horoscope AND id_category = :id_category 
	`
)
