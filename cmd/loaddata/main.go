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

	if err := database.DB.Where("1 = 1").Delete(&entity.TypeEvent{}).Error; err != nil {
		logs.Logger.Print(err)
		log.Println("No se ha podido limpiar la tabla TypeEvent.")
	}

	dataTypeEvent := []entity.TypeEvent{
		{
			ID:                 1,
			Name:               "Inicio de sesión",
			Description:        "Eventos de inicio de sesión realizados por los usuarios",
			ManagementRequired: false,
		},
		{
			ID:                 2,
			Name:               "Cambios de configuración",
			Description:        "Eventos de cambios de configuración de la aplicación",
			ManagementRequired: true,
		},
		{
			ID:                 3,
			Name:               "Detección de malware",
			Description:        "Eventos de detección de malware por el antivirus",
			ManagementRequired: true,
		},
		{
			ID:                 4,
			Name:               "Actividad en chat",
			Description:        "Eventos de actividad en chat de usuarios",
			ManagementRequired: false,
		},
	}

	database.DB.Create(dataTypeEvent)

	if err := database.DB.Where("1 = 1").Delete(&entity.Event{}).Error; err != nil {
		logs.Logger.Print(err)
		log.Println("No se ha podido limpiar la tabla TypeEvent.")
	}

	data := []entity.Event{
		{
			Name:        "Evento 1",
			TypeEventId: 1,
			Description: "Evento generado por el usuario",
			Status:      "Sin Revisar",
		},
		{
			Name:        "Evento 2",
			TypeEventId: 2,
			Description: "Se modificaron los colores de la página principal",
			Status:      "Sin Revisar",
		}, {
			Name:        "Evento 3",
			TypeEventId: 3,
			Description: "Se detectó una amenaza en el equipo del usuario cestevez",
			Status:      "Sin Revisar",
		}, {
			Name:        "Evento 4",
			TypeEventId: 4,
			Description: "El usuario cestevez se desconectó",
			Status:      "Revisar",
		}, {
			Name:        "Evento 5",
			TypeEventId: 4,
			Description: "El usuario cestevez se desconectó por inactividad",
			Status:      "Sin Revisar",
		}, {
			Name:        "Evento 6",
			TypeEventId: 2,
			Description: "Se cambio el orden del dashboard",
			Status:      "Sin Revisar",
		},
	}

	database.DB.Create(data)
}
