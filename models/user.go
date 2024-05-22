package models

type User struct {
	CustomModel
	FullName string `gorm:"column:fullname"`
	Email    string `gorm:"column:email"`
	Password string `gorm:"column:password"`
}

type CustomModel struct {
	ID uint `gorm:"primaryKey"`
}
