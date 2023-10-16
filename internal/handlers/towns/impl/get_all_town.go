package impl

import (
	"context"
	grpc "vehicle-log-grpc-api/common/grpc/proto"
	"vehicle-log-grpc-api/internal/common/converter"
)

func (h handler) GetAllTown(ctx context.Context, _ *grpc.Empty) (*grpc.TownList, error) {
	towns, err := h.townUc.GetAllTown(ctx)
	if err != nil {
		return nil, err
	}
	return converter.ConvertToListTownGrpc(towns), nil
}
