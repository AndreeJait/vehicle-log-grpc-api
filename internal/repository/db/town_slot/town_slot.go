package town_slot

import (
	"context"
	"vehicle-log-grpc-api/internal/model"
)

type TownSlotRepository interface {
	GetAllByTownCode(ctx context.Context, townCode string) ([]model.TownSlot, error)
	GetByTownCodeAndVehicleType(ctx context.Context, townCode, vehicleType string) (model.TownSlot, error)
}
