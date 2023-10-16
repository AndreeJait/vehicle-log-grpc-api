package impl

import (
	"context"
	grpc "vehicle-log-grpc-api/common/grpc/proto"
	"vehicle-log-grpc-api/internal/common/converter"
)

func (h handler) GetTownByCode(ctx context.Context, town *grpc.GetTown) (*grpc.Town, error) {
	t, err := h.townUc.GetTownByCode(ctx, town.GetTownCode())
	if err != nil {
		return nil, err
	}
	townConvert := converter.ConvertToTownGrpc(t)
	return &townConvert, nil
}
