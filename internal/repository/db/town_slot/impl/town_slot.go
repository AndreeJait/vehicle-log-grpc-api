package impl

import (
	"vehicle-log-grpc-api/internal/repository/db"
	"vehicle-log-grpc-api/internal/repository/db/town_slot"
)

type repository struct {
	db db.DBI
}

func NewTownSlotRepository(db db.DBI) town_slot.TownSlotRepository {
	return &repository{
		db: db,
	}
}
