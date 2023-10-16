package impl

import (
	"context"
	"github.com/pkg/errors"
	"github.com/uptrace/bun"
	"vehicle-log-grpc-api/internal/model"
)

func (r repository) CheckByPlatNumber(ctx context.Context, platNumber string) (bool, error) {
	var vehicle model.Vehicle
	exits, err := r.db.NewSelect().Model(&vehicle).
		Where("?=?", bun.Ident("plat_number"), platNumber).
		Exists(ctx)
	if err != nil {
		return exits, errors.Wrap(err, "failed to get data")
	}
	return exits, nil
}
