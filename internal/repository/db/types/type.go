package types

import (
	"context"
	"vehicle-log-grpc-api/internal/model"
)

type TypesRepository interface {
	GetAllTypes(ctx context.Context) ([]model.Type, error)
	BulkInsert(ctx context.Context, types []model.Type) error
	GetByName(ctx context.Context, name string) (model.Type, error)
}
