package impl

import (
	"context"
	"github.com/pkg/errors"
	"github.com/uptrace/bun"
	"vehicle-log-grpc-api/internal/model"
)

func (r repository) GetReportedLog(ctx context.Context, townCode string, date string) ([]model.LogReport, error) {
	results := make([]model.LogReport, 0)
	err := r.db.NewSelect().
		Model(&results).
		ColumnExpr("tp.name as vehicle_type", "ts.quantity as quantity", "t.town_code as town_code").
		ColumnExpr("COUNT(*) as total_logged").
		Join("INNER JOIN town_slots as ts ON ts.id = u.town_slots_id").
		Join("INNER JOIN towns as t ON t.id = ts.town_id").
		Join("INNER JOIN types as tp ON tp.id = ts.type_id ").
		Where("(status NOT IN (?)) AND t.town_code = ? ",
			bun.In([]string{model.LogStatusOut, model.LogStatusPending}), townCode).
		Group("vehicle_type", "town_code", "quantity").
		Scan(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get report")
	}
	return results, nil
}
