package main

import (
	"apps/apps/config"
	"apps/apps/parser/server"
	"log"
)

func main() {
	config, err := config.InitConfig()
	if err != nil {
		log.Fatalf("Init config error: %v\n", err)
	}
	app, err := server.NewApp(config)
	if err != nil {
		log.Fatalf("Create app error: %v\n", err)
	}
	if err := app.Run(config); err != nil {
		log.Fatalf("App run error: %v\n", err)
	}
}
