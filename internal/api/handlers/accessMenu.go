package handlers

import (
	"Sesuai/internal/api/helpers"
	"Sesuai/pkg/ahttp"
	"fmt"
	"github.com/kataras/iris/v12"
)

func GetAccessMenu(c iris.Context) {
	headers := helpers.GetHeaders(c)

	adminId := c.Params().GetString("adminId")
	if adminId == "" {
		HttpError(c, headers, fmt.Errorf("Admin Id Empty"), ahttp.ErrFailure("admin_id_empty"))
		return
	}

	accessMenu, err := app.Services.AccessMenu.GetAccessMenuByAdminId(adminId)
	if err != nil {
		HttpError(c, headers, fmt.Errorf(err.Error()), ahttp.ErrFailure(err.Error()))
		return
	}

	data := make(map[string]interface{})
	data["access_menu_list"] = []string{}

	if len(accessMenu) > 0 {
		data["access_menu_list"] = accessMenu
	}

	HttpSuccess(c, headers, data)
	return
}
