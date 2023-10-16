package logs

import (
	"context"
	grpc "vehicle-log-grpc-api/common/grpc/proto"
)

type Handler interface {
	LogOutVehicle(context.Context, *grpc.LogOutRequest) (*grpc.Log, error)
	ReportLogged(context.Context, *grpc.GetLogReport) (*grpc.TownLogReportList, error)
}
