package handlers

import (
	"apartment/internal/api/services"
	"apartment/internal/models"
	"apartment/pb"
	"context"
)

type CustomerHandler struct {
	pb.UnimplementedCustomerServiceServer
	customerService services.CustomerService
}

func NewCustomerHandler(customerService services.CustomerService) *CustomerHandler {
	return &CustomerHandler{customerService: customerService}
}

func (c *CustomerHandler) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	customer := models.Customer{
		Name:     req.Customer.GetName(),
		Phone:    req.Customer.GetPhone(),
		Passowrd: req.Customer.GetPassword(),
	}
	result, err := c.customerService.Register(customer)

	return result, err
}
func (c *CustomerHandler) DeleteCustomer(ctx context.Context, req *pb.DeleteCustomerRequest) (*pb.DeleteCustomerResponse, error) {
	panic("xxxx")
}
func (c *CustomerHandler) GetCustomer(ctx context.Context, req *pb.GetCustomerRequest) (*pb.GetCustomerResponse, error) {
	panic("xxxx")
}
func (c *CustomerHandler) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	result, err := c.customerService.Login(req.GetUsername(), req.GetPassword())
	return result, err
}
func (c *CustomerHandler) UpdateCustomer(ctx context.Context, req *pb.UpdateCustomerRequest) (*pb.UpdateCustomerResponse, error) {
	panic("xxxx")
}
func (c *CustomerHandler) FindMe(ctx context.Context, req *pb.FindMeRequest) (*pb.FindMeResponse, error) {
	customer := c.customerService.GetCustomer(ctx)
	return customer, nil
}
