package repository

import (
	"apartment/internal/models"
	"apartment/pb"

	"gorm.io/gorm"
)

type CustomerRepository interface {
	Login(username, password string) (pb.LoginResponse, error)
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
	result := c.Db.Where("id = ?", id).First(customer)
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

// Login implements CustomerRepository.
func (*CustomerRepositoryImpl) Login(username string, password string) (pb.LoginResponse, error) {
	return pb.LoginResponse{}, nil
}

func NewCustomerRepositoryImpl(db *gorm.DB) CustomerRepository {
	return &CustomerRepositoryImpl{Db: db}
}
