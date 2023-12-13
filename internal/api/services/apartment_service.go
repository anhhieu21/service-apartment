package services

import (
	"apartment/internal/api/repository"
	"apartment/internal/models"
)

type ApartmentService interface {
	CreateApartment(apartment models.Apartment) (bool, error)
	GetApartment(id string) (models.Apartment, error)
	UpdateApartment(apartment models.Apartment) (bool, error)
	DeleteApartment(id string) (bool, error)
}

type ApartmentServiceImpl struct {
	AparmentRepo repository.ApartmentRepository
}

// CreateApartment implements ApartmentService.
func (a *ApartmentServiceImpl) CreateApartment(apartment models.Apartment) (bool, error) {
	success, err := a.AparmentRepo.CreateApartment(apartment)
	return success, err
}

// DeleteApartment implements ApartmentService.
func (a *ApartmentServiceImpl) DeleteApartment(id string) (bool, error) {
	panic("unimplemented")
}

// GetApartment implements ApartmentService.
func (a *ApartmentServiceImpl) GetApartment(id string) (models.Apartment, error) {
	panic("unimplemented")
}

// UpdateApartment implements ApartmentService.
func (a *ApartmentServiceImpl) UpdateApartment(apartment models.Apartment) (bool, error) {
	panic("unimplemented")
}

func NewApartmentServiceImpl(repo repository.ApartmentRepository) ApartmentService {
	return &ApartmentServiceImpl{AparmentRepo: repo}
}
