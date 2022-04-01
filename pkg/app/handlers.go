package app

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yaska1706/quik-gaming-interview/pkg/api"
)

func (s *Server) ApiStatus() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")

		response := map[string]string{
			"status": "success",
			"data":   "Wallet API running smoothly",
		}

		c.JSON(http.StatusOK, response)
	}
}

func (s *Server) CreditHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")

		walletID := c.Param("wallet_id")

		var newcreditrequest api.CreditWalletRequest
		if err := c.ShouldBindJSON(&newcreditrequest); err != nil {
			log.Printf("handler error: %v", err)
			c.JSON(http.StatusBadRequest, nil)
			return
		}

		if err := s.walletservice.AddCredit(walletID, newcreditrequest); err != nil {
			log.Printf("handler error: %v", err)
			c.JSON(http.StatusBadRequest, nil)
			return
		}

		response := map[string]string{
			"status": "success",
			"data":   "debit created",
		}

		c.JSON(http.StatusOK, response)

	}
}

func (s *Server) DebitHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")

		walletID := c.Param("wallet_id")

		var newdebitrequest api.DebitWalletRequest
		if err := c.ShouldBindJSON(&newdebitrequest); err != nil {
			log.Printf("handler error: %v", err)
			c.JSON(http.StatusBadRequest, nil)
			return
		}

		if err := s.walletservice.AddDebit(walletID, newdebitrequest); err != nil {
			log.Printf("handler error: %v", err)
			c.JSON(http.StatusBadRequest, nil)
			return
		}

		response := map[string]string{
			"status": "success",
			"data":   "debit created",
		}

		c.JSON(http.StatusOK, response)

	}
}
