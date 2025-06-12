```go
package main

import (
	"encoding/json"
	"net/http"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Order struct {
	ID     uint   `json:"id" gorm:"primaryKey"`
	Item   string `json:"item"`
	Amount int    `json:"amount"`
}

var db *gorm.DB

func init() {
	database, _ := gorm.Open(sqlite.Open("orders.db"), &gorm.Config{})
	database.AutoMigrate(&Order{})
	db = database
}

func getOrders(c *gin.Context) {
	var orders []Order
	db.Find(&orders)
	c.JSON(http.StatusOK, orders)
}

func createOrder(c *gin.Context) {
	var order Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Create(&order)
	c.JSON(http.StatusCreated, order)
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/orders", getOrders)
	r.POST("/orders", createOrder)
	return r
}

func run() {
	r := setupRouter()
	r.Run(":8080")
}
```