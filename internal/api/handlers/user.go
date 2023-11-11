package handlers

import (
	"Sesuai/internal/api/constants"
	"Sesuai/internal/api/entities"
	"Sesuai/internal/api/helpers"
	"Sesuai/pkg/ahttp"
	"errors"
	"fmt"
	"github.com/kataras/iris/v12"
	"golang.org/x/crypto/bcrypt"
	"time"
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

	if params.Sex == "" {
		HttpError(c, headers, fmt.Errorf("Gender Can't Empty"), ahttp.ErrFailure("gender_can't_empty"))
		return
	}

	emailAlreadyUsed, err := app.Services.User.IsEmailAlreadyUsed(params.Email, userId)
	if err != nil {
		HttpError(c, headers, fmt.Errorf(err.Error()), ahttp.ErrFailure(err.Error()))
		return
	}

	if emailAlreadyUsed {
		HttpError(c, headers, fmt.Errorf("Email Already Used"), ahttp.ErrFailure("email_already_used"))
		return
	}

	phoneNumberAlreadyUsed, err := app.Services.User.IsPhoneNumberAlreadyUsed(params.PhoneNumber, userId)
	if err != nil {
		HttpError(c, headers, fmt.Errorf(err.Error()), ahttp.ErrFailure(err.Error()))
		return
	}

	if phoneNumberAlreadyUsed {
		HttpError(c, headers, fmt.Errorf("Phone Number Already Used"), ahttp.ErrFailure("phone_number_already_used"))
		return
	}

	existBloodType := app.Services.BloodType.IsBloodTypeExist(params.BloodType)
	if !existBloodType {
		HttpError(c, headers, fmt.Errorf("Blood Type Not Found"), ahttp.ErrFailure("blood_type_not_found"))
		return
	}

	// get year on date birth for set shio
	t, err := time.Parse(constants.FormatDate, params.DateBirth)
	if err != nil {
		HttpError(c, headers, fmt.Errorf(err.Error()), ahttp.ErrFailure(err.Error()))
		return
	}

	year := t.Year()

	// get all shio
	shio, err := app.Services.Shio.GetShio()
	if err != nil {
		HttpError(c, headers, fmt.Errorf(err.Error()), ahttp.ErrFailure(err.Error()))
		return
	}

	shioYear := (year - 4) % 12

	// set shio
	params.Shio = shio[shioYear].Id

	// get horoscope from date birth
	horoscopeName := helpers.GetHoroscope(t)

	horoscope, err := app.Services.Horoscope.GetHoroscopeByName(horoscopeName)
	if err != nil {
		HttpError(c, headers, fmt.Errorf(err.Error()), ahttp.ErrFailure(err.Error()))
		return
	}

	// set horoscope
	params.Horoscope = horoscope.Id

	params.UserId = userId

	err = app.Services.User.UpdateProfileUser(params)
	if err != nil {
		HttpError(c, headers, fmt.Errorf(err.Error()), ahttp.ErrFailure(err.Error()))
		return
	}

	profileUser, err := app.Services.User.GetProfileUser(userId)
	if err != nil {
		HttpError(c, headers, fmt.Errorf(err.Error()), ahttp.ErrFailure(err.Error()))
		return
	}

	HttpSuccess(c, headers, profileUser)
	return
}

func ChangePassword(c iris.Context) {
	headers := helpers.GetHeaders(c)

	userId := c.Values().GetString(constants.AuthUserId)

	params := entities.ChangePassword{}

	err := c.ReadJSON(&params)
	if err != nil {
		HttpError(c, headers, fmt.Errorf(err.Error()), ahttp.ErrFailure(err.Error()))
		return
	}

	if params.CurrentPassword == "" {
		HttpError(c, headers, fmt.Errorf("current password can't empty"), ahttp.ErrFailure("current_password_can't_empty"))
		return
	}

	if params.NewPassword == "" {
		HttpError(c, headers, fmt.Errorf("new password can't empty"), ahttp.ErrFailure("new_password_can't_empty"))
		return
	}

	if params.RepeatNewPassword == "" {
		HttpError(c, headers, fmt.Errorf("repeat new password can't empty"), ahttp.ErrFailure("repeat_new_password_can't_empty"))
		return
	}

	user, err := app.Services.User.GetUserById(userId)
	if err != nil {
		HttpError(c, headers, fmt.Errorf(err.Error()), ahttp.ErrFailure(err.Error()))
		return
	}

	if !verifyPassword(params.CurrentPassword, user.Password) {
		HttpError(c, headers, fmt.Errorf("current password not match"), ahttp.ErrFailure("current_password_not_match"))
		return
	}

	if params.NewPassword != params.RepeatNewPassword {
		HttpError(c, headers, fmt.Errorf("new password and repeat new password not match"), ahttp.ErrFailure("new_password_and_repeat_new_password_not_match"))
		return
	}

	if verifyPassword(params.NewPassword, user.Password) {
		HttpError(c, headers, fmt.Errorf("current password and new password can't be same"), ahttp.ErrFailure("current_password_and_new_password_can't_be_same"))
		return
	}

	hashNewPassword, err := bcrypt.GenerateFromPassword([]byte(params.NewPassword), 16)
	if err != nil {
		HttpError(c, headers, fmt.Errorf(err.Error()), ahttp.ErrFailure(err.Error()))
		return
	}

	err = app.Services.User.ChangePassword(userId, string(hashNewPassword))
	if err != nil {
		HttpError(c, headers, fmt.Errorf(err.Error()), ahttp.ErrFailure(err.Error()))
		return
	}

	HttpSuccess(c, headers, nil)
	return
}
