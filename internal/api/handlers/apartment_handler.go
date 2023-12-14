package handlers

import (
	"apartment/internal/api/services"
	"apartment/internal/models"
	"apartment/internal/utils"
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
	result, err := s.apartmentService.GetApartment(req.GetId())
	// msg := utils.ErrorMessage(err, "success")
	return &pb.GetApartmentResponse{
		Apartment: &pb.Apartment{
			Id:     result.ID,
			Number: result.Number,
			Status: string(result.Status),
		},
	}, err
}
func (s *ApartmentHandler) CreateApartment(ctx context.Context, req *pb.CreateApartmentRequest) (*pb.CreateApartmentResponse, error) {
	aparment := models.Apartment{
		Number: req.Apartment.GetNumber(),
	}
	success, err := s.apartmentService.CreateApartment(aparment)
	msg := utils.ErrorMessage(err, "created")
	return &pb.CreateApartmentResponse{
		Success:      success,
		ErrorMessage: msg,
	}, err
}
func (s *ApartmentHandler) UpdateApartment(ctx context.Context, req *pb.UpdateApartmentRequest) (*pb.UpdateApartmentResponse, error) {
	aparment := models.Apartment{
		ID:     req.Apartment.GetId(),
		Number: req.Apartment.GetNumber(),
		Status: models.ApartmentStatus(req.Apartment.GetStatus()),
	}
	success, err := s.apartmentService.UpdateApartment(aparment)
	msg := utils.ErrorMessage(err, "updated")
	return &pb.UpdateApartmentResponse{
		Success:      success,
		ErrorMessage: msg,
	}, err
}
func (s *ApartmentHandler) DeleteApartment(ctx context.Context, req *pb.DeleteApartmentRequest) (*pb.DeleteApartmentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteApartment not implemented")
}
