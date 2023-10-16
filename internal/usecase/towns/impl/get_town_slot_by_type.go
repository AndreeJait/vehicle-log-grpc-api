package impl

import (
	"context"
	"vehicle-log-grpc-api/internal/model"
)

func (u useCase) GetTownSlotByType(ctx context.Context, townCode, typeVehicle string) (model.TownSlot, error) {
	return u.repo.GetTownSlotRepository().GetByTownCodeAndVehicleType(ctx, townCode, typeVehicle)
}
