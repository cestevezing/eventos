package repository

import (
	"errors"

	"github.com/cestevezing/eventos/internal/app/models/entity"
	"gorm.io/gorm"
)

type TypeEventRepository struct {
	db *gorm.DB
}

func NewTypeEventRepository(db *gorm.DB) *TypeEventRepository {
	return &TypeEventRepository{
		db: db,
	}
}

func (er *TypeEventRepository) GetTypesEvent() ([]*entity.TypeEvent, error) {
	var typesEvent []*entity.TypeEvent
	err := er.db.Find(&typesEvent).Error
	return typesEvent, err
}

func (er *TypeEventRepository) GetTypeEventById(id uint) (*entity.TypeEvent, error) {
	var typeEvent entity.TypeEvent
	if err := er.db.Where("id = ?", id).First(&typeEvent).Error; err != nil {
		return nil, err
	}
	return &typeEvent, nil
}

func (er *TypeEventRepository) CreateTypeEvent(typeEvent *entity.TypeEvent) error {
	return er.db.Create(typeEvent).Error
}

func (er *TypeEventRepository) UpdateTypeEvent(typeEvent *entity.TypeEvent) error {
	var resul entity.TypeEvent
	if err := er.db.Where("id = ?", typeEvent.ID).First(&resul).Error; err != nil {
		return errors.New("no se encontró el tipo de evento a editar")
	}
	return er.db.Save(typeEvent).Error
}

func (er *TypeEventRepository) DeleteTypeEvent(id string) error {
	var resul entity.TypeEvent
	if err := er.db.Where("id = ?", id).First(&resul).Error; err != nil {
		return errors.New("no se encontró el tipo de evento a eliminar")
	}
	return er.db.Delete(&entity.TypeEvent{}, id).Error
}
