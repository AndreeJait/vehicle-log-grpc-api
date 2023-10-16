package vehicle

import (
	"context"
	"vehicle-log-grpc-api/internal/model"
)

type VehiclesRepository interface {
	GetAllVehicle(ctx context.Context) ([]model.Vehicle, error)
	CheckByPlatNumber(ctx context.Context, platNumber string) (bool, error)
	GetVehicleByPlatNumber(ctx context.Context, platNumber string) (model.Vehicle, error)
	InsertVehicle(ctx context.Context, vehicle model.Vehicle) error
}
