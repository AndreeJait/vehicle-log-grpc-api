package impl

import (
	"context"
	"github.com/pkg/errors"
	"vehicle-log-grpc-api/internal/model"
)

func (r repository) GetAllVehicle(ctx context.Context) ([]model.Vehicle, error) {
	vehicles := make([]model.Vehicle, 0)

	err := r.db.NewSelect().Model(&vehicles).Scan(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get data")
	}
	return vehicles, nil
}
