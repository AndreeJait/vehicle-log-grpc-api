package impl

import (
	"context"
	"github.com/redis/go-redis/v9"
	rredis "vehicle-log-grpc-api/internal/repository/redis"
)

func (r rRedis) CheckAndGetTypesList(ctx context.Context) (bool, string, error) {
	result, err := r.rdb.Get(ctx, rredis.TypeCache).Result()
	if err == redis.Nil {
		return false, "", nil
	} else if err != nil {
		return false, "", err
	}
	return true, result, nil
}
