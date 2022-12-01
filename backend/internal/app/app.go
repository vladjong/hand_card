package app

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"github.com/vladjong/hand_card/config"
	v1 "github.com/vladjong/hand_card/internal/controller/http/v1"
	"github.com/vladjong/hand_card/pkg/postgres"
	"github.com/vladjong/hand_card/pkg/server"
)

type App struct {
	cfg            *config.Config
	postgresClient *sqlx.DB
}

func New(cfg *config.Config) (app App, err error) {
	postgresClient, err := postgres.NewClient(
		postgres.PostgresConfig{
			Host:     cfg.PostgresSQL.Host,
			Port:     cfg.PostgresSQL.Port,
			Username: cfg.PostgresSQL.Username,
			Password: os.Getenv("POSTGRES_PASSWORD"),
			DBName:   cfg.PostgresSQL.DBName,
			SSLMode:  cfg.PostgresSQL.SSLMode,
		})
	if err != nil {
		return app, err
	}
	return App{
		cfg:            cfg,
		postgresClient: postgresClient,
	}, nil
}

func (a *App) Run() error {
	// logrus.Info("initializing")
	a.startHTTP()
	return nil
}

func (a *App) startHTTP() {
	logrus.Info("HTTP Server initializing")
	server := new(server.Server)
	handlers := v1.New()
	go func() {
		if err := server.Run(a.cfg.Listen.Port, handlers.NewRouter()); err != nil {
			logrus.Fatalf("error: occured while running HTTP Server: %s", err.Error())
		}
	}()
	logrus.Info("HTTP Server start")
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	logrus.Info("HTTP Server Shutdown")
	if err := server.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error: occured on server shutdown: %s", err.Error())
	}
	if err := a.postgresClient.Close(); err != nil {
		logrus.Errorf("error: occured on db connection close: %s", err.Error())
	}
}
