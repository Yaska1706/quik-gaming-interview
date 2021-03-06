package app

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
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
			c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprint(err)})
			return
		}

		response := map[string]string{
			"status": "success",
			"data":   "credit successful",
			"amount": newcreditrequest.Amount,
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
			c.JSON(http.StatusBadRequest, fmt.Sprint(err))
			return
		}

		response := map[string]string{
			"status": "success",
			"data":   "debit created",
			"amount": newdebitrequest.Amount,
		}

		c.JSON(http.StatusOK, response)

	}
}

func (s *Server) BalanceHandler() gin.HandlerFunc {

	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")

		walletID := c.Param("wallet_id")
		balance, err := s.walletservice.GetBalance(walletID)
		if err != nil {
			log.Printf("handler error: %v", err)
			c.JSON(http.StatusBadRequest, fmt.Sprint(err))
			return
		}

		response := map[string]string{
			"balance": balance,
		}

		response = s.VerifyCache(c, response, walletID)

		c.JSON(http.StatusOK, response)
	}

}

func (s *Server) VerifyCache(c *gin.Context, response map[string]string, walletID string) map[string]string {

	jsonresponse, _ := json.Marshal(response)
	if cacheERR := s.cache.Set(walletID, jsonresponse, 10*time.Second).Err(); cacheERR != nil {
		c.JSON(http.StatusBadRequest, fmt.Sprint(cacheERR))
		return nil
	}
	val, err := s.cache.Get(walletID).Result()
	if err != nil {
		c.Next()
		return nil
	}

	if err := json.Unmarshal([]byte(val), &response); err != nil {
		c.JSON(http.StatusBadRequest, fmt.Sprint(err))
		return nil
	}

	return response

}

func LoggerToFile() gin.HandlerFunc {
	LOG_FILE := os.Getenv("LOG_FILE")
	os.Create(LOG_FILE)
	src, err := os.OpenFile(LOG_FILE, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("err", err)
	}

	logger := logrus.New()
	logger.Out = src
	logger.SetLevel(logrus.DebugLevel)
	logger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	return func(c *gin.Context) {
		// start time
		startTime := time.Now()
		// Processing request
		c.Next()
		// End time
		endTime := time.Now()
		// execution time
		latencyTime := endTime.Sub(startTime)
		// Request mode
		reqMethod := c.Request.Method
		// Request routing
		reqUri := c.Request.RequestURI
		// Status code
		statusCode := c.Writer.Status()
		// Request IP
		clientIP := c.ClientIP()
		// Log format
		logger.Infof("| %3d | %13v | %15s | %s | %s |",
			statusCode,
			latencyTime,
			clientIP,
			reqMethod,
			reqUri,
		)
		logger.WithFields(logrus.Fields{
			"status_code":  statusCode,
			"latency_time": latencyTime,
			"client_ip":    clientIP,
			"req_method":   reqMethod,
			"req_uri":      reqUri,
		}).Info()
	}
}
