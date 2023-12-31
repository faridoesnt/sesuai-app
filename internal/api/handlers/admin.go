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

	adminId := c.Values().GetString(constants.AuthUserId)

	hasAccess, err := app.Services.AccessMenu.IsAdminHasAccessMenu(adminId, constants.AdminList)
	if err != nil {
		HttpError(c, headers, fmt.Errorf(err.Error()), ahttp.ErrFailure(err.Error()))
		return
	}

	if !hasAccess {
		HttpError(c, headers, fmt.Errorf("admin doesn't have access"), ahttp.ErrFailure("admin_doesn't_have_access"))
		return
	}

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

	adminId := c.Values().GetString(constants.AuthUserId)

	hasAccess, err := app.Services.AccessMenu.IsAdminHasAccessMenu(adminId, constants.AdminList)
	if err != nil {
		HttpError(c, headers, fmt.Errorf(err.Error()), ahttp.ErrFailure(err.Error()))
		return
	}

	if !hasAccess {
		HttpError(c, headers, fmt.Errorf("admin doesn't have access"), ahttp.ErrFailure("admin_doesn't_have_access"))
		return
	}

	adminId = c.Params().GetString("adminId")

	admin, err := app.Services.Admin.GetAdminById(adminId)
	if err != nil {
		HttpError(c, headers, fmt.Errorf("Admin Not Found"), ahttp.ErrFailure("admin_not_found"))
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

	adminId := c.Values().GetString(constants.AuthUserId)

	hasAccess, err := app.Services.AccessMenu.IsAdminHasAccessMenu(adminId, constants.AdminList)
	if err != nil {
		HttpError(c, headers, fmt.Errorf(err.Error()), ahttp.ErrFailure(err.Error()))
		return
	}

	if !hasAccess {
		HttpError(c, headers, fmt.Errorf("admin doesn't have access"), ahttp.ErrFailure("admin_doesn't_have_access"))
		return
	}

	var accessId []string

	params := entities.RequestAdmin{}

	err = c.ReadJSON(&params)
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

		params.Access = helpers.UniqueString(params.Access)

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
			case constants.EnumPoint:
				id, err = app.Services.Menu.GetMenuIdByName(constants.Point)
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

	adminId := c.Values().GetString(constants.AuthUserId)

	hasAccess, err := app.Services.AccessMenu.IsAdminHasAccessMenu(adminId, constants.AdminList)
	if err != nil {
		HttpError(c, headers, fmt.Errorf(err.Error()), ahttp.ErrFailure(err.Error()))
		return
	}

	if !hasAccess {
		HttpError(c, headers, fmt.Errorf("admin doesn't have access"), ahttp.ErrFailure("admin_doesn't_have_access"))
		return
	}

	var accessId []string

	adminId = c.Params().GetString("adminId")

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

	err = c.ReadJSON(&params)
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

		params.Access = helpers.UniqueString(params.Access)

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
			case constants.EnumPoint:
				id, err = app.Services.Menu.GetMenuIdByName(constants.Point)
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

	adminId := c.Values().GetString(constants.AuthUserId)

	hasAccess, err := app.Services.AccessMenu.IsAdminHasAccessMenu(adminId, constants.AdminList)
	if err != nil {
		HttpError(c, headers, fmt.Errorf(err.Error()), ahttp.ErrFailure(err.Error()))
		return
	}

	if !hasAccess {
		HttpError(c, headers, fmt.Errorf("admin doesn't have access"), ahttp.ErrFailure("admin_doesn't_have_access"))
		return
	}

	adminId = c.Params().GetString("adminId")

	if adminId == "" {
		HttpError(c, headers, ahttp.Error{Message: "Admin Id Empty"}, ahttp.ErrFailure("admin_id_empty"))
		return
	}

	adminExist := app.Services.Admin.IsAdminExist(adminId)
	if !adminExist {
		HttpError(c, headers, ahttp.Error{Message: "Admin Not Found"}, ahttp.ErrFailure("admin_not_found"))
		return
	}

	err = app.Services.Admin.DeleteAdmin(adminId)
	if err != nil {
		HttpError(c, headers, ahttp.Error{Message: "error delete admin"}, ahttp.ErrFailure(err.Error()))
		return
	}

	HttpSuccess(c, headers, nil)
	return
}

func ChangePasswordAdmin(c iris.Context) {
	headers := helpers.GetHeaders(c)

	adminId := c.Values().GetString(constants.AuthUserId)

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

	admin, err := app.Services.Admin.GetAdminById(adminId)
	if err != nil {
		HttpError(c, headers, fmt.Errorf(err.Error()), ahttp.ErrFailure(err.Error()))
		return
	}

	if !verifyPassword(params.CurrentPassword, admin.Password) {
		HttpError(c, headers, fmt.Errorf("current password not match"), ahttp.ErrFailure("current_password_not_match"))
		return
	}

	if params.NewPassword != params.RepeatNewPassword {
		HttpError(c, headers, fmt.Errorf("new password and repeat new password not match"), ahttp.ErrFailure("new_password_and_repeat_new_password_not_match"))
		return
	}

	if verifyPassword(params.NewPassword, admin.Password) {
		HttpError(c, headers, fmt.Errorf("current password and new password can't be same"), ahttp.ErrFailure("current_password_and_new_password_can't_be_same"))
		return
	}

	hashNewPassword, err := bcrypt.GenerateFromPassword([]byte(params.NewPassword), 16)
	if err != nil {
		HttpError(c, headers, fmt.Errorf(err.Error()), ahttp.ErrFailure(err.Error()))
		return
	}

	err = app.Services.Admin.ChangePassword(adminId, string(hashNewPassword))
	if err != nil {
		HttpError(c, headers, fmt.Errorf(err.Error()), ahttp.ErrFailure(err.Error()))
		return
	}

	HttpSuccess(c, headers, nil)
	return
}
