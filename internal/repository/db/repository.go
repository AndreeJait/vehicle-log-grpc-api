package db

import (
	"context"
	"vehicle-log-grpc-api/internal/repository/db/log"
	"vehicle-log-grpc-api/internal/repository/db/town"
	"vehicle-log-grpc-api/internal/repository/db/town_slot"
	"vehicle-log-grpc-api/internal/repository/db/types"
	"vehicle-log-grpc-api/internal/repository/db/vehicle"
)

type InTransaction func(ctx context.Context, repo Repository) (interface{}, error)

type Repository interface {
	DoInTransaction(ctx context.Context, txFunc InTransaction) (out interface{}, err error)
	GetVehicleRepo() vehicle.VehiclesRepository
	GetTownRepo() town.TownsRepository
	GetLogRepo() log.LogsRepository
	GetTownSlotRepository() town_slot.TownSlotRepository
	GetTypeRepository() types.TypesRepository
}
