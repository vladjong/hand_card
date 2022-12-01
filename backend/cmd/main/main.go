package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/vladjong/hand_card/config"
	"github.com/vladjong/hand_card/internal/app"
)

func main() {
	logrus.Info("init config")
	cfg := config.GetConfig()
	logrus.Info("env variables initializing")
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading env variables: %s", err.Error())
	}
	app, err := app.New(cfg)
	if err != nil {
		logrus.Fatal(err)
	}
	app.Run()
}
