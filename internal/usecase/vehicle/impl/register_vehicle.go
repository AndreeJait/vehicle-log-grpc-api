package impl

import (
	"context"
	"github.com/google/uuid"
	"vehicle-log-grpc-api/internal/model"
)

func (uc useCase) RegisterVehicle(ctx context.Context, platNumber string) error {
	vehicleRepo := uc.repo.GetVehicleRepo()
	vehicle := model.Vehicle{
		ID:         uuid.NewString(),
		PlatNumber: platNumber,
	}
	err := vehicleRepo.InsertVehicle(ctx, vehicle)
	if err != nil {
		return err
	}
	return nil
}
