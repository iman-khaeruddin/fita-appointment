package main

import (
	"github.com/gin-gonic/gin"
	"github.com/iman-khaeruddin/fita-appointment/config"
	"github.com/iman-khaeruddin/fita-appointment/modules/appointment"
	_ "github.com/joho/godotenv/autoload"
	"log"
	"os"
)

func main() {
	router := gin.Default()

	db := config.GormMysql(os.Getenv("CONNECTION_STRING"))

	appointment := appointment.NewRequestHandler(db)
	appointment.Handle(router)

	err := router.Run()
	if err != nil {
		log.Println("main router.Run:", err)
		return
	}
}
