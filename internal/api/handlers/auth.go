package handlers

import (
	"Sesuai/internal/api/constants"
	"Sesuai/internal/api/entities"
	"Sesuai/internal/api/helpers"
	"Sesuai/internal/api/models"
	"Sesuai/internal/api/models/response"
	"Sesuai/pkg/ahttp"
	"Sesuai/pkg/autils"
	"errors"
	"fmt"
	"github.com/kataras/iris/v12"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func ValidateToken(c iris.Context, headers *models.Headers) (token, status string) {
	message := ""

	if status, token, message = helpers.GetAuthToken(headers.Authorization); status != constants.SESSION_VALID {
		HttpError(c, headers, errors.New(message), ahttp.ErrDenied(message))
		return
	}

	status = constants.SESSION_INVALID
	if headers.GetDateTime() != "" {
		if deviceTime, err := time.Parse(constants.FormatDateTime, headers.GetDateTime()); err == nil {
			status, message = autils.GetIntervalDeviceTime(deviceTime)
		}
	}

	if status == constants.SESSION_INVALID {
		HttpError(c, headers, errors.New(message), ahttp.ErrInvalid("DATETIME_VALIDATION"))
	}

	return
}

func CheckEmail(c iris.Context) {
	headers := helpers.GetHeaders(c)

	_, err := time.Parse(constants.FormatDateTime, headers.GetDateTime())
	if err != nil {
		HttpError(c, headers, err, ahttp.ErrInvalid(err.Error()))
		return
	}

	params := entities.RequestCheckEmail{}

	err = c.ReadJSON(&params)
	if err != nil {
		HttpError(c, headers, err, ahttp.ErrFailure(err.Error()))
		return
	}

	if params.Email == "" {
		HttpError(c, headers, errors.New("email can't empty"), ahttp.ErrFailure("email_cant_empty"))
		return
	}

	data := make(map[string]interface{})

	user, err := app.Services.User.GetUserByEmail(params.Email)
	if err == nil {
		data["email"] = user.Email
		data["type"] = "user"
	} else {
		if err.Error() == "sql: no rows in result set" {
			admin, err := app.Services.Admin.GetAdminByEmail(params.Email)
			if err != nil {
				HttpError(c, headers, err, ahttp.ErrFailure("email_not_found"))
				return
			}

			data["email"] = admin.Email
			data["type"] = "admin"
		}
	}

	HttpSuccess(c, headers, data)
	return
}

func Login(c iris.Context) {
	headers := helpers.GetHeaders(c)

	params := entities.RequestLogin{}

	err := c.ReadJSON(&params)
	if err != nil {
		HttpError(c, headers, err, ahttp.ErrInvalid(err.Error()))
		return
	}

	if params.Email == "" || params.Password == "" {
		HttpError(c, headers, err, ahttp.ErrFailure("email_or_password_cant_empty"))
		return
	}

	if params.Type == "" {
		HttpError(c, headers, err, ahttp.ErrFailure("type_cant_empty"))
		return
	}

	// generate new token
	newToken := helpers.GenerateAccessToken(constants.TOKEN_AUTH_LENGTH)

	if params.Type != "admin" {
		user, err := app.Services.User.GetUserByEmail(params.Email)
		if err != nil {
			HttpError(c, headers, fmt.Errorf(err.Error()), ahttp.ErrFailure(err.Error()))
			return
		}

		// verify password
		if !verifyPassword(params.Password, user.Password) {
			HttpError(c, headers, err, ahttp.ErrFailure("password_not_match"))
			return
		}

		err = app.Services.User.RefreshToken(user.Email, newToken)
		if err != nil {
			HttpError(c, headers, err, ahttp.ErrFailure("error while refresh token user"))
			return
		}

		data := &response.Auth{
			Token:       newToken,
			Id:          user.UserId,
			FullName:    user.FullName,
			Email:       user.Email,
			PhoneNumber: user.PhoneNumber,
			DateBirth:   user.DateBirth,
			TimeBirth:   user.BirthTime,
			BloodType:   user.BloodType,
			Shio:        user.Shio,
			Horoscope:   user.Horoscope,
			Sex:         user.Sex,
			Language:    user.Language,
			Type:        params.Type,
		}

		HttpSuccess(c, headers, data)
		return
	} else {
		admin, err := app.Services.Admin.GetAdminByEmail(params.Email)
		if err != nil {
			HttpError(c, headers, fmt.Errorf(err.Error()), ahttp.ErrFailure(err.Error()))
			return
		}

		// verify password
		if !verifyPassword(params.Password, admin.Password) {
			HttpError(c, headers, err, ahttp.ErrFailure("password_not_match"))
			return
		}

		// get access menu admin
		accessMenu, err := app.Services.AccessMenu.GetAccessMenuByAdminId(admin.AdminId)
		if err != nil {
			HttpError(c, headers, fmt.Errorf(err.Error()), ahttp.ErrFailure(err.Error()))
			return
		}

		err = app.Services.Admin.RefreshToken(admin.Email, newToken)
		if err != nil {
			HttpError(c, headers, err, ahttp.ErrFailure("error while refresh token admin"))
			return
		}

		data := make(map[string]interface{})
		data["token"] = newToken
		data["id"] = admin.AdminId
		data["full_name"] = admin.FullName
		data["email"] = admin.Email
		data["type"] = params.Type
		data["access_menu"] = []string{}

		if len(accessMenu) > 0 {
			data["access_menu"] = accessMenu
		}

		HttpSuccess(c, headers, data)
		return
	}
}

func Register(c iris.Context) {
	headers := helpers.GetHeaders(c)

	params := entities.RequestRegister{}

	err := c.ReadJSON(&params)
	if err != nil {
		HttpError(c, headers, err, ahttp.ErrInvalid(err.Error()))
		return
	}

	if params.Agreement {
		// get user for check email not used
		user, _ := app.Services.User.GetUserByEmail(params.Email)

		// check email not used
		if user.Email == params.Email {
			HttpError(c, headers, err, ahttp.ErrFailure("email_cant_same"))
			return
		}

		// check phone number not used
		existPhoneNumber := app.Services.User.IsExistPhoneNumber(params.PhoneNumber)
		if existPhoneNumber {
			HttpError(c, headers, err, ahttp.ErrFailure("phone_number_cant_same"))
			return
		}

		// generate token and set token
		token := helpers.GenerateAccessToken(constants.TOKEN_AUTH_LENGTH)
		params.Token = token

		// hash password
		hashPassword, err := bcrypt.GenerateFromPassword([]byte(params.Password), 16)
		if err != nil {
			HttpError(c, headers, err, ahttp.ErrFailure(err.Error()))
			return
		}

		params.Password = string(hashPassword)

		// get year on date birth
		t, err := time.Parse(constants.FormatDate, params.BirthDate)
		if err != nil {
			HttpError(c, headers, err, ahttp.ErrFailure(err.Error()))
			return
		}

		year := t.Year()

		// get all shio
		shio, err := app.Services.Shio.GetShio()
		if err != nil {
			HttpError(c, headers, err, ahttp.ErrFailure(err.Error()))
			return
		}

		shioYear := (year - 4) % 12

		// set shio
		params.Shio = shio[shioYear].Id

		// get horoscope from date birth
		horoscopeName := helpers.GetHoroscope(t)

		// set horoscope
		horoscope, err := app.Services.Horoscope.GetHoroscopeByName(horoscopeName)
		if err != nil {
			HttpError(c, headers, err, ahttp.ErrFailure(err.Error()))
			return
		}

		params.Horoscope = horoscope.Id

		// insert user
		err = app.Services.User.InsertUser(params)
		if err != nil {
			HttpError(c, headers, err, ahttp.ErrFailure("error while register"))
			return
		}

		// get user for response
		user, _ = app.Services.User.GetUserByEmail(params.Email)

		data := &response.Auth{
			Token:       token,
			Id:          user.UserId,
			FullName:    user.FullName,
			Email:       user.Email,
			PhoneNumber: user.PhoneNumber,
			DateBirth:   user.DateBirth,
			TimeBirth:   user.BirthTime,
			BloodType:   user.BloodType,
			Shio:        user.Shio,
			Horoscope:   user.Horoscope,
			Sex:         user.Sex,
			Language:    user.Language,
			Type:        "user",
		}

		HttpSuccess(c, headers, data)
		return
	} else {
		HttpError(c, headers, fmt.Errorf("user agreement must be true"), ahttp.ErrFailure("user_agreement_must_be_true"))
		return
	}

}

func verifyPassword(password, hashPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
	return err == nil
}
