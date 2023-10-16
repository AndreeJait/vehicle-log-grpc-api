package impl

import (
	"context"
	grpc "vehicle-log-grpc-api/common/grpc/proto"
	"vehicle-log-grpc-api/internal/common/converter"
)

func (h handler) GetAllTownSlotByType(ctx context.Context, slot *grpc.GetAllTownSlot) (*grpc.TownSlotList, error) {
	townSlot, err := h.townUc.GetAllTownSlotByType(ctx, slot.GetTownCode())
	if err != nil {
		return nil, err
	}
	return converter.ConvertToTownSlotListGrpc(townSlot), nil
}
