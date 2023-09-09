package handlers

import (
	"net/http"
	"strconv"

	"github.com/cestevezing/eventos/internal/app/models/entity"
	"github.com/cestevezing/eventos/internal/app/models/request"
	"github.com/cestevezing/eventos/internal/app/models/response"
	"github.com/cestevezing/eventos/internal/app/services"
	"github.com/gin-gonic/gin"
)

type EventHandler struct {
	eventService *services.EventService
}

func NewEventHandler(eventService *services.EventService) *EventHandler {
	return &EventHandler{
		eventService: eventService,
	}
}

// GetEvents godoc
// @Summary Listar eventos
// @Description Lista todos los eventos almacenados
// @Tags Eventos
// @Produce json
// @Success 200 {array} response.GenericResponse
// @Router /events [get]
func (eh *EventHandler) GetEvents(c *gin.Context) {

	eventsResponse, err := eh.eventService.GetEvents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.NewGenericResponse(http.StatusInternalServerError, err.Error(), nil))
		return
	}
	c.JSON(http.StatusOK, response.NewGenericResponse(http.StatusOK, "OK", eventsResponse))
}

// @Summary Crea un nuevo evento
// @Description Crea un nuevo evento con los datos proporcionados.
// @Tags Eventos
// @Accept json
// @Produce json
// @Param event body request.EventCreate true "Datos del evento a crear"
// @Success 201 {object} response.GenericResponse "Evento creado con éxito"
// @Router /events [post]
func (eh *EventHandler) CreateEvent(c *gin.Context) {
	var event request.EventCreate
	if err := c.BindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, response.NewGenericResponse(http.StatusBadRequest, err.Error(), nil))
		return
	}
	if event.Status == "Revisado" || event.Status == "Sin Revisar" {
		eventEntity := entity.Event{
			Name:        event.Name,
			Description: event.Description,
			TypeEventId: event.TypeEventId,
			Date:        event.Date,
			Status:      event.Status,
		}
		err := eh.eventService.CreateEvent(&eventEntity)
		if err != nil {
			c.JSON(http.StatusInternalServerError, response.NewGenericResponse(http.StatusInternalServerError, err.Error(), nil))
			return
		}
		c.JSON(http.StatusCreated, response.NewGenericResponse(http.StatusOK, "OK", eventEntity))
	} else {
		c.JSON(http.StatusBadRequest, response.NewGenericResponse(http.StatusBadRequest, "El valor de State debe ser (Sin Revisar o Revisado)", nil))
		return
	}

}

// @Summary Actualiza un evento existente
// @Description Actualiza un evento existente con los datos proporcionados.
// @Tags Eventos
// @Accept json
// @Produce json
// @Param event body entity.Event true "Datos del evento a actualizar, el valor de Status debe ser (Sin Revisar o Revisado)"
// @Success 200 {object} response.GenericResponse "Evento actualizado con éxito"
// @Router /events [put]
func (eh *EventHandler) UpdateEvent(c *gin.Context) {
	var event entity.Event
	if err := c.BindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, response.NewGenericResponse(http.StatusBadRequest, err.Error(), nil))
		return
	}

	if event.Status == "Revisado" || event.Status == "Sin Revisar" {
		updatedEvent, err := eh.eventService.UpdateEvent(&event)
		if err != nil {
			c.JSON(http.StatusInternalServerError, response.NewGenericResponse(http.StatusInternalServerError, err.Error(), nil))
			return
		}
		c.JSON(http.StatusOK, response.NewGenericResponse(http.StatusOK, "OK", updatedEvent))
	} else {
		c.JSON(http.StatusBadRequest, response.NewGenericResponse(http.StatusBadRequest, "El valor de Status debe ser (Sin Revisar o Revisado)", nil))
		return
	}
}

// @Summary Elimina un evento por ID
// @Description Elimina un evento existente según su ID.
// @Tags Eventos
// @Produce json
// @Param id path string true "ID del evento a eliminar"
// @Success 200 {object} response.GenericResponse "Evento eliminado con éxito"
// @Router /events/{id} [delete]
func (eh *EventHandler) DeleteEvent(c *gin.Context) {
	id := c.Param("id")
	if err := eh.eventService.DeleteEvent(id); err != nil {
		c.JSON(http.StatusInternalServerError, response.NewGenericResponse(http.StatusInternalServerError, err.Error(), nil))
		return
	}
	c.JSON(http.StatusAccepted, response.NewGenericResponse(http.StatusAccepted, "Eliminado con éxito", nil))
}

// @Summary Actualiza el estado de un evento
// @Description Actualiza el estado de un evento existente con los datos proporcionados.
// @Tags Eventos
// @Accept json
// @Produce json
// @Param state body request.UpdateStatusRequest true "Datos del estado a actualizar,  el valor de State debe ser (Sin Revisar o Revisado)"
// @Success 202 {object} response.GenericResponse "Estado del evento actualizado con éxito"
// @Router /events/update-status [put]
func (eh *EventHandler) UpdateStatusEvent(c *gin.Context) {
	var state request.UpdateStatusRequest
	if err := c.BindJSON(&state); err != nil {
		c.JSON(http.StatusBadRequest, response.NewGenericResponse(http.StatusBadRequest, err.Error(), nil))
		return
	}

	if state.Status == "Revisado" || state.Status == "Sin Revisar" {
		event, err := eh.eventService.UpdateStatusEvent(state)
		if err != nil {
			c.JSON(http.StatusInternalServerError, response.NewGenericResponse(http.StatusInternalServerError, err.Error(), nil))
			return
		}
		c.JSON(http.StatusAccepted, response.NewGenericResponse(http.StatusAccepted, "Se actualizó con éxito", event))
	} else {
		c.JSON(http.StatusBadRequest, response.NewGenericResponse(http.StatusBadRequest, "El valor de status debe ser (Sin Revisar o Revisado)", nil))
		return
	}

}

// @Summary Obtiene eventos filtrados por requerimiento de gestión
// @Description Obtiene una lista de eventos filtrados por requerimiento de gestión (true o false), solo si es evento ya esta en estado Revisado.
// @Tags Eventos
// @Produce json
// @Param value path boolean true "Valor de requerimiento de gestión (true o false)"
// @Success 200 {array} response.GenericResponse "Eventos filtrados por requerimiento de gestión"
// @Router /events/get-management-required/{value} [get]
func (eh *EventHandler) GetEventsByManagementRequired(c *gin.Context) {
	valorBool, errc := strconv.ParseBool(c.Param("value"))

	if errc != nil {
		c.JSON(http.StatusInternalServerError, response.NewGenericResponse(http.StatusBadRequest, c.Param("value")+" No es un valor correcto, ingrese true o false", nil))
		return
	}

	events, err := eh.eventService.GetEventsByManagementRequired(valorBool)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.NewGenericResponse(http.StatusInternalServerError, err.Error(), nil))
		return
	}
	c.JSON(http.StatusOK, response.NewGenericResponse(http.StatusOK, "OK", events))
}
