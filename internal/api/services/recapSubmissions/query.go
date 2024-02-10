package recapSubmissions

const (
	findRecapSubmissions = `
		SELECT
		    u.id_user,
			u.fullname as name,
			u.date_birth as birth_date,
			u.sex as gender,
			h.name as horoscope,
			s.name as shio,
			b.name as blood_type
		FROM
		    user as u
		LEFT JOIN horoscope as h 
			ON u.id_horoscope = h.id_horoscope
		LEFT JOIN shio as s
			ON u.id_shio = s.id_shio
		LEFT JOIN blood_type as b 
			ON u.id_blood_type = b.id_blood_type
	`

	whereHoroscope = ` u.id_horoscope = ?`
	whereShio      = ` u.id_shio = ?`
	whereBloodType = ` u.id_blood_type = ?`
	whereGender    = ` u.sex = ?`

	orderBy = ` ORDER BY u.fullname ASC`

	countRecapSubmissionsUser = `
		SELECT
			COUNT(s.id_user) as total_submissions,
			COUNT(ut.id_submission) as total_unlock_submissions
		FROM
			submission as s
		LEFT JOIN used_token as ut
			ON s.id_submission = ut.id_submission
		WHERE
			s.id_user = ?
	`

	findSummariesSubmissionByUserId = `
		SELECT
		    s.id_submission as submission_id,
			IFNULL(gt.token, "-") as token
		FROM
			submission s
		LEFT JOIN used_token ut ON s.id_submission = ut.id_submission
		LEFT JOIN generate_token gt ON ut.id_token = gt.id_token
		WHERE s.id_user = ?
		ORDER BY s.id_submission ASC
	`

	findSummariesPointSubmissionBySubmissionId = `
		SELECT
			c.name as element_name,
			ps.point
		FROM
		    point_submission ps
		LEFT JOIN category c ON ps.id_category = c.id_category
		WHERE ps.id_submission = ?
		ORDER BY c.name ASC
	`
)
