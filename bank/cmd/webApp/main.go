package main

import (
	"flag"
	"log"
	"os"
	"sync"

	"github.com/onlineTraveling/bank/api/grpc"
	"github.com/onlineTraveling/bank/app"
	"github.com/onlineTraveling/bank/config"
	"github.com/onlineTraveling/bank/pkg/logger"
)

var configPath = flag.String("config", "config.json", "service configuration file")

func main() {
	flag.Parse()

	if v := os.Getenv("BANK_CONFIG_PATH"); len(v) > 0 {
		*configPath = v
	}
	c := config.MustReadConfig(*configPath)
	err := logger.InitLogger(c)
	if err != nil {
		log.Fatal("can not initialize logger")
	}
	logger.Info("Starting the program", nil)
	appContainer := app.NewMustApp(c)
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		grpc.Run(c, appContainer)
	}()
	wg.Wait()

}
