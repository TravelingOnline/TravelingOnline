package main

import (
	"agency/api/handlers/http"
	"agency/app"
	"agency/config"
	"flag"
	"log"
	"os"
)

var configPath = flag.String("agency", "config.json", "service configuration file")

func main() {
	flag.Parse()

	if v := os.Getenv("AGENCY_PATH"); len(v) > 0 {
		*configPath = v
	}

	c := config.MustReadConfig(*configPath)

	appContainer := app.NewMustApp(c)

	log.Fatal(http.Run(appContainer, c.Server))
}
