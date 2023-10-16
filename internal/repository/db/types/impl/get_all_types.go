package impl

import (
	"context"
	"github.com/pkg/errors"
	"vehicle-log-grpc-api/internal/model"
)

func (r repository) GetAllTypes(ctx context.Context) ([]model.Type, error) {
	results := make([]model.Type, 0)
	err := r.db.NewSelect().Model(&results).Scan(ctx)
	if err != nil {
		return results, errors.Wrap(err, "failed to get all types")
	}
	return results, nil
}
