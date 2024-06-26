package models

import "time"

type Application struct {
	CustomModel
	Position    string    `json:"position" gorm:"column:position"`
	Company     string    `json:"company" gorm:"column:company"`
	Type        string    `json:"type" gorm:"column:type"`
	Location    string    `json:"location" gorm:"column:location"`
	Note        string    `json:"note,omitempty" gorm:"column:note"`
	Status      string    `json:"status,omitempty" gorm:"column:status"`
	AppliedDate time.Time `json:"applied_date" gorm:"column:applied_date"`
	WorkType    string    `json:"worktype" gorm:"column:worktype"`
}
