package impl

import (
	"context"
	"encoding/json"
	"vehicle-log-grpc-api/internal/model"
)

func (u useCase) GetAllTypes(ctx context.Context) ([]model.Type, error) {
	result := make([]model.Type, 0)
	check, redis, err := u.redis.CheckAndGetTypesList(ctx)
	if err != nil {
		return result, err
	}
	if check {
		err = json.Unmarshal([]byte(redis), &result)
		if err != nil {
			return result, err
		}
		return result, nil
	}
	typesFromRepo, err := u.repo.GetTypeRepository().GetAllTypes(ctx)
	if err != nil {
		return result, err
	}
	jsonStrByte, err := json.Marshal(&typesFromRepo)
	if err != nil {
		return typesFromRepo, err
	}
	err = u.redis.SetListTypes(ctx, string(jsonStrByte))
	return typesFromRepo, err
}
