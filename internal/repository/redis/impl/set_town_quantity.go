package impl

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"log"
	"time"
	"vehicle-log-grpc-api/internal/common/utils"
	rredis "vehicle-log-grpc-api/internal/repository/redis"
)

func (r rRedis) SetTownQuantity(ctx context.Context, townCode, vehicleType string, value int) error {
	timeIntervalEOD := utils.GetEndEndOfDateTime(r.cfg.Cache.TimeRefresh)
	err := r.rdb.Set(ctx, fmt.Sprintf("%s_%s_%s", rredis.TypeQuantity, townCode, vehicleType), value, timeIntervalEOD).Err()
	if err != nil {
		return errors.Wrap(err, "failed to insert data to redis")
	}
	log.Printf("inserted to redis %s: %d", time.Now().Format(utils.DefaultFormatDate), value)
	return nil
}
