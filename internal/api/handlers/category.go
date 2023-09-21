package handlers

import (
	"Sesuai/internal/api/constants"
	"Sesuai/internal/api/entities"
	"Sesuai/internal/api/helpers"
	"Sesuai/pkg/ahttp"
	"errors"
	"github.com/kataras/iris/v12"
)

func GetCategory(c iris.Context) {
	headers := helpers.GetHeaders(c)

	categories := app.Services.Category.GetCategory()

	HttpSuccess(c, headers, categories)
	return
}

func GetCategoryDetail(c iris.Context) {
	headers := helpers.GetHeaders(c)

	categoryId := c.Params().GetString("categoryId")

	if categoryId == "" {
		HttpError(c, headers, ahttp.Error{Message: "Category Id Is Empty"}, ahttp.ErrFailure("Category Id Is Empty"))
		return
	}

	category := app.Services.Category.GetCategoryDetail(categoryId)
	if category.Id == "" {
		HttpError(c, headers, ahttp.Error{Message: "Category Not Found"}, ahttp.ErrNotFound("Category Not Found"))
		return
	}

	HttpSuccess(c, headers, category)
	return
}

func SaveCategory(c iris.Context) {
	headers := helpers.GetHeaders(c)

	adminId := c.Values().GetString(constants.AuthUserId)

	params := entities.RequestCategory{}

	err := c.ReadJSON(&params)
	if err != nil {
		HttpError(c, headers, err, ahttp.ErrInvalid(err.Error()))
		return
	}

	params.AdminId = adminId

	err = app.Services.Category.InsertCategory(params)
	if err != nil {
		HttpError(c, headers, errors.New("error insert category"), ahttp.ErrFailure(err.Error()))
		return
	}

	HttpSuccess(c, headers, nil)
	return
}

func UpdateCategory(c iris.Context) {
	headers := helpers.GetHeaders(c)

	adminId := c.Values().GetString(constants.AuthUserId)

	categoryId := c.Params().GetString("categoryId")

	params := entities.RequestCategory{}

	err := c.ReadJSON(&params)
	if err != nil {
		HttpError(c, headers, err, ahttp.ErrInvalid(err.Error()))
		return
	}

	params.AdminId = adminId

	if categoryId == "" {
		HttpError(c, headers, ahttp.Error{Message: "Category Id Is Empty"}, ahttp.ErrFailure("Category Id Is Empty"))
		return
	}

	if existCategory := app.Services.Category.IsExistCategory(categoryId); !existCategory {
		HttpError(c, headers, ahttp.Error{Message: "Category Not Found"}, ahttp.ErrFailure("Category Not Found"))
		return
	}

	err = app.Services.Category.UpdateCategory(categoryId, params)
	if err != nil {
		HttpError(c, headers, ahttp.Error{Message: "error update category"}, ahttp.ErrFailure(err.Error()))
		return
	}

	HttpSuccess(c, headers, nil)
	return
}

func DeleteCategory(c iris.Context) {
	headers := helpers.GetHeaders(c)

	categoryId := c.Params().GetString("categoryId")

	if categoryId == "" {
		HttpError(c, headers, ahttp.Error{Message: "Category Id Is Empty"}, ahttp.ErrFailure("Category Id Is Empty"))
		return
	}

	if existCategory := app.Services.Category.IsExistCategory(categoryId); !existCategory {
		HttpError(c, headers, ahttp.Error{Message: "Category Not Found"}, ahttp.ErrFailure("Category Not Found"))
		return
	}

	err := app.Services.Category.DeleteCategory(categoryId)
	if err != nil {
		HttpError(c, headers, ahttp.Error{Message: "error delete category"}, ahttp.ErrFailure(err.Error()))
		return
	}

	HttpSuccess(c, headers, nil)
	return
}
