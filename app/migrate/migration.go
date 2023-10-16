package migrate

import (
	"context"
	"fmt"
	"github.com/rs/zerolog/log"
	migrate "github.com/rubenv/sql-migrate"
	"github.com/uptrace/bun"
	"vehicle-log-grpc-api/config"
	"vehicle-log-grpc-api/internal/common/utils"
	"vehicle-log-grpc-api/internal/repository/redis"
	redisImpl "vehicle-log-grpc-api/internal/repository/redis/impl"
)

const MigrationTypeUp = "up"
const MigrationTypeDown = "down"
const MigrationTypeFresh = "fresh"

type Migration struct {
	db    *bun.DB
	cfg   *config.Config
	redis redis.Redis
}

func New() *Migration {
	cfg := config.LoadDefault()
	db, err := utils.NewBunPostgresSQLCon(config.Env(cfg.Server.Mode), *cfg)
	if err != nil {
		panic(err)
	}
	r := redisImpl.New(cfg)
	return &Migration{
		cfg:   cfg,
		db:    db,
		redis: r,
	}
}

func (m *Migration) Start(migrateType string) {
	log.Printf("start migration %s", migrateType)

	if migrateType == MigrationTypeFresh {
		if config.Env(m.cfg.Server.Mode).IsProd() {
			log.Printf("cannot migrate fresh in production")
			return
		}
		err := m.redis.ClearCacheRedis(context.Background(), redis.TypeCache, redis.TownCache)
		if err != nil {
			panic(err)
		}

		err = m.redis.ClearCacheRedisPattern(context.Background(), fmt.Sprintf("%s*", redis.TypeQuantity))
		if err != nil {
			panic(err)
		}

		err = m.redis.ClearCacheRedisPattern(context.Background(), fmt.Sprintf("%s*", redis.TypeAvailable))
		if err != nil {
			panic(err)
		}

		log.Print("drop database")
		_, err = m.db.Exec("DROP SCHEMA public CASCADE; CREATE SCHEMA public;")
		if err != nil {
			panic(err)
		}
		m.db, err = utils.NewBunPostgresSQLCon(config.Env(m.cfg.Server.Mode), *m.cfg)
		if err != nil {
			panic(err)
		}
	}

	migrations := &migrate.FileMigrationSource{
		Dir: "./scripts/migrations",
	}

	var direction migrate.MigrationDirection
	switch migrateType {
	case MigrationTypeUp:
		direction = migrate.Up
	case MigrationTypeDown:
		direction = migrate.Down
	case MigrationTypeFresh:
		direction = migrate.Up
	}

	count, err := migrate.Exec(m.db.DB, "postgres", migrations, direction)
	if err != nil {
		panic(err)
	}
	log.Printf("applied %d migrations", count)
}
