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

	if err := database.DB.Where("1 = 1").Delete(&entity.Event{}).Error; err != nil {
		logs.Logger.Print(err)
		log.Println("No se ha podido limpiar la tabla TypeEvent.")
	}

	if err := database.DB.Where("1 = 1").Delete(&entity.TypeEvent{}).Error; err != nil {
		logs.Logger.Print(err)
		log.Println("No se ha podido limpiar la tabla TypeEvent.")
	}

	typeEvent1 := entity.TypeEvent{
		Name:               "Inicio de sesión",
		Description:        "Eventos de inicio de sesión realizados por los usuarios",
		ManagementRequired: false,
	}

	database.DB.Create(&typeEvent1)
	typeEvent2 := entity.TypeEvent{
		Name:               "Cambios de configuración",
		Description:        "Eventos de cambios de configuración de la aplicación",
		ManagementRequired: true,
	}

	database.DB.Create(&typeEvent2)
	typeEvent3 := entity.TypeEvent{
		Name:               "Detección de malware",
		Description:        "Eventos de detección de malware por el antivirus",
		ManagementRequired: true,
	}

	database.DB.Create(&typeEvent3)
	typeEvent4 := entity.TypeEvent{
		Name:               "Actividad en chat",
		Description:        "Eventos de actividad en chat de usuarios",
		ManagementRequired: false,
	}
	database.DB.Create(&typeEvent4)

	data := []entity.Event{
		{
			Name:        "Evento 1",
			TypeEventId: typeEvent1.ID,
			Description: "Evento generado por el usuario",
			Status:      "Sin Revisar",
		},
		{
			Name:        "Evento 2",
			TypeEventId: typeEvent2.ID,
			Description: "Se modificaron los colores de la página principal",
			Status:      "Sin Revisar",
		}, {
			Name:        "Evento 3",
			TypeEventId: typeEvent3.ID,
			Description: "Se detectó una amenaza en el equipo del usuario cestevez",
			Status:      "Revisado",
		}, {
			Name:        "Evento 4",
			TypeEventId: typeEvent4.ID,
			Description: "El usuario cestevez se desconectó",
			Status:      "Sin Revisar",
		}, {
			Name:        "Evento 5",
			TypeEventId: typeEvent4.ID,
			Description: "El usuario cestevez se desconectó por inactividad",
			Status:      "Revisado",
		}, {
			Name:        "Evento 6",
			TypeEventId: typeEvent2.ID,
			Description: "Se cambio el orden del dashboard",
			Status:      "Sin Revisar",
		},
	}

	database.DB.Create(data)
	log.Println("Successful!")
}
