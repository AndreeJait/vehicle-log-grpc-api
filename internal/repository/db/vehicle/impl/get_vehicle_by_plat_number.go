package impl

import (
	"context"
	"github.com/pkg/errors"
	"github.com/uptrace/bun"
	"vehicle-log-grpc-api/internal/model"
)

func (r repository) GetVehicleByPlatNumber(ctx context.Context, platNumber string) (model.Vehicle, error) {
	var vehicle model.Vehicle
	err := r.db.NewSelect().Model(&vehicle).
		Relation("Type").
		Where("?=?", bun.Ident("plat_number"), platNumber).
		Scan(ctx)
	if err != nil {
		return vehicle, errors.Wrap(err, "failed to get data")
	}
	return vehicle, nil
}
