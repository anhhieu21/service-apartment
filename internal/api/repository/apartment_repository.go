package repository

import (
	"apartment/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ApartmentRepository interface {
	CreateApartment(apartment models.Apartment) (bool, error)
	GetApartment(id string) (models.Apartment, error)
	UpdateApartment(apartment models.Apartment) (bool, error)
	DeleteApartment(id string) (bool, error)
}

type ApartmentRepositoryImpl struct {
	Db *gorm.DB
}

// CreateApartment implements ApartmentRepository.
func (a *ApartmentRepositoryImpl) CreateApartment(apartment models.Apartment) (bool, error) {
	apartment = models.Apartment{
		ID:     uuid.NewString(),
		Number: apartment.Number,
		Status: models.VacantReady,
	}
	result := a.Db.Create(&apartment)
	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}

// DeleteApartment implements ApartmentRepository.
func (a *ApartmentRepositoryImpl) DeleteApartment(id string) (bool, error) {

	panic("unimplemented")
}

// GetApartment implements ApartmentRepository.
func (a *ApartmentRepositoryImpl) GetApartment(id string) (models.Apartment, error) {
	user := models.Apartment{}
	result := a.Db.Where("id = ?", id).Find(&user)
	if result.Error != nil {
		return user, result.Error
	}
	return user, nil
}

// UpdateApartment implements ApartmentRepository.
func (a *ApartmentRepositoryImpl) UpdateApartment(apartment models.Apartment) (bool, error) {
	id := apartment.ID
	result := a.Db.Model(&models.Apartment{}).Select("status").Where("id = ?", id).Updates(apartment)
	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}

func NewApartmentRepositoryImpl(db *gorm.DB) ApartmentRepository {
	return &ApartmentRepositoryImpl{Db: db}
}
