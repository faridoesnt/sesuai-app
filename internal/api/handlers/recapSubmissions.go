package handlers

import (
	"Sesuai/internal/api/constants"
	"Sesuai/internal/api/entities"
	"Sesuai/internal/api/helpers"
	"Sesuai/pkg/ahttp"
	"fmt"
	"github.com/kataras/iris/v12"
	"os"
	"time"
)

func RecapSubmissions(c iris.Context) {
	headers := helpers.GetHeaders(c)

	adminId := c.Values().GetString(constants.AuthUserId)

	admin, err := app.Services.Admin.GetAdminById(adminId)
	if err != nil {
		HttpError(c, headers, fmt.Errorf(err.Error()), ahttp.ErrFailure(err.Error()))
		return
	}

	params := entities.RequestRecapSubmissions{}

	err = c.ReadJSON(&params)
	if err != nil {
		HttpError(c, headers, fmt.Errorf(err.Error()), ahttp.ErrFailure(err.Error()))
		return
	}

	fetchData, err := app.Services.RecapSubmissions.GetRecapSubmissions(params)
	if err != nil {
		HttpError(c, headers, fmt.Errorf(err.Error()), ahttp.ErrFailure(err.Error()))
		return
	}

	if len(fetchData) == 0 {
		HttpError(c, headers, fmt.Errorf("tidak ada data submissions"), ahttp.ErrFailure("no_data"))
		return
	}

	excelFile, err := app.Services.RecapSubmissions.GenerateExcel(fetchData)
	if err != nil {
		HttpError(c, headers, fmt.Errorf(err.Error()), ahttp.ErrFailure(err.Error()))
		return
	}

	t, err := time.Parse(constants.FormatDateTime, headers.DateTime)
	if err != nil {
		HttpError(c, headers, fmt.Errorf(err.Error()), ahttp.ErrFailure(err.Error()))
		return
	}

	fileName := "Fetch Data " + t.Format("2 January 2006")

	currentDir, _ := os.Getwd()
	tmpFile := currentDir + "/internal/api/files/" + fileName + ".xlsx"
	if err := excelFile.SaveAs(tmpFile); err != nil {
		HttpError(c, headers, fmt.Errorf(err.Error()), ahttp.ErrFailure(err.Error()))
		return
	}

	sendEmail, err := helpers.SendEmail(tmpFile, admin.Email, fileName)
	if !sendEmail {
		HttpError(c, headers, fmt.Errorf(err.Error()), ahttp.ErrFailure(err.Error()))
		return
	}

	if err := os.Remove(tmpFile); err != nil {
		HttpError(c, headers, fmt.Errorf(err.Error()), ahttp.ErrFailure(err.Error()))
		return
	}

	HttpSuccess(c, headers, nil)
	return
}
