package impl

import (
	"context"
	"github.com/pkg/errors"
	"log"
	"time"
	"vehicle-log-grpc-api/internal/common/utils"
	rredis "vehicle-log-grpc-api/internal/repository/redis"
)

func (r rRedis) SetListTypes(ctx context.Context, typeListStr string) error {
	err := r.rdb.Set(ctx, rredis.TypeCache, typeListStr, 0).Err()
	if err != nil {
		return errors.Wrap(err, "failed to insert data to redis")
	}
	log.Printf("inserted to redis %s: %s", time.Now().Format(utils.DefaultFormatDate), typeListStr)
	return nil
}
