package impl

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"strconv"
	rredis "vehicle-log-grpc-api/internal/repository/redis"
)

func (r rRedis) CheckAndGetTownAvailable(ctx context.Context, townCode, vehicleType string) (bool, int, error) {
	result, err := r.rdb.Get(ctx, fmt.Sprintf("%s_%s_%s", rredis.TypeAvailable, townCode, vehicleType)).Result()
	if err == redis.Nil {
		return false, 0, nil
	} else if err != nil {
		return false, 0, err
	}
	converted, _ := strconv.Atoi(result)
	return true, converted, nil
}
