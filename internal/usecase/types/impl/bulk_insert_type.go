package impl

import (
	"context"
	"github.com/google/uuid"
	"vehicle-log-grpc-api/internal/model"
	"vehicle-log-grpc-api/internal/repository/redis"
)

func (u useCase) BulkInsertType(ctx context.Context, types []string) error {
	typesModel := make([]model.Type, 0)
	for _, v := range types {
		typesModel = append(typesModel, model.Type{
			ID:   uuid.NewString(),
			Name: v,
		})
	}
	err := u.repo.GetTypeRepository().BulkInsert(ctx, typesModel)
	if err != nil {
		return err
	}
	err = u.redis.ClearCacheRedis(ctx, redis.TypeCache)
	return err
}
