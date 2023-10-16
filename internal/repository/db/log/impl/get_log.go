package impl

import (
	"context"
	"database/sql"
	"github.com/pkg/errors"
	"github.com/uptrace/bun"
	"vehicle-log-grpc-api/internal/model"
)

func (r repository) GetLog(ctx context.Context, townCode string, platNumber string, status []string) (model.Log, error) {
	var log model.Log
	err := r.db.NewSelect().
		Model(&log).
		Relation("TownSlot").
		Relation("TownSlot.Town", func(query *bun.SelectQuery) *bun.SelectQuery {
			return query.Where("town_code = ?", townCode)
		}).
		Relation("TownSlot.Type").
		Relation("Vehicle", func(query *bun.SelectQuery) *bun.SelectQuery {
			return query.Where("plat_number = ?", platNumber)
		}).
		Where("status not in (?)", bun.In(status)).
		Scan(ctx)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.Log{}, nil
		} else {
			return log, errors.Wrap(err, "failed to fetch data")
		}
	}
	return log, nil
}
