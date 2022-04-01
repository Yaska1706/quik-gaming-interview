package main

import (
	"fmt"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/yaska1706/quik-gaming-interview/pkg/api"
	"github.com/yaska1706/quik-gaming-interview/pkg/app"
	"github.com/yaska1706/quik-gaming-interview/pkg/repository"
)

func init() {
	if err := godotenv.Load(".env"); err != nil {
		panic(err)
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "this is the startup error: %s\n", err)
		os.Exit(1)
	}
}

func run() error {

	db, err := repository.SetupDB()
	if err != nil {
		return err
	}

	storage := repository.NewStorage(db)

	// create router dependency
	router := gin.Default()
	router.Use(app.LoggerToFile())
	router.Use(cors.Default())

	walletService := api.NewWalletService(storage)
	server := app.NewServer(router, walletService)

	// start the server
	if err := server.Run(); err != nil {
		return err
	}

	return nil
}
