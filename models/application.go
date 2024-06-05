package models

import "time"

type Application struct {
	CustomModel
	Position    string    `json:"position" gorm:"column:postion"`
	Company     string    `json:"company" gorm:"column:company"`
	Type        string    `json:"type" gorm:"column:type"`
	Location    string    `json:"location" gorm:"column:location"`
	Note        string    `json:"note,omitempty" gorm:"column:note"`
	Status      string    `json:"status" gorm:"column:status"`
	AppliedDate time.Time `json:"applied_date" gorm:"column:applied_date"`
	WorkType    string    `json:"work_type" gorm:"column:work_type"`
}
