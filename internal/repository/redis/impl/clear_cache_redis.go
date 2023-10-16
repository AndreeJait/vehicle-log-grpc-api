package impl

import (
	"context"
)

func (r rRedis) ClearCacheRedis(ctx context.Context, types ...string) error {
	err := r.rdb.Del(ctx, types...).Err()
	return err
}
