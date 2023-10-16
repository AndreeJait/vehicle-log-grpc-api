package impl

import (
	"context"
	"github.com/pkg/errors"
	"vehicle-log-grpc-api/internal/model"
)

func (r repository) BulkInsert(ctx context.Context, types []model.Type) error {
	_, err := r.db.NewInsert().Model(&types).
		On("CONFLICT (name) DO NOTHING").
		Exec(ctx)
	if err != nil {
		return errors.Wrap(err, "failed to bulk insert")
	}
	return nil
}
