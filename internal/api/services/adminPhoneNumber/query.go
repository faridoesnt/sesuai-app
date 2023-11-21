package adminPhoneNumber

const (
	findAdminPhoneNumber = `
		SELECT
			IFNULL(id, "") as id,
			IFNULL(phone_number, "") as phone_number
		FROM
		    phone_number_request_token
		LIMIT 1
	`
)
