package main

import (
	"flag"
	"log"
	"os"

	"github.com/onlineTraveling/hotel/api/handlers/http"
	"github.com/onlineTraveling/hotel/app"
	"github.com/onlineTraveling/hotel/config"
	"github.com/onlineTraveling/hotel/pkg/logger"
)

var configPath = flag.String("config", "config.json", "service configuration file")

func main() {
	flag.Parse()

	if v := os.Getenv("HOTEL_CONFIG_PATH"); len(v) > 0 {
		*configPath = v
	}
	c := config.MustReadConfig(*configPath)
	err := logger.InitLogger(c)
	if err != nil {
		log.Fatal("can not initialize logger")
	}
	logger.Info("Starting the program", nil)
	appContainer := app.NewMustApp(c)
	err = http.Run(appContainer, c.Server)
	if err != nil {
		log.Fatal("can not start the programm")
	}

}
