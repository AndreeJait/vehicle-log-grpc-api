package impl

import (
	"context"
	grpc "vehicle-log-grpc-api/common/grpc/proto"
	"vehicle-log-grpc-api/internal/common/converter"
)

func (h handler) GetTownSlotByType(ctx context.Context, slot *grpc.GetTownSlot) (*grpc.TownSlot, error) {
	townSlot, err := h.townUc.GetTownSlotByType(ctx, slot.GetTownCode(), slot.GetType())
	if err != nil {
		return nil, err
	}
	converted := converter.ConvertToTownSlotGrpc(townSlot)
	return &converted, err
}
