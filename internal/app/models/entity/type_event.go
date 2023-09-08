package entity

type TypeEvent struct {
	ID                 uint    `json:"id" gorm:"id"`
	Name               string  `json:"name" gorm:"name"`
	Description        string  `json:"description" gorm:"description"`
	ManagementRequired bool    `json:"management_required" gorm:"management_required"`
	Events             []Event `json:"-" gorm:"foreignKey:TypeEventId"`
}
