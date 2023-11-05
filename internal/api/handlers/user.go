package handlers

import (
	"Sesuai/internal/api/constants"
	"Sesuai/internal/api/entities"
	"Sesuai/internal/api/helpers"
	"Sesuai/pkg/ahttp"
	"errors"
	"fmt"
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

func GetProfileUser(c iris.Context) {
	headers := helpers.GetHeaders(c)

	userId := c.Values().GetString(constants.AuthUserId)

	profileUser, err := app.Services.User.GetProfileUser(userId)
	if err != nil {
		HttpError(c, headers, fmt.Errorf(err.Error()), ahttp.ErrFailure(err.Error()))
		return
	}

	HttpSuccess(c, headers, profileUser)
	return
}

func UpdateProfileUser(c iris.Context) {
	headers := helpers.GetHeaders(c)

	userId := c.Values().GetString(constants.AuthUserId)

	params := entities.UpdateProfile{}

	err := c.ReadJSON(&params)
	if err != nil {
		HttpError(c, headers, fmt.Errorf(err.Error()), ahttp.ErrFailure(err.Error()))
		return
	}

	if params.FullName == "" {
		HttpError(c, headers, fmt.Errorf("Full Name Can't Empty"), ahttp.ErrFailure("full_name_can't_empty"))
		return
	}

	if params.Email == "" {
		HttpError(c, headers, fmt.Errorf("Email Can't Empty"), ahttp.ErrFailure("email_can't_empty"))
		return
	}

	if params.PhoneNumber == "" {
		HttpError(c, headers, fmt.Errorf("Phone Number Can't Empty"), ahttp.ErrFailure("phone_number_can't_empty"))
		return
	}

	if params.DateBirth == "" {
		HttpError(c, headers, fmt.Errorf("Birth Date Can't Empty"), ahttp.ErrFailure("birth_date_can't_empty"))
		return
	}

	if params.BirthTime == "" {
		HttpError(c, headers, fmt.Errorf("Birth Time Can't Empty"), ahttp.ErrFailure("birth_time_can't_empty"))
		return
	}

	if params.BloodType == "" {
		HttpError(c, headers, fmt.Errorf("Blood Type Can't Empty"), ahttp.ErrFailure("blood_type_can't_empty"))
		return
	}

	if params.Shio == "" {
		HttpError(c, headers, fmt.Errorf("Shio Can't Empty"), ahttp.ErrFailure("shio_can't_empty"))
		return
	}

	if params.Horoscope == "" {
		HttpError(c, headers, fmt.Errorf("Horoscope Can't Empty"), ahttp.ErrFailure("horoscope_can't_empty"))
		return
	}

	if params.Sex == "" {
		HttpError(c, headers, fmt.Errorf("Gender Can't Empty"), ahttp.ErrFailure("gender_can't_empty"))
		return
	}

	user, err := app.Services.User.GetUserByEmail(params.Email)
	if err == nil {
		if user.Email != params.Email {
			HttpError(c, headers, fmt.Errorf("Email Already Used"), ahttp.ErrFailure("email_already_used"))
			return
		}

		if user.PhoneNumber != params.PhoneNumber {
			HttpError(c, headers, fmt.Errorf("Phone Number Already Used"), ahttp.ErrFailure("phone_number_already_used"))
			return
		}
	}

	existHoroscope := app.Services.Horoscope.IsHoroscopeExist(params.Horoscope)
	if !existHoroscope {
		HttpError(c, headers, fmt.Errorf("Horoscope Not Found"), ahttp.ErrFailure("horoscope_not_found"))
		return
	}

	existShio := app.Services.Shio.IsShioExist(params.Shio)
	if !existShio {
		HttpError(c, headers, fmt.Errorf("Shio Not Found"), ahttp.ErrFailure("shio_not_found"))
		return
	}

	existBloodType := app.Services.BloodType.IsBloodTypeExist(params.BloodType)
	if !existBloodType {
		HttpError(c, headers, fmt.Errorf("Blood Type Not Found"), ahttp.ErrFailure("blood_type_not_found"))
		return
	}

	params.UserId = userId

	err = app.Services.User.UpdateProfileUser(params)
	if err != nil {
		HttpError(c, headers, fmt.Errorf(err.Error()), ahttp.ErrFailure(err.Error()))
		return
	}

	HttpSuccess(c, headers, nil)
	return
}
