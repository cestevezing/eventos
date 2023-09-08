package routes

import (
	"github.com/cestevezing/eventos/internal/app/handlers"
	"github.com/gin-gonic/gin"
)

func EventRoutes(router *gin.Engine, eventHandler *handlers.EventHandler) {

	eventRoutes := router.Group("/events")

	eventRoutes.GET("/", eventHandler.GetEvents)
	eventRoutes.POST("/", eventHandler.CreateEvent)
	eventRoutes.PUT("/", eventHandler.UpdateEvent)
	eventRoutes.DELETE("/:id", eventHandler.DeleteEvent)
	eventRoutes.PUT("/update-status", eventHandler.UpdateSatusEvent)
	eventRoutes.GET("/get-management-required/:value", eventHandler.GetEventsByManagementRequired)
}
