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

	err := godotenv.Load()
	if err != nil {
		logrus.Fatal(err.Error())
	}

	config, err := internal.NewConfig()
	if err != nil {
		logrus.Fatal(err.Error())
	}

	err = internal.NewStorage(config).Init()
	if err != nil {
		logrus.Fatal(err.Error())
	}

	db, err := internal.NewSqliteDB()
	if err != nil {
		logrus.Fatal(err.Error())
	}

	repositories := repository.New(db)
	services := service.New(repositories, config.Get("storage.root"))
	handlers := handler.New(services)

	server := new(internal.Server)
	server.RunInGoroutine(config.Get("http.port"), handlers.InitRouter())

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
