package services

import (
	"apartment/internal/api/repository"
	"apartment/internal/models"
	"apartment/internal/utils"
	"apartment/pb"
	"context"
	"time"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CustomerService interface {
	Login(username, passowrd string) (*pb.LoginResponse, error)
	Register(customer models.Customer) (*pb.RegisterResponse, error)
	GetCustomer(ctx context.Context) *pb.FindMeResponse
}

type CustomerServiceImpl struct {
	CustomerRepo repository.CustomerRepository
}

// GetCustomer implements CustomerService.
func (c *CustomerServiceImpl) GetCustomer(ctx context.Context) *pb.FindMeResponse {
	id, err := utils.GetCustomerIdFromContext(ctx)
	if err != nil{
		return &pb.FindMeResponse{}
	}
	customer := c.CustomerRepo.GetUserById(id)
	
	return &pb.FindMeResponse{
		Customer: &pb.Customer{
			Id:    customer.ID,
			Name:  customer.Name,
			Phone: customer.Phone,
		},
	}
}

// Register implements CustomerService.
func (c *CustomerServiceImpl) Register(customer models.Customer) (*pb.RegisterResponse, error) {
	newCustomer := c.CustomerRepo.GetUserByName(customer.Name)

	if newCustomer != nil {
		return nil, status.Errorf(codes.AlreadyExists, "Customer already")
	}

	password := utils.HashPassword(customer.Passowrd)

	uuid := uuid.NewString()

	customer.Passowrd = password
	customer.ID = uuid
	success, err := c.CustomerRepo.CreateCustomer(customer)

	return &pb.RegisterResponse{
		Success: success,
	}, err
}

// Login implements CustomerService.
func (c *CustomerServiceImpl) Login(username string, passowrd string) (*pb.LoginResponse, error) {
	customer := c.CustomerRepo.GetUserByName(username)

	if customer == nil {
		return nil, status.Errorf(codes.NotFound, "Customer not found")
	}

	if err := utils.CompareHashAndPassword(customer.Passowrd, passowrd); err != nil {
		return nil, err
	}

	expired_token := time.Now().Add(time.Hour * 24).Unix()
	access_token, err := utils.GenerateAccessToken(customer.ID, expired_token)
	utils.ErrorMessage(err, "Gen access token failed")

	refresh_token, err := utils.GenerateRefreshToken(customer.ID)
	utils.ErrorMessage(err, "Gen refresh token failed")

	return &pb.LoginResponse{
		AccessToken:  access_token,
		RefreshToken: refresh_token,
		ExpiredToken: expired_token,
		Success:      true,
	}, err
}

func NewCustomerServiceImpl(repo repository.CustomerRepository) CustomerService {
	return &CustomerServiceImpl{CustomerRepo: repo}
}
