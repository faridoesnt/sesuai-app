package main

import (
	"Sesuai/internal/api/constants"
	"Sesuai/internal/api/constracts"
	"Sesuai/internal/api/handlers"
	"Sesuai/internal/api/helpers"
	"Sesuai/internal/api/middlewares"
	"Sesuai/internal/api/routers"
	"Sesuai/pkg/alog"
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12"
	"net/http"
	"os"
	"time"
)

var app *constracts.App

func main() {
	os.Setenv("TZ", "Asia/Jakarta")

	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Access-Control-Allow-Origin", "Accept", "content-type", "X-Requested-With", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "Screen"},
		AllowCredentials: true,
	})

	irisApp := iris.New()
	irisApp.Use()
	irisApp.AllowMethods(iris.MethodOptions)

	app = &constracts.App{
		Iris: irisApp,
	}

	InitConfig()
	alog.Init()

	InitDatasource()
	InitServices()

	middlewares.Init(app)
	handlers.Init(app)
	routers.Init(app, crs)
	helpers.Init(app)

	srv := &http.Server{
		Addr:         ":" + os.Getenv(constants.ServerPort),
		ReadTimeout:  10 * time.Minute,
		WriteTimeout: 10 * time.Minute,
	}

	_ = irisApp.Run(iris.Server(srv), iris.WithOptimizations, iris.WithoutBodyConsumptionOnUnmarshal)
}
