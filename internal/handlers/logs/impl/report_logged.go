package impl

import (
	"context"
	grpc "vehicle-log-grpc-api/common/grpc/proto"
	"vehicle-log-grpc-api/internal/common/converter"
)

func (h handler) ReportLogged(ctx context.Context, report *grpc.GetLogReport) (*grpc.TownLogReportList, error) {
	reports, err := h.ucLog.ReportLogged(ctx, report.GetDate(), report.GetTownCode())
	if err != nil {
		return nil, err
	}

	return converter.ConvertFromLogReportToLogReportGrpcList(reports, report.GetDate()), nil
}
