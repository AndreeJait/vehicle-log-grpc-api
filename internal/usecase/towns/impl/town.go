package impl

import (
	"vehicle-log-grpc-api/config"
	"vehicle-log-grpc-api/internal/repository/db"
	"vehicle-log-grpc-api/internal/repository/redis"
	"vehicle-log-grpc-api/internal/usecase/towns"
)

type useCase struct {
	repo  db.Repository
	cfg   *config.Config
	redis redis.Redis
}

func New(repo db.Repository, cfg *config.Config, redis redis.Redis) towns.TownUseCase {
	return &useCase{
		repo:  repo,
		cfg:   cfg,
		redis: redis,
	}
}
