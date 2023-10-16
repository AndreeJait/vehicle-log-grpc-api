package converter

import (
	grpc "vehicle-log-grpc-api/common/grpc/proto"
	"vehicle-log-grpc-api/internal/common/utils"
	"vehicle-log-grpc-api/internal/model"
)

func ConvertFromLogReportToLogReportGrpcList(logs []model.LogReport, date string) *grpc.TownLogReportList {
	var listsGrpc grpc.TownLogReportList
	listsGrpc.Date = date
	listGrpcItem := make([]*grpc.TownLogReport, 0)

	for _, v := range logs {
		convert := ConvertFromLogReportToLogReportGrpc(v)
		listGrpcItem = append(listGrpcItem, &convert)
	}

	listsGrpc.List = listGrpcItem
	return &listsGrpc
}

func ConvertFromLogReportToLogReportGrpc(log model.LogReport) grpc.TownLogReport {
	return grpc.TownLogReport{
		TotalLogged: int64(log.TotalLogged),
		TownCode:    log.TownCode,
		Quantity:    int64(log.Quantity),
		VehicleType: log.VehicleType,
	}
}

func ConvertFromLogToLogGrpc(log model.Log) grpc.Log {
	convertedTimeIn := utils.ConvertToTime(log.DateAt, log.TimeIn)
	convertedTimeOut := utils.ConvertToTime(utils.PtrStr(log.DateOutAt), utils.PtrStr(log.TimeOut))
	totalDif := convertedTimeOut.Sub(convertedTimeIn)
	return grpc.Log{
		TownName:    log.TownSlot.Town.TownName,
		TownCode:    log.TownSlot.Town.TownCode,
		TimeOut:     utils.PtrStr(log.TimeOut),
		TimeIn:      log.TimeIn,
		PlatNumber:  log.Vehicle.PlatNumber,
		DateAt:      log.DateAt,
		DateOutAt:   utils.PtrStr(log.DateOutAt),
		VehicleType: log.TownSlot.Type.Name,
		TotalTime:   float32(totalDif.Seconds()),
	}
}
