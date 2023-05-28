package main

import (
	"github.com/gin-gonic/gin"
	"github.com/iman-khaeruddin/fita-appointment/config"
	_ "github.com/joho/godotenv/autoload"
	"log"
	"os"
)

func main() {
	router := gin.Default()

	config.GormMysql(os.Getenv("CONNECTION_STRING"))

	err := router.Run()
	if err != nil {
		log.Println("main router.Run:", err)
		return
	}
}
