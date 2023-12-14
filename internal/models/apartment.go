package models

type Apartment struct {
	ID     string `gorm:"primarykey"`
	Number string `gorm:"unique"`
	Status ApartmentStatus
}

// status : vacant, occupied

type ApartmentStatus string

const (
	VacantReady ApartmentStatus = "VR"
	Occupied    ApartmentStatus = "OC"
)
