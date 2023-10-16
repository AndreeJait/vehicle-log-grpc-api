package impl

import (
	"context"
	"github.com/pkg/errors"
	"github.com/uptrace/bun"
	"vehicle-log-grpc-api/internal/model"
)

func (r repository) GetByTownCodeAndVehicleType(ctx context.Context, townCode, vehicleType string) (model.TownSlot, error) {
	var townSlot model.TownSlot
	err := r.db.NewSelect().
		Model(&townSlot).
		Relation("Town", func(query *bun.SelectQuery) *bun.SelectQuery {
			return query.Where("town_code = ?", townCode)
		}).
		Relation("Type", func(query *bun.SelectQuery) *bun.SelectQuery {
			return query.Where("name = ?", vehicleType)
		}).
		Scan(ctx)
	if err != nil {
		return townSlot, errors.Wrap(err, "failed to get town")
	}
	return townSlot, nil
}
