package vehicle

import (
	"context"
	grpc "vehicle-log-grpc-api/common/grpc/proto"
)

type Handler interface {
	GetVehicleInTown(context.Context, *grpc.VehicleInTownRequest) (*grpc.VehicleList, error)
}
