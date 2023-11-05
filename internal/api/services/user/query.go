package user

const (
	findUserByEmail = `
				SELECT 
    				u.id_user as user_id,
    				u.fullname as full_name,
    				u.email,
    				u.password,
    				u.phone_number,
    				u.date_birth,
    				u.birth_time,
    				bt.name as blood_type,
    				s.name as shio,
    				h.name as horoscope,
    				IFNULL(sex, "") as sex,
    				IFNULL(language, "") as language
				FROM 
				    user as u
				LEFT JOIN blood_type as bt
					ON u.id_blood_type = bt.id_blood_type
				LEFT JOIN shio as s
					ON u.id_shio = s.id_shio
				LEFT JOIN horoscope as h
					ON u.id_horoscope = h.id_horoscope
				WHERE 
				    u.email = ? 
				LIMIT 1
	`

	refreshToken = `UPDATE user SET token = ? WHERE email = ?`

	insertUser = `
		INSERT INTO user 
		    (email, password, fullname, phone_number, date_birth, birth_time, id_blood_type, id_shio, id_horoscope, sex, token)
		VALUES 
		    (:email, :password, :fullname, :phone_number, :date_birth, :birth_time, :id_blood_type, :id_shio, :id_horoscope, :sex, :token)
	`

	countPhoneNumber = `
		SELECT
			count(phone_number) as total
		FROM
		    user
		WHERE
		    phone_number = ?
	`

	findUserLoggedIn = `
		SELECT 
			u.id_user as user_id,
			u.fullname as full_name,
			u.email,
			u.password,
			u.phone_number,
			u.date_birth,
			bt.name as blood_type,
			s.name as shio,
			h.name as horoscope,
			IFNULL(sex, "") as sex,
			IFNULL(language, "") as language
		FROM 
			user as u
		LEFT JOIN blood_type as bt
			ON u.id_blood_type = bt.id_blood_type
		LEFT JOIN shio as s
			ON u.id_shio = s.id_shio
		LEFT JOIN horoscope as h
			ON u.id_horoscope = h.id_horoscope
		WHERE 
			u.id_user = ? and token = ? 
	`

	findProfileUser = `
		SELECT 
			u.id_user as user_id,
			u.fullname as full_name,
			u.email,
			u.password,
			u.phone_number,
			u.date_birth,
			u.birth_time,
			bt.name as blood_type,
			s.name as shio,
			h.name as horoscope,
			IFNULL(sex, "") as sex,
			IFNULL(language, "") as language
		FROM 
			user as u
		LEFT JOIN blood_type as bt
			ON u.id_blood_type = bt.id_blood_type
		LEFT JOIN shio as s
			ON u.id_shio = s.id_shio
		LEFT JOIN horoscope as h
			ON u.id_horoscope = h.id_horoscope
		WHERE 
			u.id_user = ? 
		LIMIT 1
	`

	updateProfileUser = `
		UPDATE 
		    user 
		SET 
		    fullname = :full_name, email = :email, phone_number = :phone_number, date_birth = :date_birth, 
		    birth_time = :birth_time, id_blood_type = :id_blood_type, id_shio = :id_shio, id_horoscope = :id_horoscope,
		    sex = :sex
		WHERE 
		    id_user = :id_user
	`
)
