package services

import (
	"github.com/cestevezing/eventos/internal/app/models/entity"
	"github.com/cestevezing/eventos/internal/app/models/request"
	"github.com/cestevezing/eventos/internal/app/repository"
)

type EventService struct {
	eventRepository *repository.EventRepository
}

func NewEventService(eventRepository *repository.EventRepository) *EventService {
	return &EventService{
		eventRepository: eventRepository,
	}
}

func (s *EventService) GetEvents() ([]*entity.Event, error) {
	return s.eventRepository.GetEvents()
}

func (s *EventService) CreateEvent(event *entity.Event) (*entity.Event, error) {
	if err := s.eventRepository.CreateEvent(event); err != nil {
		return nil, err
	}
	return event, nil
}

func (s *EventService) UpdateEvent(event *entity.Event) (*entity.Event, error) {
	if err := s.eventRepository.UpdateEvent(event); err != nil {
		return nil, err
	}
	return event, nil
}

func (s *EventService) DeleteEvent(id string) error {
	if err := s.eventRepository.DeleteEvent(id); err != nil {
		return err
	}
	return nil
}

func (s *EventService) UpdateSatusEvent(status request.UpdateStatusRequest) error {
	if err := s.eventRepository.UpdateSatusEvent(status); err != nil {
		return err
	}
	return nil
}

func (s *EventService) GetEventsByManagementRequired(managementRequired bool) ([]*entity.Event, error) {
	return s.eventRepository.GetEventsByManagementRequired(managementRequired)
}
