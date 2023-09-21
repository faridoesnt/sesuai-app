package main

import (
	"Sesuai/pkg/alog"
	"github.com/joho/godotenv"
)

func InitConfig() {
	config, err := godotenv.Read()
	if err != nil {
		alog.Logger.Fatalf(err.Error())
	}

	app.Config = config
}
