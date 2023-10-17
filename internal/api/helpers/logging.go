package helpers

import (
	"Sesuai/internal/api/constants"
	"Sesuai/internal/api/entities"
	"Sesuai/internal/api/models"
)

func LogActionAdmin(eventType, res, message string, headers *models.Headers, response string) {

	go func() {
		result := res
		switch res {
		case constants.ResponseFalse:
			result = constants.ResponseFail
		case constants.ResponseFail:
			result = constants.ResponseFail
		case constants.ResponseSuccess:
			result = constants.ResponseSuccess
		}

		if eventType != "" && (eventType == "HTTP_RESPONSE" || (result != constants.ResponseSuccess && result != "")) {

			dataActivityLog := entities.MobileLogAdmin{
				AdminId:    headers.ID,
				Endpoint:   headers.Endpoint,
				EventType:  eventType,
				Result:     result,
				Message:    message,
				Header:     headers.Header,
				Params:     headers.Params,
				Response:   response,
				Device:     headers.Device,
				DeviceTime: headers.GetDateTime(),
			}

			_ = app.Services.Logging.InsertMobileLogAdmin(dataActivityLog)
		}
	}()
}

func LogActionUser(eventType, res, message string, headers *models.Headers, response string) {

	go func() {
		result := res
		switch res {
		case constants.ResponseFalse:
			result = constants.ResponseFail
		case constants.ResponseFail:
			result = constants.ResponseFail
		case constants.ResponseSuccess:
			result = constants.ResponseSuccess
		}

		if eventType != "" && (eventType == "HTTP_RESPONSE" || (result != constants.ResponseSuccess && result != "")) {

			dataActivityLog := entities.MobileLogUser{
				UserId:     headers.ID,
				Endpoint:   headers.Endpoint,
				EventType:  eventType,
				Result:     result,
				Message:    message,
				Header:     headers.Header,
				Params:     headers.Params,
				Response:   response,
				Device:     headers.Device,
				DeviceTime: headers.GetDateTime(),
			}

			_ = app.Services.Logging.InsertMobileLogUser(dataActivityLog)
		}
	}()
}
