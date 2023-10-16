package towns

import (
	"context"
	grpc "vehicle-log-grpc-api/common/grpc/proto"
)

type Handler interface {
	GetTownByCode(context.Context, *grpc.GetTown) (*grpc.Town, error)
	GetAllTown(context.Context, *grpc.Empty) (*grpc.TownList, error)
	GetTownSlotByType(context.Context, *grpc.GetTownSlot) (*grpc.TownSlot, error)
	GetAllTownSlotByType(context.Context, *grpc.GetAllTownSlot) (*grpc.TownSlotList, error)
}
