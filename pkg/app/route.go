package app

import "github.com/gin-gonic/gin"

func (s *Server) Routes() *gin.Engine {
	router := s.router

	// group all routes under /v1/api
	v1 := router.Group("/v1/api/wallets")
	{
		v1.GET("/status", s.ApiStatus())
		v1.POST("/:wallet_id/credit", s.CreditHandler())
		v1.POST("/:wallet_id/debit", s.DebitHandler())
		v1.GET("/:wallet_id/balance", s.BalanceHandler())

	}

	return router
}
