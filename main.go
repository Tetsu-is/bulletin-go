package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Todo struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

func main() {
	dsn := "host=localhost user=tetsuro password=postgres dbname=bulletin port=5432 sslmode=disable TimeZone=Asia/Tokyo"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Todo{})

	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.GET("/todos", func(c *gin.Context) {
		var todos []Todo
		db.Find(&todos)
		c.JSON(200, todos)
	})

	router.POST("/todos", func(c *gin.Context) {
		var todo Todo
		c.BindJSON(&todo)
		db.Create(&todo)
		c.JSON(200, todo)
	})

	router.Run()
}
