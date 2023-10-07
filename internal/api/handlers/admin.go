package handlers

import (
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

		admins[index].AccessMenu = []entities.AccessMenu{}

		if len(accessMenu) > 0 {
			for _, val := range accessMenu {
				admins[index].AccessMenu = append(admins[index].AccessMenu, entities.AccessMenu{
					AccessId: val.AccessId,
					MenuId:   val.MenuId,
					MenuName: val.MenuName,
					Status:   val.Status,
				})
			}
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

	admin.AccessMenu = []entities.AccessMenu{}

	if len(accessMenu) > 0 {
		for _, val := range accessMenu {
			admin.AccessMenu = append(admin.AccessMenu, entities.AccessMenu{
				AccessId: val.AccessId,
				MenuId:   val.MenuId,
				MenuName: val.MenuName,
				Status:   val.Status,
			})
		}
	}

	HttpSuccess(c, headers, admin)
	return
}

func SaveAdmin(c iris.Context) {
	headers := helpers.GetHeaders(c)

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

	adminId := c.Params().GetString("adminId")

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

	err = app.Services.Admin.UpdateAdmin(adminId, params)
	if err != nil {
		HttpError(c, headers, fmt.Errorf(err.Error()), ahttp.ErrFailure(err.Error()))
		return
	}

	HttpSuccess(c, headers, nil)
	return
}
