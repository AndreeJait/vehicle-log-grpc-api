package impl

import (
	"vehicle-log-grpc-api/internal/repository/db"
	"vehicle-log-grpc-api/internal/repository/db/log"
)

type repository struct {
	db db.DBI
}

func NewLogRepository(db db.DBI) log.LogsRepository {
	return &repository{
		db: db,
	}
}
