package app

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/yaska1706/quik-gaming-interview/pkg/api"
)

type Server struct {
	router        *gin.Engine
	cache         *redis.Client
	walletservice api.WalletService
}

func NewServer(router *gin.Engine, cache *redis.Client, walletservice api.WalletService) *Server {
	return &Server{
		router:        router,
		walletservice: walletservice,
		cache:         cache,
	}
}

func (s *Server) Run() error {
	// run function that initializes the routes
	r := s.Routes()

	LISTEN_ADDR := os.Getenv("LISTEN_ADDRESS")
	LISTEN_PORT := os.Getenv("LISTEN_PORT")

	if err := r.Run(LISTEN_ADDR + ":" + LISTEN_PORT); err != nil {
		log.Printf("Server - there was an error calling Run on router: %v", err)
		return err
	}

	return nil
}
