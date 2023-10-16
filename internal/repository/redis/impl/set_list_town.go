package impl

import (
	"context"
	"github.com/pkg/errors"
	"log"
	"time"
	"vehicle-log-grpc-api/internal/common/utils"
	rredis "vehicle-log-grpc-api/internal/repository/redis"
)

func (r rRedis) SetListTown(ctx context.Context, townListStr string) error {

	timeIntervalEOD := utils.GetEndEndOfDateTime(r.cfg.Cache.TimeRefresh)
	err := r.rdb.Set(ctx, rredis.TownCache, townListStr, timeIntervalEOD).Err()
	if err != nil {
		return errors.Wrap(err, "failed to insert data to redis")
	}
	log.Printf("inserted to redis %s: %s", time.Now().Format(utils.DefaultFormatDate), townListStr)
	return nil
}
