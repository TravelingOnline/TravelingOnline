package main

import (
	"flag"
	"log"
	"os"
	"sync"

	"github.com/onlineTraveling/auth/api/grpc"
	"github.com/onlineTraveling/auth/api/http"
	"github.com/onlineTraveling/auth/api/message_broker"
	"github.com/onlineTraveling/auth/app"
	"github.com/onlineTraveling/auth/config"
	"github.com/onlineTraveling/auth/pkg/logger"
)

var configPath = flag.String("config", "config.json", "service configuration file")

func main() {
	flag.Parse()

	if v := os.Getenv("AUTH_CONFIG_PATH"); len(v) > 0 {
		*configPath = v
	}
	c := config.MustReadConfig(*configPath)
	err := logger.InitLogger(c)
	if err != nil {
		log.Fatal("can not initialize logger")
	}
	logger.Info("Starting Authentication Service..", nil)
	appContainer := *app.NewMustApp(c)
	err = http.Run(appContainer, c.Server)
	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		defer wg.Done()
		http.Run(appContainer, c.Server)
	}()
	go func() {
		defer wg.Done()
		grpc.Run(c, &appContainer)
	}()
	go func() {
		defer wg.Done()
		message_broker.Run(&appContainer)
	}()
	wg.Wait()

}
