package models

type User struct {
	CustomModel
	FullName string `json:"fullname,omitempty" gorm:"column:fullname"`
	Email    string `json:"email" gorm:"column:email"`
	Password string `json:"password" gorm:"column:password"`
	Status   string `json:"status" gorm:"column:status"`
}

type CustomModel struct {
	ID uint `gorm:"primaryKey"`
}
