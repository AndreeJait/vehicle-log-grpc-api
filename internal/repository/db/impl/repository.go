package impl

import (
	"context"
	"github.com/uptrace/bun"
	"vehicle-log-grpc-api/internal/repository/db"
	"vehicle-log-grpc-api/internal/repository/db/log"
	logImpl "vehicle-log-grpc-api/internal/repository/db/log/impl"
	"vehicle-log-grpc-api/internal/repository/db/town"
	townImpl "vehicle-log-grpc-api/internal/repository/db/town/impl"
	"vehicle-log-grpc-api/internal/repository/db/town_slot"
	townSlotImpl "vehicle-log-grpc-api/internal/repository/db/town_slot/impl"
	"vehicle-log-grpc-api/internal/repository/db/types"
	typeImpl "vehicle-log-grpc-api/internal/repository/db/types/impl"
	"vehicle-log-grpc-api/internal/repository/db/vehicle"
	vehicleImpl "vehicle-log-grpc-api/internal/repository/db/vehicle/impl"
)

type repository struct {
	db         *bun.DB
	dbExecutor db.DBI
}

func (r *repository) GetTypeRepository() types.TypesRepository {
	return typeImpl.NewTypeRepository(r.getInstance())
}

func (r *repository) GetTownSlotRepository() town_slot.TownSlotRepository {
	return townSlotImpl.NewTownSlotRepository(r.getInstance())
}

func (r *repository) GetLogRepo() log.LogsRepository {
	return logImpl.NewLogRepository(r.getInstance())
}

func (r *repository) GetTownRepo() town.TownsRepository {
	return townImpl.NewTownRepository(r.getInstance())
}

func (r *repository) GetVehicleRepo() vehicle.VehiclesRepository {
	return vehicleImpl.NewVehicleRepository(r.getInstance())
}

func (r *repository) DoInTransaction(ctx context.Context, txFunc db.InTransaction) (out interface{}, err error) {
	var tx bun.Tx
	registry := r
	if r.dbExecutor == nil {
		tx, err = r.db.BeginTx(ctx, nil)
		if err != nil {
			return
		}
		defer func() {
			if p := recover(); p != nil {
				_ = tx.Rollback()
				panic(p)
			} else if err != nil {
				rErr := tx.Rollback()
				if rErr != nil {
					err = rErr
				}
			} else {
				err = tx.Commit()
			}
		}()

		registry = &repository{
			db:         r.db,
			dbExecutor: tx,
		}
	}

	out, err = txFunc(ctx, registry)

	return
}

func (r *repository) getInstance() db.DBI {
	if r.dbExecutor != nil {
		return r.dbExecutor
	}
	return r.db
}

func NewRepo(db *bun.DB) db.Repository {
	return &repository{
		db: db,
	}
}
