package impl

import (
	"vehicle-log-grpc-api/internal/repository/db"
	"vehicle-log-grpc-api/internal/repository/db/types"
)

type repository struct {
	db db.DBI
}

func NewTypeRepository(db db.DBI) types.TypesRepository {
	return &repository{
		db: db,
	}
}
