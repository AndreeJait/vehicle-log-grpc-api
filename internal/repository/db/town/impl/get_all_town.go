package impl

import (
	"context"
	"github.com/pkg/errors"
	"vehicle-log-grpc-api/internal/model"
)

func (r *repository) GetAllTown(ctx context.Context) ([]model.Town, error) {
	towns := make([]model.Town, 0)
	err := r.db.NewSelect().Model(&towns).Scan(ctx)
	if err != nil {
		return towns, errors.Wrap(err, "failed to get all town")
	}
	return towns, nil
}
