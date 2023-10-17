package logging

const (
	insertMobileLogAdmin = `
		INSERT INTO mobile_log_admin 
		    (id_admin, endpoint, eventType, result, message, header, params, response, device, device_time)
		VALUES 
		    (:admin_id, :endpoint, :eventType, :result, :message, :header, :params, :response, :device, :device_time)
	`

	insertMobileLogUser = `
		INSERT INTO mobile_log_user
		    (id_user, endpoint, eventType, result, message, header, params, response, device, device_time)
		VALUES 
		    (:user_id, :endpoint, :eventType, :result, :message, :header, :params, :response, :device, :device_time)
	`
)
