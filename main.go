package main

import (
	"fmt"
	"github.com/AnthonyNixon/Debate-Bingo-Backend/auth"
	"github.com/AnthonyNixon/Debate-Bingo-Backend/handlers"
	"github.com/AnthonyNixon/Debate-Bingo-Backend/storage"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"time"
)

var PORT = ""

func init() {
	auth.Initialize()
	storage.Initialize()
	PORT = os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
}

func main() {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowMethods:     []string{"POST", "GET", "PUT", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		AllowAllOrigins:  true,
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.POST("/boards", handlers.NewBoard)
	r.POST("/toggleBox", handlers.ToggleBox)

	r.GET("/boards/:board_id", handlers.GetBoard)

	log.Printf("Running Debate Bingo API on :%s...", PORT)

	err := r.Run(fmt.Sprintf(":%s", PORT)) // listen and serve on 0.0.0.0:8080
	if err != nil {
		log.Fatal(err.Error())
	}
}