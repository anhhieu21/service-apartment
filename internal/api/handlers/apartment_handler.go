package handlers

import (
	"apartment/internal/api/services"
	"apartment/internal/models"
	"apartment/pb"
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ApartmentHandler struct {
	pb.UnimplementedApartmentServiceServer
	apartmentService services.ApartmentService
}

func NewApartment(apartmentService services.ApartmentService) *ApartmentHandler {
	return &ApartmentHandler{apartmentService: apartmentService}
}

func (s *ApartmentHandler) GetApartment(ctx context.Context, req *pb.GetApartmentRequest) (*pb.GetApartmentResponse, error) {

	return nil, status.Errorf(codes.Unimplemented, "method GetApartment not implemented")
}
func (s *ApartmentHandler) CreateApartment(ctx context.Context, req *pb.CreateApartmentRequest) (*pb.CreateApartmentResponse, error) {
	msg := ""
	aparment := models.Apartment{
		Number: req.Apartment.GetNumber(),
	}
	success, err := s.apartmentService.CreateApartment(aparment)
	if err != nil {
		msg = err.Error()
	} else {
		msg = "created"
	}
	return &pb.CreateApartmentResponse{
		Success:      success,
		ErrorMessage: msg,
	}, err
}
func (s *ApartmentHandler) UpdateApartment(ctx context.Context, req *pb.UpdateApartmentRequest) (*pb.UpdateApartmentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateApartment not implemented")
}
func (s *ApartmentHandler) DeleteApartment(ctx context.Context, req *pb.DeleteApartmentRequest) (*pb.DeleteApartmentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteApartment not implemented")
}
