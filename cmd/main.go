package main

import (
	"encoding/json"
	env "github.com/caarlos0/env/v6"
	app "github.com/pedroxer/BookingManagerSystem/internal/app"
	config "github.com/pedroxer/BookingManagerSystem/internal/configs"
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	data, err := os.ReadFile("./configs/config.json")
	if err != nil {
		log.Fatal(err)
	}
	cfg := new(config.Config)
	if err := json.Unmarshal(data, cfg); err != nil {
		log.Fatal(err)
	}
	if err := env.Parse(cfg); err != nil {
		log.Fatal(err)
	}

	//db, err := database.ConnToPostgres(&cfg.Postgres)
	//if err != nil {
	//	log.Fatal(err)
	//}
	logger := logrus.New()
	logger.SetLevel(logrus.InfoLevel)

	//storage := storage.New(db, logger)

	application := app.NewApp(logger, *cfg)
	errChan := make(chan error, 1)
	sigChan := make(chan os.Signal, 1)
	logger.Info("starting the server...")
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	go func() {
		if err := application.Router.Start(); err != nil {
			errChan <- err
		}
	}()
	logger.Info("server started")
	var startErr error
	select {
	case sig := <-sigChan:
		logger.Infoln("got signal:", sig)
	case err := <-errChan:
		logger.Info(err)
		startErr = err
	}
	logger.Info("gracefully shutting down the server...")
	if err := application.Router.Shutdown(); err != nil {
		logger.Warn(err)
	}

	if startErr != nil {
		log.Fatal(startErr)
	}

	logger.Info("exited")

}
