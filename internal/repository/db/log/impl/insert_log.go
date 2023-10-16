package impl

import (
	"context"
	"github.com/pkg/errors"
	"vehicle-log-grpc-api/internal/model"
)

func (r repository) InsertLog(ctx context.Context, log model.Log) error {
	_, err := r.db.NewInsert().Model(&log).Exec(ctx)
	if err != nil {
		return errors.Wrap(err, "failed to insert log")
	}
	return nil
}
