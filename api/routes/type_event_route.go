package routes

import (
	"github.com/cestevezing/eventos/internal/app/handlers"
	"github.com/gin-gonic/gin"
)

func TypeEventRoutes(router *gin.Engine, typeEventHandler *handlers.TypeEventHandler) {
	eventRoutes := router.Group("/types-event")

	eventRoutes.GET("/", typeEventHandler.GetTypesEvent)
	eventRoutes.POST("/", typeEventHandler.CreateTypeEvent)
	eventRoutes.PUT("/", typeEventHandler.UpdateTypeEvent)
	eventRoutes.DELETE("/:id", typeEventHandler.DeleteTypeEvent)
}
