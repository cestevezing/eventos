package handlers

import (
	"net/http"

	"github.com/cestevezing/eventos/internal/app/models/entity"
	"github.com/cestevezing/eventos/internal/app/models/response"
	"github.com/cestevezing/eventos/internal/app/services"
	"github.com/gin-gonic/gin"
)

type TypeEventHandler struct {
	typeEventService *services.TypeEventService
}

func NewTypeEventHandler(typeEventService *services.TypeEventService) *TypeEventHandler {
	return &TypeEventHandler{
		typeEventService: typeEventService,
	}
}

// GetTypesEvent godoc
// @Summary Obtiene tipos de eventos
// @Description Obtiene una lista de tipos de eventos.
// @Tags Tipos Evento
// @Produce json
// @Success 200 {array} response.GenericResponse "Tipos de eventos"
// @Router /types-event [get]
func (eh *TypeEventHandler) GetTypesEvent(c *gin.Context) {
	typesEvent, err := eh.typeEventService.GetTypesEvent()
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.NewGenericResponse(http.StatusInternalServerError, err.Error(), nil))
		return
	}
	c.JSON(http.StatusOK, response.NewGenericResponse(http.StatusOK, "OK", typesEvent))
}

// @Summary Crea un nuevo tipo de evento
// @Description Crea un nuevo tipo de evento con los datos proporcionados.
// @Tags Tipos Evento
// @Accept json
// @Produce json
// @Param typeEvent body entity.TypeEvent true "Datos del tipo de evento a crear"
// @Success 201 {object} response.GenericResponse "Tipo de evento creado con éxito"
// @Router /types-event [post]
func (eh *TypeEventHandler) CreateTypeEvent(c *gin.Context) {
	var typeEvent entity.TypeEvent
	if err := c.BindJSON(&typeEvent); err != nil {
		c.JSON(http.StatusBadRequest, response.NewGenericResponse(http.StatusBadRequest, err.Error(), nil))
		return
	}

	createdTypeEvent, err := eh.typeEventService.CreateTypeEvent(&typeEvent)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.NewGenericResponse(http.StatusBadRequest, err.Error(), nil))
		return
	}
	c.JSON(http.StatusCreated, response.NewGenericResponse(http.StatusCreated, "OK", createdTypeEvent))

}

// @Summary Actualiza un tipo de evento existente
// @Description Actualiza un tipo de evento existente con los datos proporcionados.
// @Tags Tipos Evento
// @Accept json
// @Produce json
// @Param typeEvent body entity.TypeEvent true "Datos del tipo de evento a actualizar"
// @Success 200 {object} response.GenericResponse "Tipo de evento actualizado con éxito"
// @Router /types-event [put]
func (eh *TypeEventHandler) UpdateTypeEvent(c *gin.Context) {
	var typeEvent entity.TypeEvent
	if err := c.BindJSON(&typeEvent); err != nil {
		c.JSON(http.StatusBadRequest, response.NewGenericResponse(http.StatusBadRequest, err.Error(), nil))
		return
	}

	updatedTypeEvent, err := eh.typeEventService.UpdateTypeEvent(&typeEvent)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.NewGenericResponse(http.StatusBadRequest, err.Error(), nil))
		return
	}
	c.JSON(http.StatusOK, response.NewGenericResponse(http.StatusCreated, "OK", updatedTypeEvent))
}

// @Summary Elimina un tipo de evento por ID
// @Description Elimina un tipo de evento existente según su ID.
// @Tags Tipos Evento
// @Produce json
// @Param id path string true "ID del tipo de evento a eliminar"
// @Success 200 {object} response.GenericResponse "Tipo de evento eliminado con éxito"
// @Router /types-event/{id} [delete]
func (eh *TypeEventHandler) DeleteTypeEvent(c *gin.Context) {
	id := c.Param("id")
	if err := eh.typeEventService.DeleteTypeEvent(id); err != nil {
		c.JSON(http.StatusInternalServerError, response.NewGenericResponse(http.StatusBadRequest, err.Error(), nil))
		return
	}
	c.JSON(http.StatusCreated, response.NewGenericResponse(http.StatusCreated, "Se eliminó con éxito", nil))
}
