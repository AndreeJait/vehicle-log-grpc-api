package impl

import (
	"context"
	grpc "vehicle-log-grpc-api/common/grpc/proto"
	"vehicle-log-grpc-api/internal/common/converter"
	"vehicle-log-grpc-api/internal/usecase/vehicle"
)

func (h handlers) GetVehicleInTown(ctx context.Context, req *grpc.VehicleInTownRequest) (*grpc.VehicleList, error) {
	logs, err := h.vcUseCase.GetVehicleInTown(ctx, req.GetTownCode(), vehicle.FilterVehicleInTown{
		PlatNumber:  req.GetFilter().GetPlatNumber(),
		VehicleType: req.GetFilter().GetVehicleType(),
	})
	if err != nil {
		return nil, err
	}
	resp := converter.ConvertFromListLogTOVehicleGrpcList(logs)
	return resp, nil
}
