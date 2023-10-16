package impl

import (
	"context"
	"database/sql"
	"github.com/pkg/errors"
	"github.com/uptrace/bun"
	"vehicle-log-grpc-api/internal/model"
	"vehicle-log-grpc-api/internal/shared/ierr"
)

func (r repository) GetAllByTownCode(ctx context.Context, townCode string) ([]model.TownSlot, error) {
	townSlots := make([]model.TownSlot, 0)
	err := r.db.NewSelect().Model(&townSlots).
		Relation("Town", func(query *bun.SelectQuery) *bun.SelectQuery {
			return query.Where("town_code = ?", townCode)
		}).
		Relation("Type").
		Scan(ctx)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, ierr.ErrTownCodeNotHaveParkingSlot
	}

	if err != nil {
		return nil, errors.Wrap(err, "failed to get town slot")
	}
	return townSlots, nil
}
