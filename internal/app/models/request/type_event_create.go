package request

type TypeEventCreate struct {
	Name               string `json:"name" gorm:"name"`
	Description        string `json:"description" gorm:"description"`
	ManagementRequired bool   `json:"management_required" gorm:"management_required"`
}
