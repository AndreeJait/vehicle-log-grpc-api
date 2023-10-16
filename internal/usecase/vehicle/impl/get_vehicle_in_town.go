package impl

import (
	"context"
	"vehicle-log-grpc-api/internal/model"
	"vehicle-log-grpc-api/internal/usecase/vehicle"
)

func (uc useCase) GetVehicleInTown(ctx context.Context, townCode string, filter vehicle.FilterVehicleInTown) ([]model.Log, error) {
	logs, err := uc.repo.GetLogRepo().GetVehicleInTown(ctx, townCode, filter)
	if err != nil {
		return logs, err
	}

	return logs, nil
}
