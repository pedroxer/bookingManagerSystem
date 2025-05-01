package app

import (
	config "github.com/pedroxer/BookingManagerSystem/internal/configs"
	"github.com/pedroxer/BookingManagerSystem/internal/routes"
	"github.com/pedroxer/BookingManagerSystem/pkg"
	"github.com/sirupsen/logrus"
)

type App struct {
	Router routes.Router
}

func NewApp(log *logrus.Logger, config config.Config) *App {
	resourceClient, err := pkg.CreateResourceClient(config.ResourceService)
	if err != nil {
		log.Fatalf("%s", err)
		return nil
	}
	bookingClient, err := pkg.CreateBookingClient(config.BookingService)
	if err != nil {
		log.Fatalf("%s", err)
		return nil
	}
	return &App{
		Router: *routes.NewRouter(log, &config, resourceClient, bookingClient),
	}
}
