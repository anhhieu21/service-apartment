package repository

import (
	"apartment/internal/models"
	"gorm.io/gorm"
)

type CustomerRepository interface {
	GetUserByName(username string) *models.Customer
	GetUserById(id string) *models.Customer
	CreateCustomer(customer models.Customer) (bool, error)
}

type CustomerRepositoryImpl struct {
	Db *gorm.DB
}

// GetUserById implements CustomerRepository.
func (c *CustomerRepositoryImpl) GetUserById(id string) *models.Customer {
	customer := models.Customer{}
	result := c.Db.Where("id = ?", id).First(&customer)
	if result.Error != nil {
		return nil
	}
	return &customer
}

// CreateCustomer implements CustomerRepository.
func (c *CustomerRepositoryImpl) CreateCustomer(customer models.Customer) (bool, error) {
	result := c.Db.Create(&customer)
	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}

// GetUser implements CustomerRepository.
func (c *CustomerRepositoryImpl) GetUserByName(username string) *models.Customer {
	customer := models.Customer{}
	result := c.Db.Where("name = ?", username).First(&customer)
	if result.Error != nil {
		return nil
	}
	return &customer
}

func NewCustomerRepositoryImpl(db *gorm.DB) CustomerRepository {
	return &CustomerRepositoryImpl{Db: db}
}
