package models

type Admin struct{
	ID string `gorm:"primarykey"`
	UserName string `gorm:"unique"`
	Password string `gorm:"unique"`
}