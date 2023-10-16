package towns

import (
	"context"
	"vehicle-log-grpc-api/internal/model"
)

type TownUseCase interface {
	GetAllTown(ctx context.Context) ([]model.Town, error)
	GetTownByCode(ctx context.Context, townCode string) (model.Town, error)
	GetAllTownSlotByType(ctx context.Context, townCode string) ([]model.TownSlot, error)
	GetTownSlotByType(ctx context.Context, townCode, typeVehicle string) (model.TownSlot, error)
}
