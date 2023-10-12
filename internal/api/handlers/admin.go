package handlers

import (
	"Sesuai/internal/api/constants"
	"Sesuai/internal/api/entities"
	"Sesuai/internal/api/helpers"
	"Sesuai/pkg/ahttp"
	"fmt"
	"github.com/kataras/iris/v12"
	"golang.org/x/crypto/bcrypt"
)

func GetAdmins(c iris.Context) {
	headers := helpers.GetHeaders(c)

	admins, err := app.Services.Admin.GetAdmins()
	if err != nil {
		HttpError(c, headers, fmt.Errorf(err.Error()), ahttp.ErrFailure(err.Error()))
		return
	}

	for index, admin := range admins {
		accessMenu, err := app.Services.AccessMenu.GetAccessMenuByAdminId(admin.AdminId)
		if err != nil {
			HttpError(c, headers, fmt.Errorf(err.Error()), ahttp.ErrFailure(err.Error()))
			return
		}

		admins[index].AccessMenu = []string{}

		if len(accessMenu) > 0 {
			admins[index].AccessMenu = accessMenu
		}
	}

	data := make(map[string]interface{})
	data["admin_list"] = admins

	HttpSuccess(c, headers, data)
	return
}

func GetAdmin(c iris.Context) {
	headers := helpers.GetHeaders(c)

	adminId := c.Params().GetString("adminId")

	admin, err := app.Services.Admin.GetAdminById(adminId)
	if err != nil {
		HttpError(c, headers, fmt.Errorf(err.Error()), ahttp.ErrFailure(err.Error()))
		return
	}

	accessMenu, err := app.Services.AccessMenu.GetAccessMenuByAdminId(admin.AdminId)
	if err != nil {
		HttpError(c, headers, fmt.Errorf(err.Error()), ahttp.ErrFailure(err.Error()))
		return
	}

	admin.AccessMenu = []string{}

	if len(accessMenu) > 0 {
		admin.AccessMenu = accessMenu
	}

	HttpSuccess(c, headers, admin)
	return
}

func SaveAdmin(c iris.Context) {
	headers := helpers.GetHeaders(c)

	var accessId []string

	params := entities.RequestAdmin{}

	err := c.ReadJSON(&params)
	if err != nil {
		HttpError(c, headers, fmt.Errorf(err.Error()), ahttp.ErrFailure(err.Error()))
		return
	}

	if params.FullName == "" {
		HttpError(c, headers, fmt.Errorf("Full Name empty"), ahttp.ErrFailure("full_name_empty"))
		return
	}

	if params.Email == "" {
		HttpError(c, headers, fmt.Errorf("Email empty"), ahttp.ErrFailure("email_empty"))
		return
	}

	if params.Password == "" {
		HttpError(c, headers, fmt.Errorf("Password empty"), ahttp.ErrFailure("password_empty"))
		return
	}

	if params.PhoneNumber == "" {
		HttpError(c, headers, fmt.Errorf("Phone Number empty"), ahttp.ErrFailure("phone_number_empty"))
		return
	}

	phoneNumberExist := app.Services.Admin.IsPhoneNumberExist(params.PhoneNumber)
	if phoneNumberExist {
		HttpError(c, headers, fmt.Errorf("Phone number already registered"), ahttp.ErrFailure("phone_number_already_registered"))
		return
	}

	emailExist := app.Services.Admin.IsEmailExist(params.Email)
	if emailExist {
		HttpError(c, headers, fmt.Errorf("Email already registered"), ahttp.ErrFailure("email_already_registered"))
		return
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(params.Password), 16)
	if err != nil {
		HttpError(c, headers, err, ahttp.ErrFailure(err.Error()))
		return
	}

	params.Password = string(hashPassword)

	if len(params.Access) > 0 {
		for _, val := range params.Access {
			var id string

			switch val {
			case constants.EnumGenerateToken:
				id, err = app.Services.Menu.GetMenuIdByName(constants.GenerateToken)
				if err != nil {
					HttpError(c, headers, fmt.Errorf(err.Error()), ahttp.ErrBadRequest)
					return
				}
			case constants.EnumAdminList:
				id, err = app.Services.Menu.GetMenuIdByName(constants.AdminList)
				if err != nil {
					HttpError(c, headers, fmt.Errorf(err.Error()), ahttp.ErrBadRequest)
					return
				}
			case constants.EnumQuestionList:
				id, err = app.Services.Menu.GetMenuIdByName(constants.QuestionList)
				if err != nil {
					HttpError(c, headers, fmt.Errorf(err.Error()), ahttp.ErrBadRequest)
					return
				}
			case constants.EnumSubmition:
				id, err = app.Services.Menu.GetMenuIdByName(constants.Submition)
				if err != nil {
					HttpError(c, headers, fmt.Errorf(err.Error()), ahttp.ErrBadRequest)
					return
				}
			case constants.EnumElement:
				id, err = app.Services.Menu.GetMenuIdByName(constants.Element)
				if err != nil {
					HttpError(c, headers, fmt.Errorf(err.Error()), ahttp.ErrBadRequest)
					return
				}
			case constants.EnumPointAnswer:
				id, err = app.Services.Menu.GetMenuIdByName(constants.PointAnswer)
				if err != nil {
					HttpError(c, headers, fmt.Errorf(err.Error()), ahttp.ErrBadRequest)
					return
				}
			}

			accessId = append(accessId, id)
		}
	}

	params.Access = accessId

	err = app.Services.Admin.InsertAdmin(params)
	if err != nil {
		HttpError(c, headers, fmt.Errorf(err.Error()), ahttp.ErrFailure(err.Error()))
		return
	}

	HttpSuccess(c, headers, nil)
	return
}

func UpdateAdmin(c iris.Context) {
	headers := helpers.GetHeaders(c)

	var accessId []string

	adminId := c.Params().GetString("adminId")

	if adminId == "" {
		HttpError(c, headers, ahttp.Error{Message: "Admin Id Empty"}, ahttp.ErrFailure("admin_id_empty"))
		return
	}

	adminExist := app.Services.Admin.IsAdminExist(adminId)
	if !adminExist {
		HttpError(c, headers, ahttp.Error{Message: "Admin Not Found"}, ahttp.ErrFailure("admin_not_found"))
		return
	}

	params := entities.RequestAdmin{}

	err := c.ReadJSON(&params)
	if err != nil {
		HttpError(c, headers, fmt.Errorf(err.Error()), ahttp.ErrFailure(err.Error()))
		return
	}

	if params.FullName == "" {
		HttpError(c, headers, fmt.Errorf("Full Name empty"), ahttp.ErrFailure("full_name_empty"))
		return
	}

	if params.Email == "" {
		HttpError(c, headers, fmt.Errorf("Email empty"), ahttp.ErrFailure("email_empty"))
		return
	}

	if params.PhoneNumber == "" {
		HttpError(c, headers, fmt.Errorf("Phone Number empty"), ahttp.ErrFailure("phone_number_empty"))
		return
	}

	admin, err := app.Services.Admin.GetAdminById(adminId)
	if err != nil {
		HttpError(c, headers, fmt.Errorf(err.Error()), ahttp.ErrFailure(err.Error()))
		return
	}

	if admin.PhoneNumber != params.PhoneNumber {
		phoneNumberExist := app.Services.Admin.IsPhoneNumberExist(params.PhoneNumber)
		if phoneNumberExist {
			HttpError(c, headers, fmt.Errorf("Phone number already registered"), ahttp.ErrFailure("phone_number_already_registered"))
			return
		}
	}

	if admin.Email != params.Email {
		emailExist := app.Services.Admin.IsEmailExist(params.Email)
		if emailExist {
			HttpError(c, headers, fmt.Errorf("Email already registered"), ahttp.ErrFailure("email_already_registered"))
			return
		}
	}

	if len(params.Access) > 0 {
		for _, val := range params.Access {
			var id string

			switch val {
			case constants.EnumGenerateToken:
				id, err = app.Services.Menu.GetMenuIdByName(constants.GenerateToken)
				if err != nil {
					HttpError(c, headers, fmt.Errorf(err.Error()), ahttp.ErrBadRequest)
					return
				}
			case constants.EnumAdminList:
				id, err = app.Services.Menu.GetMenuIdByName(constants.AdminList)
				if err != nil {
					HttpError(c, headers, fmt.Errorf(err.Error()), ahttp.ErrBadRequest)
					return
				}
			case constants.EnumQuestionList:
				id, err = app.Services.Menu.GetMenuIdByName(constants.QuestionList)
				if err != nil {
					HttpError(c, headers, fmt.Errorf(err.Error()), ahttp.ErrBadRequest)
					return
				}
			case constants.EnumSubmition:
				id, err = app.Services.Menu.GetMenuIdByName(constants.Submition)
				if err != nil {
					HttpError(c, headers, fmt.Errorf(err.Error()), ahttp.ErrBadRequest)
					return
				}
			case constants.EnumElement:
				id, err = app.Services.Menu.GetMenuIdByName(constants.Element)
				if err != nil {
					HttpError(c, headers, fmt.Errorf(err.Error()), ahttp.ErrBadRequest)
					return
				}
			case constants.EnumPointAnswer:
				id, err = app.Services.Menu.GetMenuIdByName(constants.PointAnswer)
				if err != nil {
					HttpError(c, headers, fmt.Errorf(err.Error()), ahttp.ErrBadRequest)
					return
				}
			}

			accessId = append(accessId, id)
		}
	}

	params.Access = accessId

	err = app.Services.Admin.UpdateAdmin(adminId, params)
	if err != nil {
		HttpError(c, headers, fmt.Errorf(err.Error()), ahttp.ErrFailure(err.Error()))
		return
	}

	HttpSuccess(c, headers, nil)
	return
}

func DeleteAdmin(c iris.Context) {
	headers := helpers.GetHeaders(c)

	adminId := c.Params().GetString("adminId")

	if adminId == "" {
		HttpError(c, headers, ahttp.Error{Message: "Admin Id Empty"}, ahttp.ErrFailure("admin_id_empty"))
		return
	}

	adminExist := app.Services.Admin.IsAdminExist(adminId)
	if !adminExist {
		HttpError(c, headers, ahttp.Error{Message: "Admin Not Found"}, ahttp.ErrFailure("admin_not_found"))
		return
	}

	err := app.Services.Admin.DeleteAdmin(adminId)
	if err != nil {
		HttpError(c, headers, ahttp.Error{Message: "error delete admin"}, ahttp.ErrFailure(err.Error()))
		return
	}

	HttpSuccess(c, headers, nil)
	return
}
