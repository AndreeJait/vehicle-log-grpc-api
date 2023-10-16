package impl

import (
	"context"
	"vehicle-log-grpc-api/internal/model"
)

func (u useCase) GetAllTownSlotByType(ctx context.Context, townCode string) ([]model.TownSlot, error) {
	return u.repo.GetTownSlotRepository().GetAllByTownCode(ctx, townCode)
}
