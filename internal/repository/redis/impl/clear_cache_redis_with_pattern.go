package impl

import "context"

func (r rRedis) ClearCacheRedisPattern(ctx context.Context, pattern string) error {
	keys, err := r.rdb.Keys(ctx, pattern).Result()
	if err != nil {
		return err
	}
	if len(keys) > 0 {
		err = r.rdb.Del(ctx, keys...).Err()
		if err != nil {
			return err
		}
	}
	return nil
}
