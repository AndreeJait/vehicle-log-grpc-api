package vehicle

import (
	"context"
	"vehicle-log-grpc-api/internal/model"
)

type VehiclesUseCase interface {
	RegisterVehicle(ctx context.Context, platNumber string) error
	CheckVehicle(ctx context.Context, platNumber string) (bool, error)
	GetVehicleInTown(ctx context.Context, townCode string, filter FilterVehicleInTown) ([]model.Log, error)
}
