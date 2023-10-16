package impl

import (
	"context"
	"github.com/redis/go-redis/v9"
	"log"
	"vehicle-log-grpc-api/config"
	rd "vehicle-log-grpc-api/internal/repository/redis"
)

type rRedis struct {
	rdb *redis.Client
	cfg *config.Config
}

func New(cfg *config.Config) rd.Redis {
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.Addr,
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.DB,
	})
	pong, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}
	log.Printf("connected to redis [%s]", pong)
	return &rRedis{
		rdb: rdb,
		cfg: cfg,
	}
}
