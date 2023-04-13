package service

import (
	"gitlab.com/nodejs_vs_go/go_api_gateway/config"
	"gitlab.com/nodejs_vs_go/go_api_gateway/genproto/driver_service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceManagerI interface {
	DriverService() driver_service.DriverServiceClient
	CarService() driver_service.CarServiceClient
}

type grpcClients struct {
	driverService driver_service.DriverServiceClient
	carService driver_service.CarServiceClient
}

func NewGrpcClients(cfg config.Config) (ServiceManagerI, error) {
	connDriverService, err := grpc.Dial(
		cfg.DriverServiceHost +":"+ cfg.DriverServicePort,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}

	return &grpcClients{
		driverService: driver_service.NewDriverServiceClient(connDriverService),
		carService: driver_service.NewCarServiceClient(connDriverService),
	}, nil
}

func (g *grpcClients) DriverService() driver_service.DriverServiceClient {
	return g.driverService
}
 
func (g *grpcClients) CarService() driver_service.CarServiceClient {
	return g.carService
}