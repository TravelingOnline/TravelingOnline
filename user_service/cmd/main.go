package main

import (
	"fmt"
	"log"
	"user_service/config"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	fmt.Printf("Loaded configuration: %+v\n", cfg)
}

// func main() {
// 	flag.Parse()
// 	if v := os.Getenv("CONFIG_PATH"); len(v) > 0 {
// 		*configPath = v
// 	}

//		c := config.MustReadConfig(*configPath)
//		fmt.Println(c)
//	}
