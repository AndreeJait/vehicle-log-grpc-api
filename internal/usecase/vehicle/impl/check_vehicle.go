package impl

import (
	"context"
)

func (uc useCase) CheckVehicle(ctx context.Context, platNumber string) (bool, error) {
	return uc.repo.GetVehicleRepo().CheckByPlatNumber(ctx, platNumber)
}
