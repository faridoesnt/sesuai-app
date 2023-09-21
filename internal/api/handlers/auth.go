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

	//data := response.Auth{}
	data := make(map[string]interface{})
	hashPassword := ""

	user := entities.User{}
	admin := entities.Admin{}

	if params.Type != "admin" {
		user, _ = app.Services.User.GetUserByEmail(params.Email)

		data["id"] = user.UserId
		data["full_name"] = user.FullName
		data["email"] = user.Email
		data["phone_number"] = user.PhoneNumber
		data["date_birth"] = user.DateBirth
		data["birth_time"] = user.BirthTime
		data["sex"] = user.Sex
		data["blood_type"] = user.BloodType
		data["language"] = user.Language

		hashPassword = user.Password
	} else {
		admin, _ = app.Services.Admin.GetAdminByEmail(params.Email)

		data["id"] = admin.AdminId
		data["full_name"] = admin.FullName
		data["email"] = admin.Email

		hashPassword = admin.Password
	}

	// verify password
	if !verifyPassword(params.Password, hashPassword) {
		HttpError(c, headers, err, ahttp.ErrFailure("password_not_match"))
		return
	}

	// refresh token
	newToken := helpers.GenerateAccessToken(constants.TOKEN_AUTH_LENGTH)

	if params.Type != "admin" {
		err = app.Services.User.RefreshToken(user.Email, newToken)
		if err != nil {
			HttpError(c, headers, err, ahttp.ErrFailure("error while refresh token user"))
			return
		}
	} else {
		err = app.Services.Admin.RefreshToken(admin.Email, newToken)
		if err != nil {
			HttpError(c, headers, err, ahttp.ErrFailure("error while refresh token admin"))
			return
		}
	}

	data["token"] = newToken

	HttpSuccess(c, headers, data)
	return
}

func Register(c iris.Context) {
	headers := helpers.GetHeaders(c)

	params := entities.RequestRegister{}

	err := c.ReadJSON(&params)
	if err != nil {
		HttpError(c, headers, err, ahttp.ErrInvalid(err.Error()))
		return
	}

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
	}

	HttpSuccess(c, headers, data)
	return
}

func verifyPassword(password, hashPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
	return err == nil
}
