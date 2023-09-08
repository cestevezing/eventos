package api

import (
	"github.com/cestevezing/eventos/api/routes"
	"github.com/cestevezing/eventos/database"
	docs "github.com/cestevezing/eventos/docs"
	"github.com/cestevezing/eventos/internal/app/handlers"
	"github.com/cestevezing/eventos/internal/app/repository"
	"github.com/cestevezing/eventos/internal/app/services"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitApi() *gin.Engine {

	router := gin.Default()

	ConfigureRoutes(router)

	return router

}

func ConfigureRoutes(router *gin.Engine) {
	db, _ := database.Connect()

	docs.SwaggerInfo.Title = "API REST de Eventos"
	docs.SwaggerInfo.Description = "Documentación de proyecto"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:3001"
	docs.SwaggerInfo.BasePath = "/"

	typeEventRepository := repository.NewTypeEventRepository(db)
	typeEventService := services.NewTypeEventService(typeEventRepository)
	typeEventHandler := handlers.NewTypeEventHandler(typeEventService)

	routes.TypeEventRoutes(router, typeEventHandler)

	eventRepository := repository.NewEventRepository(db, typeEventRepository)
	eventService := services.NewEventService(eventRepository)
	eventHandler := handlers.NewEventHandler(eventService)

	routes.EventRoutes(router, eventHandler)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
