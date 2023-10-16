package impl

import (
	"context"
	"vehicle-log-grpc-api/internal/model"
)

func (u useCase) GetTownByCode(ctx context.Context, townCode string) (model.Town, error) {
	townFromRepo, err := u.repo.GetTownRepo().GetTownByCode(ctx, townCode)
	if err != nil {
		return townFromRepo, err
	}
	return townFromRepo, nil
}
