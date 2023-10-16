package impl

import (
	"vehicle-log-grpc-api/internal/repository/db"
	"vehicle-log-grpc-api/internal/repository/db/town"
)

type repository struct {
	db db.DBI
}

func NewTownRepository(db db.DBI) town.TownsRepository {
	return &repository{
		db: db,
	}
}
