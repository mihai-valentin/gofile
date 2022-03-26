package main

import (
	"context"
	_ "github.com/mattn/go-sqlite3"
	"gofile"
	"gofile/pkg/handler"
	"gofile/pkg/repository"
	"gofile/pkg/service"
	"os"
	"os/signal"
	"syscall"
)

// ToDo:
// 		config
//		env
//		log
//		goroutines

func main() {
	db, err := repository.NewSqliteDB(repository.Config{
		File: "./files.db",
	})
	if err != nil {
		panic(err.Error())
	}

	repositories := repository.NewRepository(db)
	services := service.NewService(repositories)
	handlers := handler.NewHandler(services)

	server := new(gofile.Server)

	go func() {
		if err = server.Run("8080", handlers.InitRouter()); err != nil {
			panic(err.Error())
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	if err := server.Shutdown(context.Background()); err != nil {
		panic(err.Error())
	}

	if err := db.Close(); err != nil {
		panic(err.Error())
	}
}
