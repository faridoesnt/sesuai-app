package main

import (
	"Sesuai/internal/api/constracts"
	"Sesuai/internal/api/services/accessMenu"
	"Sesuai/internal/api/services/admin"
	"Sesuai/internal/api/services/bloodType"
	"Sesuai/internal/api/services/bloodTypePoint"
	"Sesuai/internal/api/services/element"
	"Sesuai/internal/api/services/generateToken"
	"Sesuai/internal/api/services/horoscope"
	"Sesuai/internal/api/services/horoscopePoint"
	"Sesuai/internal/api/services/logging"
	"Sesuai/internal/api/services/menu"
	"Sesuai/internal/api/services/pointAnswer"
	"Sesuai/internal/api/services/question"
	"Sesuai/internal/api/services/shio"
	"Sesuai/internal/api/services/shioPoint"
	"Sesuai/internal/api/services/submission"
	"Sesuai/internal/api/services/user"
	"Sesuai/pkg/alog"
)

func InitServices() {
	app.Services = &constracts.Services{
		User:           user.Init(app),
		Admin:          admin.Init(app),
		BloodType:      bloodType.Init(app),
		BloodTypePoint: bloodTypePoint.Init(app),
		GenerateToken:  generateToken.Init(app),
		Element:        element.Init(app),
		Question:       question.Init(app),
		Submission:     submission.Init(app),
		Shio:           shio.Init(app),
		ShioPoint:      shioPoint.Init(app),
		Horoscope:      horoscope.Init(app),
		HoroscopePoint: horoscopePoint.Init(app),
		PointAnswer:    pointAnswer.Init(app),
		AccessMenu:     accessMenu.Init(app),
		Menu:           menu.Init(app),
		Logging:        logging.Init(app),
	}

	alog.Logger.Printf("Initializing Services: Pass")
}
