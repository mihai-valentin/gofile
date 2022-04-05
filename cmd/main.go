package main

import (
	"context"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"gofile/internal"
	"gofile/pkg/handler"
	"gofile/pkg/repository"
	"gofile/pkg/service"
	"os"
	"os/signal"
	"syscall"
)

/**
ToDo:
	tests
	readme
	aws
*/

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})

	if err := godotenv.Load(); err != nil {
		logrus.Fatal(err.Error())
	}

	config, err := internal.NewConfig()
	if err != nil {
		logrus.Fatal(err.Error())
	}

	if err := internal.NewStorage(config).Init(); err != nil {
		logrus.Fatal(err.Error())
	}

	db, err := internal.NewSqliteDB()
	if err != nil {
		logrus.Fatal(err.Error())
	}

	httpPort := config.Get("http.port")
	storageRoot := config.Get("storage.root")

	repositories := repository.New(db.DB)
	services := service.New(repositories, storageRoot)
	handlers := handler.New(services)

	server := new(internal.Server)
	server.RunInGoroutine(httpPort, handlers.InitRouter())

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	if err := server.Shutdown(context.Background()); err != nil {
		logrus.Fatal(err.Error())
	}

	if err := db.Close(); err != nil {
		logrus.Fatal(err.Error())
	}
}
