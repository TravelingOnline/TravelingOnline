package main

import (
	"flag"
	"log"
	"os"
	"sync"

	"github.com/onlineTraveling/transport/api/handlers/grpc"
	"github.com/onlineTraveling/transport/app"
	"github.com/onlineTraveling/transport/config"
)

func main() {
	var configPath = flag.String("config", "config.yml", "configuration file path")
	flag.Parse()

	if v := os.Getenv("CONFIG_PATH"); len(v) > 0 {
		*configPath = v
	}
	c := config.MustReadConfig(*configPath)

	// err := logger.InitLogger(c)
	// if err != nil {
	// 	log.Fatal("can not initialize logger")
	// }
	log.Println("Starting the program")
	appContainer := app.NewMustApp(c)

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		grpc.Run(c, appContainer)
	}()
	wg.Wait()
}
