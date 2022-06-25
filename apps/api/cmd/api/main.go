package main

import (
	"log"
	"task/api/config"
	"task/api/server"
)

func main() {
	config, err := config.InitConfig()
	if err != nil {
		log.Fatalf("Init config error: %v\n", err)
	}
	app := server.NewApp()
	if err != nil {
		log.Fatalf("Create api error: %v\n", err)
	}
	if err := app.Run(config); err != nil {
		log.Fatalf("App run error: %v\n", err)
	}
}
