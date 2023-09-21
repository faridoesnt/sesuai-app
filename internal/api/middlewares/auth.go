package middlewares

import (
	"Sesuai/internal/api/constants"
	"Sesuai/internal/api/handlers"
	"Sesuai/internal/api/helpers"
	"Sesuai/pkg/ahttp"
	"github.com/kataras/iris/v12"
)

func Auth(c iris.Context) {
	headers := helpers.GetHeaders(c)

	token, status := handlers.ValidateToken(c, headers)
	if status != constants.SESSION_VALID {
		return
	}

	var id string

	user, err := app.Services.User.GetUserLoggedIn(headers.ID, token)
	if err == nil {
		id = user.UserId
	} else {
		if err.Error() == "sql: no rows in result set" {
			admin, err := app.Services.Admin.GetAdminLoggedIn(headers.ID, token)
			if err != nil {
				handlers.HttpError(c, headers, err, ahttp.ErrDenied(err.Error()))
			}

			id = admin.AdminId
		}
	}

	c.Values().Set(constants.AuthUserId, id)

	c.Next()
}
