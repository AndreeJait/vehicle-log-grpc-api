package impl

import (
	"context"
	"vehicle-log-grpc-api/internal/model"
)

func (r repository) GetByName(ctx context.Context, name string) (model.Type, error) {
	var result model.Type
	err := r.db.NewSelect().Model(&result).Where("name=?", name).Scan(ctx)
	if err != nil {
		return result, err
	}
	return result, nil
}
