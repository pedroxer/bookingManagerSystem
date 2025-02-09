package main

import (
	"encoding/json"
	env "github.com/caarlos0/env/v6"
	app "github.com/pedroxer/BookingManagerSystem/internal/app"
	config "github.com/pedroxer/BookingManagerSystem/internal/configs"
	"github.com/pedroxer/BookingManagerSystem/internal/database"
	"github.com/pedroxer/BookingManagerSystem/internal/storage"
	"github.com/sirupsen/logrus"
	"log"
	"os"
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

	db, err := database.ConnToPostgres(&cfg.Postgres)
	if err != nil {
		log.Fatal(err)
	}
	logger := logrus.New()
	logger.SetLevel(logrus.InfoLevel)

	storage := storage.New(db, logger)

	application := app.NewApp(logger, storage, cfg.Api)

}
