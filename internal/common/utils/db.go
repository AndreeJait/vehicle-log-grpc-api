package utils

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/extra/bundebug"
	"log"
	"vehicle-log-grpc-api/config"
)

func NewBunPostgresSQLCon(env config.Env, cfg config.Config) (*bun.DB, error) {
	dbEnv := cfg.DB
	connectionStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbEnv.Host, dbEnv.Port, dbEnv.User, dbEnv.Password, dbEnv.Name)
	hsqldb, err := sql.Open("postgres", connectionStr)
	if err != nil {
		return nil, errors.Wrap(err, "cannot open connection")
	}
	db := bun.NewDB(hsqldb, pgdialect.New())
	if env.IsLocal() {
		db.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))
	}
	log.Printf("connected to database")
	return db, nil
}
