package main

import (
	"fmt"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
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
	walletService := api.NewWalletService(storage)

	// create router dependency
	router := gin.Default()
	router.Use(app.LoggerToFile())
	router.Use(cors.Default())

	// setup cache
	cache := setupRedis()

	server := app.NewServer(router, cache, walletService)

	// start the server
	if err := server.Run(); err != nil {
		return err
	}

	return nil
}

func setupRedis() *redis.Client {
	REDIS_ADDR := os.Getenv("REDIS_ADDRESS")
	REDIS_PORT := os.Getenv("REDIS_PORT")

	ADDR := REDIS_ADDR + ":" + REDIS_PORT

	cache := redis.NewClient(&redis.Options{
		Addr:     ADDR,
		Password: "",
		DB:       0,
	})

	pong, err := cache.Ping().Result()
	fmt.Println(pong, err)
	return cache

}
