package alog

import (
	"github.com/sirupsen/logrus"
	"github.com/snowzach/rotatefilehook"
	"os"
)

var Logger *logrus.Logger

func Init() {
	Logger = logrus.New()

	rotateFileHook, err := rotatefilehook.NewRotateFileHook(rotatefilehook.RotateFileConfig{
		Filename:   os.Getenv(LogLocation),
		MaxSize:    1, // megabytes
		MaxBackups: 3,
		MaxAge:     90, //days
		Level:      logrus.DebugLevel,
		Formatter: &logrus.JSONFormatter{
			TimestampFormat: "2006-01-02T15:04:05.999999999Z07:00",
		},
	})

	if err != nil {
		Logger.Fatalf("Failed to initialize file rotate hook: %v", err)
	}

	Logger.SetFormatter(&logrus.TextFormatter{
		ForceColors:   true,
		FullTimestamp: true,
	})

	Logger.SetReportCaller(true)

	if os.Getenv(AppEnv) == "production" {
		Logger.AddHook(rotateFileHook)
	}
}
