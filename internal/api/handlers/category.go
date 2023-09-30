package handlers

import (
	"Sesuai/internal/api/constants"
	"Sesuai/internal/api/entities"
	"Sesuai/internal/api/helpers"
	"Sesuai/internal/api/libs"
	"Sesuai/internal/api/models/response"
	"Sesuai/pkg/ahttp"
	"errors"
	"fmt"
	"github.com/kataras/iris/v12"
	"path/filepath"
	"time"
)

func GetCategory(c iris.Context) {
	headers := helpers.GetHeaders(c)

	categories := app.Services.Category.GetCategory()
	data := make(map[string]interface{})
	data["list_category"] = []response.Category{}

	if len(categories) > 0 {
		data["list_category"] = categories
	}

	HttpSuccess(c, headers, data)
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

	params := &entities.RequestCategory{}
	file := constants.IMAGE_MULTIPART
	parsed := false

	if c.GetContentTypeRequested() == "application/json" {
		file = constants.IMAGE_LOCALLY

		err := c.ReadJSON(params)
		if err != nil {
			HttpError(c, headers, err, ahttp.ErrInvalid(err.Error()))
			return
		}
	} else {
		parsed, params = parseSaveCategory(c.FormValues())

		if !parsed {
			HttpError(c, headers, fmt.Errorf("error parsed params save category"), ahttp.ErrFailure("error_while_parsing_params"))
			return
		}
	}

	params.AdminId = adminId

	if file == constants.IMAGE_MULTIPART {
		file, info, err := c.Request().FormFile("image")

		if err == nil {
			if file != nil {
				defer file.Close()
			}

			extension := filepath.Ext(info.Filename)

			timestamp := fmt.Sprintf("%d", time.Now().UnixNano()/int64(time.Millisecond))

			filename := timestamp + extension

			bucket := app.Config[constants.AwsS3Bucket]
			access_key := app.Config[constants.AwsS3Key]
			secret := app.Config[constants.AwsS3Secret]

			go libs.AWSMultipartUpload(bucket, access_key, secret, filename, file, info)

			params.FileName = filename
		}
	}

	err := app.Services.Category.InsertCategory(*params)
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

	params := &entities.RequestCategory{}
	file := constants.IMAGE_MULTIPART
	parsed := false

	if c.GetContentTypeRequested() == "application/json" {
		file = constants.IMAGE_LOCALLY

		err := c.ReadJSON(params)
		if err != nil {
			HttpError(c, headers, err, ahttp.ErrInvalid(err.Error()))
			return
		}
	} else {
		parsed, params = parseSaveCategory(c.FormValues())

		if !parsed {
			HttpError(c, headers, fmt.Errorf("error parsed params update category"), ahttp.ErrFailure("error_while_parsing_params"))
			return
		}
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

	if file == constants.IMAGE_MULTIPART {
		file, info, err := c.Request().FormFile("image")

		if err == nil {
			if file != nil {
				defer file.Close()
			}

			extension := filepath.Ext(info.Filename)

			timestamp := fmt.Sprintf("%d", time.Now().UnixNano()/int64(time.Millisecond))

			filename := timestamp + extension

			bucket := app.Config[constants.AwsS3Bucket]
			access_key := app.Config[constants.AwsS3Key]
			secret := app.Config[constants.AwsS3Secret]

			go libs.AWSMultipartUpload(bucket, access_key, secret, filename, file, info)

			params.FileName = filename
		}
	}

	err := app.Services.Category.UpdateCategory(categoryId, *params)
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

func parseSaveCategory(values map[string][]string) (parsed bool, params *entities.RequestCategory) {
	parsed = false
	tmp := &entities.RequestCategory{}
	params = tmp

	if len(values["name"]) > 0 {
		tmp.Name = values["name"][0]
	}

	if len(values["filename"]) > 0 {
		tmp.FileName = values["filename"][0]
	}

	if len(values["image"]) > 0 {
		tmp.Image = values["image"][0]
	}

	parsed = true
	params = tmp

	return
}
