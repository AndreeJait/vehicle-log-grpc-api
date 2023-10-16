package impl

import (
	"context"
	grpc "vehicle-log-grpc-api/common/grpc/proto"
	"vehicle-log-grpc-api/internal/common/converter"
	"vehicle-log-grpc-api/internal/usecase/logs"
)

func (h handler) LogOutVehicle(ctx context.Context, request *grpc.LogOutRequest) (*grpc.Log, error) {
	report, err := h.ucLog.LogOutDetail(ctx, logs.RequestLogOut{
		TownCode:   request.GetTownCode(),
		TimeOut:    request.GetTimeOut(),
		PlatNumber: request.PlatNumber,
	})
	if err != nil {
		return nil, err
	}
	convert := converter.ConvertFromLogToLogGrpc(report)
	return &convert, err
}
