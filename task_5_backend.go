```go
package main

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

type OrderNotification struct {
	OrderID     string `json:"order_id"`
	Status      string `json:"status"`
	Description string `json:"description"`
}

type TelegramMessage struct {
	ChatID int64  `json:"chat_id"`
	Text   string `json:"text"`
}

func notifyTelegram(chatID int64, message string) error {
	telegramAPI := "https://api.telegram.org/bot<YOUR_BOT_TOKEN>/sendMessage"
	msg := TelegramMessage{
		ChatID: chatID,
		Text:   message,
	}
	data, _ := json.Marshal(msg)
	_, err := http.Post(telegramAPI, "application/json", bytes.NewBuffer(data))
	return err
}

func orderHandler(c *gin.Context) {
	var notification OrderNotification
	if err := c.ShouldBindJSON(&notification); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}
	message := "Order ID: " + notification.OrderID + "\nStatus: " + notification.Status + "\nDescription: " + notification.Description
	if err := notifyTelegram(<YOUR_CHAT_ID>, message); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "failed to send notification"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "notification sent"})
}

func setupRouter() *gin.Engine {
	router := gin.Default()
	router.POST("/order_notification", orderHandler)
	return router
}

func RunServer() {
	router := setupRouter()
	router.Run(":8080")
}
```