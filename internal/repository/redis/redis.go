package redis

import "context"

type Redis interface {
	SetListTown(ctx context.Context, townListStr string) error
	CheckAndGetTownList(ctx context.Context) (bool, string, error)
	SetListTypes(ctx context.Context, typeListStr string) error
	CheckAndGetTypesList(ctx context.Context) (bool, string, error)
	ClearCacheRedis(ctx context.Context, types ...string) error
	ClearCacheRedisPattern(ctx context.Context, pattern string) error

	CheckAndGetTownQuantity(ctx context.Context, townCode, vehicleType string) (bool, int, error)
	SetTownQuantity(ctx context.Context, townCode, vehicleType string, value int) error

	CheckAndGetTownAvailable(ctx context.Context, townCode, vehicleType string) (bool, int, error)
	SetTownAvailable(ctx context.Context, townCode, vehicleType string, value int) error
}
