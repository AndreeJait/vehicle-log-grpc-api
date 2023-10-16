package impl

import (
	"context"
	"github.com/pkg/errors"
	"vehicle-log-grpc-api/internal/model"
)

func (r repository) UpdateLog(ctx context.Context, log model.Log) error {

	_, err := r.db.NewUpdate().Model(&log).WherePK().OmitZero().Exec(ctx)
	if err != nil {
		return errors.Wrap(err, "failed to update log")
	}
	return nil
}
