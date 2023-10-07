package handlers

import (
	"Sesuai/internal/api/entities"
	"Sesuai/internal/api/helpers"
	"Sesuai/pkg/ahttp"
	"fmt"
	"github.com/kataras/iris/v12"
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
