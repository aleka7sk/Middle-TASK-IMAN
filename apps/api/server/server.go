package server

import (
	"context"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"os"
	"os/signal"
	"task/api/config"
	servicehttp "task/api/internal/service/delivery/http"
	grpcservice "task/pkg/grpc"
	"time"
)

type App struct {
	httpServer *http.Server
}

func NewApp() *App {
	return &App{}
}

func (a *App) Run(config *config.Config) error {
	router := gin.Default()
	connParser, err := grpc.Dial(":"+config.ParserPort, grpc.WithInsecure())
	if err != nil {
		return err
	}
	connCrud, err := grpc.Dial(":"+config.CrudPort, grpc.WithInsecure())
	if err != nil {
		return err
	}
	parser := grpcservice.NewCreatorClient(connParser)
	crud := grpcservice.NewEditorClient(connCrud)
	servicehttp.RegisterHTTPEndpoints(router, parser, crud)
	a.httpServer = &http.Server{
		Addr:           ":" + config.Port,
		Handler:        router,
		WriteTimeout:   20 * time.Second,
		ReadTimeout:    20 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := a.httpServer.ListenAndServe(); err != nil {
			log.Fatalf("Server listen and server error: %v\n", err)
		}
	}()
	log.Printf("Server run on port: %v\n", config.Port)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Interrupt)
	<-quit
	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()
	return a.httpServer.Shutdown(ctx)
}
