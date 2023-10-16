package town

import (
	"context"
	"vehicle-log-grpc-api/internal/model"
)

type TownsRepository interface {
	GetAllTown(ctx context.Context) ([]model.Town, error)
	GetTownByCode(ctx context.Context, townCode string) (model.Town, error)
	GetTownByID(ctx context.Context, townID string) (model.Town, error)
}
