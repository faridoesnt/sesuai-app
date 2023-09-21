package main

import (
	"Sesuai/internal/api/constracts"
	"Sesuai/internal/api/services/admin"
	"Sesuai/internal/api/services/bloodType"
	"Sesuai/internal/api/services/category"
	"Sesuai/internal/api/services/generateToken"
	"Sesuai/internal/api/services/question"
	"Sesuai/internal/api/services/submission"
	"Sesuai/internal/api/services/user"
	"Sesuai/pkg/alog"
)

func InitServices() {
	app.Services = &constracts.Services{
		User:          user.Init(app),
		Admin:         admin.Init(app),
		BloodType:     bloodType.Init(app),
		GenerateToken: generateToken.Init(app),
		Category:      category.Init(app),
		Question:      question.Init(app),
		Submission:    submission.Init(app),
	}

	alog.Logger.Printf("Initializing Services: Pass")
}
