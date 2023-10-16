package impl

import (
	"context"
	"database/sql"
	"github.com/pkg/errors"
	"github.com/uptrace/bun"
	"vehicle-log-grpc-api/internal/model"
)

func (r repository) GetAvailableCount(ctx context.Context, townCode string, date string) (int, error) {
	count, err := r.db.NewSelect().
		Model((*model.LogReport)(nil)).
		ColumnExpr("tp.name as vehicle_type", "ts.quantity as quantity", "t.town_code as town_code").
		Join("INNER JOIN town_slots as ts ON ts.id = u.town_slots_id").
		Join("INNER JOIN towns as t ON t.id = ts.town_id").
		Join("INNER JOIN types as tp ON tp.id = ts.type_id ").
		Where("(date_at = ? OR status NOT IN (?)) AND t.town_code = ? ", date,
			bun.In([]string{model.LogStatusOut, model.LogStatusPending}), townCode).
		Group("vehicle_type", "town_code", "quantity").
		Count(ctx)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, nil
		}
		return 0, errors.Wrap(err, "failed to get report")
	}
	return count, nil
}
