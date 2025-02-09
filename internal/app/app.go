package app

import (
	"github.com/pedroxer/BookingManagerSystem/internal/app/routes"
	config "github.com/pedroxer/BookingManagerSystem/internal/configs"
	"github.com/pedroxer/BookingManagerSystem/internal/storage"
	"github.com/sirupsen/logrus"
)

type App struct {
	Router routes.Router
}

func NewApp(log *logrus.Logger, storage *storage.Storage, config config.Api) *App {
	// init services

	// init router

}
