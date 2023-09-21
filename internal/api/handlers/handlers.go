package handlers

import (
	"Sesuai/internal/api/constants"
	"Sesuai/internal/api/models"
	"Sesuai/pkg/ahttp"
	"Sesuai/pkg/alog"
	"encoding/json"
	"github.com/kataras/iris/v12"
	"github.com/sirupsen/logrus"
)

func HttpError(c iris.Context, headers *models.Headers, err error, httpError ahttp.Error) {
	c.StatusCode(httpError.Code)
	c.Header("Access-Control-Allow-Origin", "*")
	c.Request().Header.Set("Access-Control-Allow-Origin", "*")
	c.ResponseWriter().Header().Set("Access-Control-Allow-Origin", "*")

	alog.Logger.Errorf("%s %s error (%d): %s", c.Method(), c.Path(), httpError.Status, err)

	if httpError.Message == "DATETIME_VALIDATION" {
		//helpers.LogAction("DATETIME_VALIDATION", constants.ResponseInvalid, err.Error(), headers, err.Error())
	} else {
		_, errJson := json.Marshal(httpError) // _ : response
		if errJson != nil {
			logrus.Error("HTTP Result Error ("+c.Path()+"): ", err)
			//helpers.LogAction("HTTP_RESPONSE", "fail", "", headers, errJson.Error())
		}

		//helpers.LogAction("HTTP_RESPONSE", "fail", err.Error(), headers, string(response))
	}

	apiError := ahttp.CastError(httpError)
	if app.Config[constants.ServerEnv] == constants.EnvDevelopment {
		res := ahttp.ErrorResponse{
			Status:  apiError.Status,
			Message: apiError.Message,
			Data:    apiError.Data,
			Debug: &ahttp.ErrorDebug{
				Message: err.Error(),
			},
		}

		_ = c.JSON(res)
		c.StopExecution()
	}

	_ = c.JSON(apiError)
	c.StopExecution()
}

func HttpSuccess(c iris.Context, headers *models.Headers, data interface{}) {
	response := models.Response{}
	response.Data = data
	response.Status = "success"

	_, err := json.Marshal(response) // _ : res
	if err != nil {
		logrus.Error("HTTP Result Error ("+c.Path()+"): ", err)
		//helpers.LogAction("HTTP_RESPONSE", "fail", response.Message, headers, err.Error())
	}

	//helpers.LogAction("HTTP_RESPONSE", response.Status, response.Message, headers, string(res))

	c.Header("Access-Control-Allow-Origin", "*")
	c.Request().Header.Set("Access-Control-Allow-Origin", "*")
	c.ResponseWriter().Header().Set("Access-Control-Allow-Origin", "*")

	_ = c.JSON(response)
	c.Next()
}
