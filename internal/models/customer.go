package models

type Customer struct {
	ID    string `gorm:"primarykey"`
	Name  string `gorm:"unique"`
	Phone string `gorm:"unique"`
}
