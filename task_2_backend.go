```go
package main

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var telegramAPIURL = "https://api.telegram.org/bot"

func webhookHandler(c *gin.Context) {
	var update Update
	if err := c.BindJSON(&update); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	processUpdate(update)
	c.Status(http.StatusOK)
}

func processUpdate(update Update) {
	if update.Message != nil {
		sendMessage(update.Message.Chat.ID, "Hello, "+update.Message.From.FirstName+"!")
	}
}

func sendMessage(chatID int64, text string) {
	url := telegramAPIURL + os.Getenv("TELEGRAM_BOT_TOKEN") + "/sendMessage"
	body, _ := json.Marshal(map[string]interface{}{
		"chat_id": chatID,
		"text":    text,
	})
	http.Post(url, "application/json", bytes.NewBuffer(body))
}

type Update struct {
	Message *Message `json:"message"`
}

type Message struct {
	Chat *Chat   `json:"chat"`
	From *User   `json:"from"`
	Text string  `json:"text"`
}

type Chat struct {
	ID int64 `json:"id"`
}

type User struct {
	FirstName string `json:"first_name"`
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.POST("/webhook", webhookHandler)
	return r
}

func init() {
	r := setupRouter()
	r.Run(":8080")
}
```