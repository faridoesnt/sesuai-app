package main

import (
	"Sesuai/pkg/asm"
	"log"
	"os"
)

func InitConfig() {
	env := "prod"
	secretName := "sesuai-prod"
	region := "ap-southeast-3"

	config, err := asm.GetSecret(secretName, region)
	if err != nil {
		log.Fatalf("Load Secret %s (%s) Failed, err: %s", secretName, env, err.Error())
		os.Exit(0)
	} else {
		log.Printf("Secret Loaded: %s (%s)", secretName, env)
	}

	app.Config = config
}
