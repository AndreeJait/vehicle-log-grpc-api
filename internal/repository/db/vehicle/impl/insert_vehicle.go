package impl

import (
	"context"
	"github.com/pkg/errors"
	"vehicle-log-grpc-api/internal/model"
)

func (r repository) InsertVehicle(ctx context.Context, vehicle model.Vehicle) error {

	_, err := r.db.NewInsert().Model(&vehicle).Exec(ctx)
	if err != nil {
		return errors.Wrap(err, "failed to insert vehicle")
	}
	return nil
}
