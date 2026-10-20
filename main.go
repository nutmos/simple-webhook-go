package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r := setupRouter()
	log.Printf("listening on :%s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/v1")
	v1.POST("/webhook", handleWebhook)

	return r
}

func handleWebhook(c *gin.Context) {
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to read request body"})
		return
	}

	log.Printf("webhook received: content-type=%q headers=%v body=%s",
		c.ContentType(), c.Request.Header, body)

	c.JSON(http.StatusOK, gin.H{
		"status":      "received",
		"received_at": time.Now().UTC().Format(time.RFC3339),
		"bytes":       len(body),
	})
}
