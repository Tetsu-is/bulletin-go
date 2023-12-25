package main

import (
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Tweet struct {
	Id    string `json:"id" gorm:"primaryKey"`
	Title string `json:"title"`
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dsn := "host=" + host + " user=" + user + " password=" + password + " dbname=" + dbName + " port=" + port + " sslmode=disable TimeZone=Asia/Tokyo"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Tweet{})

	router := gin.Default()

	//cors setting from middleware
	router.Use(cors.New(cors.Config{

		AllowOrigins: []string{
			"http://localhost:3000",
		},

		AllowMethods: []string{
			"POST",
			"GET",
			"OPTIONS",
			"PUT",
			"DELETE",
		},

		AllowHeaders: []string{
			"Content-Type",
		},
	}))

	//rewite this with for loop
	for i := 1; i <= 10; i++ {
		db.Create(&Tweet{Id: string(i), Title: "Hello World" + string(i)})
	}

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	//receive json -> {offset: 0, limit: 10}
	router.GET("/tweets", func(c *gin.Context) {
		var tweets []Tweet
		db.Offset(0).Limit(10).Find(&tweets)
		c.JSON(200, tweets)
	})

	router.POST("/tweets", func(c *gin.Context) {
		var tweet Tweet
		c.BindJSON(&tweet)
		db.Create(&tweet)
		c.JSON(200, tweet)
	})

	router.Run()
}
