package repository

import (
	"errors"

	"github.com/cestevezing/eventos/internal/app/models/entity"
	"github.com/cestevezing/eventos/internal/app/models/request"
	"gorm.io/gorm"
)

type EventRepository struct {
	db                  *gorm.DB
	typeEventRepository *TypeEventRepository
}

func NewEventRepository(db *gorm.DB, typeEventRepository *TypeEventRepository) *EventRepository {
	return &EventRepository{
		db:                  db,
		typeEventRepository: typeEventRepository,
	}
}

func (er *EventRepository) GetEvents() ([]*entity.Event, error) {
	var events []*entity.Event
	err := er.db.Find(&events).Error
	return events, err
}

func (er *EventRepository) CreateEvent(event *entity.Event) error {

	_, err := er.typeEventRepository.GetTypeEventById(event.TypeEventId)
	if err != nil {
		return errors.New("no se encontró el tipo evento")
	}
	return er.db.Create(event).Error
}

func (er *EventRepository) UpdateEvent(event *entity.Event) error {
	var resul entity.Event
	if err := er.db.Where("id = ?", event.ID).First(&resul).Error; err != nil {
		return errors.New("no se encuentró el evento")
	}
	return er.db.Save(event).Error
}

func (er *EventRepository) DeleteEvent(id string) error {
	return er.db.Delete(&entity.Event{}, id).Error
}

func (er *EventRepository) UpdateStatusEvent(status request.UpdateStatusRequest) (entity.Event, error) {
	var resul entity.Event
	if err := er.db.Where("id = ?", status.Id).First(&resul).Error; err != nil {
		return entity.Event{}, errors.New("no se encuentro el evento")
	}
	resul.Status = status.Status
	errUpdate := er.db.Save(resul).Error
	return resul, errUpdate
}

func (er *EventRepository) GetEventsByManagementRequired(managementRequired bool) ([]*entity.Event, error) {
	var resul []*entity.Event
	err := er.db.Joins("JOIN type_events ON events.type_event_id = type_events.id").
		Where("type_events.management_required = ? AND events.status = 'Revisado'", managementRequired).Find(&resul).Error
	return resul, err
}
