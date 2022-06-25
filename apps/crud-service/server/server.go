package server

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"log"
	"net"
	"task/config"
	"task/crud/internal"
	grpchandler "task/crud/internal/delivery/grpc"
	"task/crud/internal/repository"
	"task/crud/internal/usecase"
)

type App struct {
	usecase internal.UseCase
}

func NewApp(config *config.Config) (*App, error) {
	db, err := InitDB(config)
	if err != nil {
		return nil, err
	}
	repository := repository.NewRepository(db)
	return &App{
		usecase: usecase.NewService(repository),
	}, nil
}

func (a *App) Run(config *config.Config) error {
	server := grpc.NewServer()
	grpchandler.NewGRPCServer(server, a.usecase)

	l, err := net.Listen("tcp", ":8055")
	if err != nil {
		return err
	}
	log.Printf("Server run on port: %v\n", "8055")
	if err := server.Serve(l); err != nil {
		return err
	}
	return nil
}

func InitDB(config *config.Config) (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%v port=%v user=%v "+
		"password=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Password)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

//
//func Run() {
//	s := grpc.NewServer()
//	srv := &GRPCServer{}
//	grpcservice.RegisterEditorServer(s, srv)
//	l, err := net.Listen("tcp", ":8055")
//	if err != nil {
//		return
//	}
//	log.Printf("Server run on port: %v\n", "8055")
//	if err := s.Serve(l); err != nil {
//		return
//	}
//}
