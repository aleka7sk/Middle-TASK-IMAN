package main

import (
	"log"
	"task/config"
	"task/crud/server"
)

func main() {
	config, err := config.InitConfig()
	if err != nil {
		log.Fatalf("Init config error: %v\n", err)
	}
	app, err := server.NewApp(config)
	if err != nil {
		log.Fatalf("Create api error: %v\n", err)
	}
	if err := app.Run(config); err != nil {
		log.Fatalf("App run error: %v\n", err)
	}
}
