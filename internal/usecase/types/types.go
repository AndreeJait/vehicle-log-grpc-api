package types

import (
	"context"
	"vehicle-log-grpc-api/internal/model"
)

type TypeUseCase interface {
	GetAllTypes(ctx context.Context) ([]model.Type, error)
	BulkInsertType(ctx context.Context, types []string) error
}
