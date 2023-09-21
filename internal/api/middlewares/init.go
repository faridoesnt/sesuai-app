package middlewares

import (
	. "Sesuai/internal/api/constracts"
)

var app *App

func Init(a *App) {
	app = a
}
