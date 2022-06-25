package server

import (
  "apps/apps/config"
  "apps/apps/parser/internal"
  grpchandler "apps/apps/parser/internal/delivery/grpc"
  "apps/apps/parser/internal/repository"
  "apps/apps/parser/internal/usecase"
  "database/sql"
  "fmt"
  _ "github.com/lib/pq"
  "google.golang.org/grpc"
  "log"
  "net"
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

  l, err := net.Listen("tcp", ":8041")
  if err != nil {
    return err
  }
  log.Printf("Server run on port: %v\n", "8041")
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
