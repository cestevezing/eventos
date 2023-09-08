package main

import (
	"log"

	"github.com/cestevezing/eventos/api"
	"github.com/cestevezing/eventos/internal/logs"
	"github.com/joho/godotenv"
)

func main() {

	logs.InitLogger()

	err := godotenv.Load()
	if err != nil {
		logs.Logger.Print(err)
		log.Println("Error al cargar el archivo .env")
	}

	app := api.InitApi()

	app.Run(":3001")

}
