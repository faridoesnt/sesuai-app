package middlewares

import (
	"Sesuai/internal/api/constants"
	"Sesuai/internal/api/handlers"
	"Sesuai/internal/api/helpers"
	"Sesuai/pkg/ahttp"
	"github.com/kataras/iris/v12"
)

func AuthAdmin(c iris.Context) {
	headers := helpers.GetHeaders(c)

	token, status := handlers.ValidateToken(c, headers)
	if status != constants.SESSION_VALID {
		return
	}

	var id string

	admin, err := app.Services.Admin.GetAdminLoggedIn(headers.ID, token)
	if err == nil {
		id = admin.AdminId
	} else {
		handlers.HttpError(c, headers, err, ahttp.ErrDenied(err.Error()))
	}

	c.Values().Set(constants.AuthUserId, id)

	c.Next()
}

func AuthUser(c iris.Context) {
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
		handlers.HttpError(c, headers, err, ahttp.ErrDenied(err.Error()))
	}

	c.Values().Set(constants.AuthUserId, id)

	c.Next()
}
