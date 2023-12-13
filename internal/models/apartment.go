package models

type Apartment struct {
	ID     string `gorm:"primarykey"`
	Number string `gorm:"unique"`
}
