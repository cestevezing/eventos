package main

import (
	"log"

	"github.com/cestevezing/eventos/database"
	"github.com/cestevezing/eventos/internal/app/models/entity"
	"github.com/cestevezing/eventos/internal/logs"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		logs.Logger.Print(err)
		log.Println("Error al cargar el archivo .env")
	}
	database.Connect()
}

func main() {
	database.DB.AutoMigrate(&entity.TypeEvent{}, &entity.Event{})
}
