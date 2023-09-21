package handlers

import (
	"Sesuai/internal/api/helpers"
	"Sesuai/pkg/ahttp"
	"errors"
	"github.com/kataras/iris/v12"
)

func GetUser(c iris.Context) {
	headers := helpers.GetHeaders(c)

	email := c.FormValue("email")

	if email == "" {
		HttpError(c, headers, errors.New("email can't empty"), ahttp.ErrFailure("email_cant_empty"))
		return
	}

	user, err := app.Services.User.GetUserByEmail(email)
	if err != nil {
		HttpError(c, headers, err, ahttp.ErrFailure("user_not_found"))
		return
	}

	data := make(map[string]interface{})
	data["email"] = user.Email
	data["full_name"] = user.FullName
	data["phone_number"] = user.PhoneNumber
	data["birth_date"] = user.DateBirth
	data["birth_time"] = user.BirthTime
	data["gender"] = user.Sex
	data["blood_type"] = user.BloodType
	data["shio"] = user.Shio
	data["horoscope"] = user.Horoscope

	HttpSuccess(c, headers, data)
	return
}
