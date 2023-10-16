package impl

import (
	"vehicle-log-grpc-api/internal/handlers/vehicle"
	vehicleUC "vehicle-log-grpc-api/internal/usecase/vehicle"
)

type handlers struct {
	vcUseCase vehicleUC.VehiclesUseCase
}

func New(vcUC vehicleUC.VehiclesUseCase) vehicle.Handler {
	return &handlers{
		vcUseCase: vcUC,
	}
}
