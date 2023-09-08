package entity

import (
	"time"
)

type Event struct {
	ID          uint      `json:"id" gorm:"id"`
	Name        string    `json:"name" gorm:"name"`
	TypeEventId uint      `json:"type_event_id" gorm:"type_event_id"`
	Description string    `json:"description" gorm:"description"`
	Date        time.Time `json:"date" gorm:"type:timestamp with time zone;default:current_timestamp"`
	Status      string    `json:"status" gorm:"status"`
}
