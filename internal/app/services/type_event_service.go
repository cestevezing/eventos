package services

import (
	"github.com/cestevezing/eventos/internal/app/models/entity"
	"github.com/cestevezing/eventos/internal/app/repository"
)

type TypeEventService struct {
	typeEventRepository *repository.TypeEventRepository
}

func NewTypeEventService(typeEventRepository *repository.TypeEventRepository) *TypeEventService {
	return &TypeEventService{
		typeEventRepository: typeEventRepository,
	}
}

func (s *TypeEventService) GetTypesEvent() ([]*entity.TypeEvent, error) {
	return s.typeEventRepository.GetTypesEvent()
}

func (s *TypeEventService) CreateTypeEvent(typeEvent *entity.TypeEvent) (*entity.TypeEvent, error) {
	if err := s.typeEventRepository.CreateTypeEvent(typeEvent); err != nil {
		return nil, err
	}
	return typeEvent, nil
}

func (s *TypeEventService) UpdateTypeEvent(typeEvent *entity.TypeEvent) (*entity.TypeEvent, error) {
	if err := s.typeEventRepository.UpdateTypeEvent(typeEvent); err != nil {
		return nil, err
	}
	return typeEvent, nil
}

func (s *TypeEventService) DeleteTypeEvent(id string) error {
	if err := s.typeEventRepository.DeleteTypeEvent(id); err != nil {
		return err
	}
	return nil
}
