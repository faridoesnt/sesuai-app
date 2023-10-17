package helpers

import (
	"Sesuai/internal/api/constants"
	"Sesuai/internal/api/models"
	"github.com/kataras/iris/v12"
	"strings"
)

func GetHeaders(c iris.Context) *models.Headers {

	headers := models.Headers{
		IP:            c.RemoteAddr(),
		Params:        "",
		Endpoint:      c.Path(),
		ID:            c.GetHeader("ID"),
		Authorization: c.GetHeader("Authorization"),
		Device:        c.GetHeader("Device-Model"),
		OS:            c.GetHeader("OS"),
		OSVersion:     c.GetHeader("OS-Version"),
		DateTime:      c.GetHeader("Device-Time"),
	}

	if headers.ID != "" {
		headers.User = "User"

		isAdmin := app.Services.Admin.IsAdminExist(headers.ID)
		if isAdmin {
			headers.User = "Admin"
		}
	}

	headers.InitParams(c)
	headers.InitHeader(c)

	return &headers
}

func GetAuthToken(authorization string) (status string, token string, message string) {
	status = constants.SESSION_DENIED

	if authorization != "" {
		if strings.Contains(authorization, "key ") {
			heads := strings.Split(authorization, " ")

			if len(heads) == 2 {
				status = constants.SESSION_VALID
				token = heads[1]
			} else {
				message = "invalid authorization value"
			}
		} else {
			message = "invalid authorization value"
		}
	} else {
		message = "authorization header not found"
	}

	return
}
