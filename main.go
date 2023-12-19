package main

import (
	"apartment/db"
	"apartment/internal/api/handlers"
	"apartment/internal/api/repository"
	"apartment/internal/api/services"
	"apartment/pb"
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

func init() {
	db.DatabasaConnection()
}

var (
	port = flag.Int("port", 3000, "gRPC server port")
)

func main() {
	fmt.Println("gRPC server running ...")
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	apartmentRepoImpl := repository.NewApartmentRepositoryImpl(db.DB)
	apartmentServiceImpl := services.NewApartmentServiceImpl(apartmentRepoImpl)
	apartmentHandler := handlers.NewApartment(apartmentServiceImpl)

	customerRepoImpl := repository.NewCustomerRepositoryImpl(db.DB)
	customerServiceImpl := services.NewCustomerServiceImpl(customerRepoImpl)
	customerHandler := handlers.NewCustomerHandler(customerServiceImpl)
	s := grpc.NewServer()
	pb.RegisterApartmentServiceServer(s, apartmentHandler)
	pb.RegisterCustomerServiceServer(s, customerHandler)

	log.Printf("Server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve : %v", err)
	}
}
