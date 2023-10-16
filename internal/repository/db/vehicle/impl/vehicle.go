package impl

import (
	"vehicle-log-grpc-api/internal/repository/db"
	"vehicle-log-grpc-api/internal/repository/db/vehicle"
)

type repository struct {
	db db.DBI
}

func NewVehicleRepository(db db.DBI) vehicle.VehiclesRepository {
	return &repository{
		db: db,
	}
}
