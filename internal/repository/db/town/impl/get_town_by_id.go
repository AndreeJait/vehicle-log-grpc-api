package impl

import (
	"context"
	"github.com/pkg/errors"
	"github.com/uptrace/bun"
	"vehicle-log-grpc-api/internal/model"
)

func (r *repository) GetTownByID(ctx context.Context, townID string) (model.Town, error) {
	var town model.Town
	err := r.db.NewSelect().Model(&town).Where("?=?", bun.Ident("id"), townID).Scan(ctx)
	if err != nil {
		return town, errors.Wrap(err, "failed to get town")
	}
	return town, nil
}
