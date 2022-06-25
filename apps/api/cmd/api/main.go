package main

import (
	"apps/apps/api/config"
	"apps/apps/api/server"
	"log"
)

func main() {
	config, err := config.InitConfig()
	if err != nil {
		log.Fatalf("Init config error: %v\n", err)
	}
	app := server.NewApp()
	if err != nil {
		log.Fatalf("Create app error: %v\n", err)
	}
	if err := app.Run(config); err != nil {
		log.Fatalf("App run error: %v\n", err)
	}
}
