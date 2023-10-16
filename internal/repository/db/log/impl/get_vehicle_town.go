package impl

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/uptrace/bun"
	"vehicle-log-grpc-api/internal/model"
	"vehicle-log-grpc-api/internal/usecase/vehicle"
)

func (r repository) GetVehicleInTown(ctx context.Context, towerCode string, filter vehicle.FilterVehicleInTown) ([]model.Log, error) {
	logs := make([]model.Log, 0)
	q := r.db.NewSelect().Model(&logs).
		Relation("TownSlot").
		Relation("Vehicle", func(query *bun.SelectQuery) *bun.SelectQuery {
			if filter.PlatNumber != "" {
				query = query.Where("plat_number ILIKE ?", fmt.Sprintf("%%%s%%", filter.PlatNumber))
			}
			return query
		}).
		Relation("TownSlot.Town", func(query *bun.SelectQuery) *bun.SelectQuery {
			query = query.Where("town_code = ?", towerCode)
			return query
		}).
		Relation("TownSlot.Type", func(query *bun.SelectQuery) *bun.SelectQuery {
			if filter.VehicleType != "" {
				query = query.Where("name ILIKE ?", fmt.Sprintf("%%%s%%", filter.VehicleType))
			}
			return query
		}).Order("date_at DESC", "time_in DESC")

	err := q.Scan(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to fetch logs")
	}
	return logs, nil
}
