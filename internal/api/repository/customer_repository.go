package repository

import (
	"apartment/internal/models"

	"gorm.io/gorm"
)

type CustomerRepository interface {
	GetUserByName(username string) *models.Customer
	GetUserById(id string) *models.Customer
	CreateCustomer(customer models.Customer) (bool, error)
	GetApartment(id string) *models.Apartment
	GetApartments() *[]models.Apartment
}

type CustomerRepositoryImpl struct {
	Db *gorm.DB
}

// GetApartments implements CustomerRepository.
// find(&âprment) điều này cho phép db query điền trực tiếp dự liệu vào apartment với dự liệu có đc sau khi query
// & : lấy địa chỉ của apartment rồi tạo 1 con trỏ (*) tới đó 
// * : đc dùng để định hướng và truy cập sử đổi các data tại địa chỉ bộ nhớ đó , ở đây lại tại ví trí bộ nhớ của apartment
func (c *CustomerRepositoryImpl) GetApartments() *[]models.Apartment {
	apartments := []models.Apartment{}
	result := c.Db.Find(&apartments)
	if result.Error != nil {
		return nil
	}
	return &apartments
}

// GetApartment implements CustomerRepository.
func (c *CustomerRepositoryImpl) GetApartment(id string) *models.Apartment {
	apartment := models.Apartment{}
	result := c.Db.Where("id = ?", id).First(&apartment)
	if result.Error != nil {
		return nil
	}
	return &apartment
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
