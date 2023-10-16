package impl

import (
	"context"
	"vehicle-log-grpc-api/internal/model"
)

func (u useCase) ReportLogged(ctx context.Context, date string, townCode string) ([]model.LogReport, error) {
	logsTemp := make([]model.LogReport, 0)
	types, err := u.repo.GetTownSlotRepository().GetAllByTownCode(ctx, townCode)
	if err != nil {
		return logsTemp, err
	}
	logs, err := u.repo.GetLogRepo().GetReportedLog(ctx, townCode, date)
	if err != nil {
		return logs, err
	}
	for _, v := range types {
		found := u.findIndex(v.Type.Name, logs)
		totalLogged := 0
		if found != -1 {
			totalLogged = logs[found].TotalLogged
		}
		logsTemp = append(logsTemp, model.LogReport{
			TownCode:    townCode,
			VehicleType: v.Type.Name,
			Quantity:    v.Quantity,
			TotalLogged: totalLogged,
		})
	}
	return logsTemp, nil
}

func (u useCase) findIndex(vehicleType string, logs []model.LogReport) int {
	for i, v := range logs {
		if v.VehicleType == vehicleType {
			return i
		}
	}
	return -1
}
