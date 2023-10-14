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

func GetElements(c iris.Context) {
	headers := helpers.GetHeaders(c)

	elements := app.Services.Element.GetElements()
	data := make(map[string]interface{})
	data["list_element"] = []response.Element{}

	if len(elements) > 0 {
		data["list_element"] = elements
	}

	HttpSuccess(c, headers, data)
	return
}

func GetElementDetail(c iris.Context) {
	headers := helpers.GetHeaders(c)

	elementId := c.Params().GetString("elementId")

	if elementId == "" {
		HttpError(c, headers, ahttp.Error{Message: "Element Id Is Empty"}, ahttp.ErrFailure("element_id_is_empty"))
		return
	}

	element := app.Services.Element.GetElementDetail(elementId)
	if element.Id == "" {
		HttpError(c, headers, ahttp.Error{Message: "Element Not Found"}, ahttp.ErrFailure("element_not_found"))
		return
	}

	HttpSuccess(c, headers, element)
	return
}

func SaveElement(c iris.Context) {
	headers := helpers.GetHeaders(c)

	adminId := c.Values().GetString(constants.AuthUserId)

	params := &entities.RequestElement{}
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
		parsed, params = parseSaveElement(c.FormValues())

		if !parsed {
			HttpError(c, headers, fmt.Errorf("error parsed params save element"), ahttp.ErrFailure("error_while_parsing_params"))
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

	err := app.Services.Element.InsertElement(*params)
	if err != nil {
		HttpError(c, headers, errors.New("error insert element"), ahttp.ErrFailure(err.Error()))
		return
	}

	HttpSuccess(c, headers, nil)
	return
}

func UpdateElement(c iris.Context) {
	headers := helpers.GetHeaders(c)

	adminId := c.Values().GetString(constants.AuthUserId)

	elementId := c.Params().GetString("elementId")

	params := &entities.RequestElement{}
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
		parsed, params = parseSaveElement(c.FormValues())

		if !parsed {
			HttpError(c, headers, fmt.Errorf("error parsed params update element"), ahttp.ErrFailure("error_while_parsing_params"))
			return
		}
	}

	params.AdminId = adminId

	if elementId == "" {
		HttpError(c, headers, ahttp.Error{Message: "Element Id Is Empty"}, ahttp.ErrFailure("element_id_is_empty"))
		return
	}

	if existElement := app.Services.Element.IsExistElement(elementId); !existElement {
		HttpError(c, headers, ahttp.Error{Message: "Element Not Found"}, ahttp.ErrFailure("element_not_found"))
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

	err := app.Services.Element.UpdateElement(elementId, *params)
	if err != nil {
		HttpError(c, headers, ahttp.Error{Message: "error update element"}, ahttp.ErrFailure(err.Error()))
		return
	}

	HttpSuccess(c, headers, nil)
	return
}

func DeleteElement(c iris.Context) {
	headers := helpers.GetHeaders(c)

	elementId := c.Params().GetString("elementId")

	if elementId == "" {
		HttpError(c, headers, ahttp.Error{Message: "Element Id Is Empty"}, ahttp.ErrFailure("element_id_is_empty"))
		return
	}

	if existElement := app.Services.Element.IsExistElement(elementId); !existElement {
		HttpError(c, headers, ahttp.Error{Message: "Element Not Found"}, ahttp.ErrFailure("element_not_found"))
		return
	}

	err := app.Services.Element.DeleteElement(elementId)
	if err != nil {
		HttpError(c, headers, ahttp.Error{Message: "error delete element"}, ahttp.ErrFailure(err.Error()))
		return
	}

	HttpSuccess(c, headers, nil)
	return
}

func parseSaveElement(values map[string][]string) (parsed bool, params *entities.RequestElement) {
	parsed = false
	tmp := &entities.RequestElement{}
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
