package converter

import (
	grpc "vehicle-log-grpc-api/common/grpc/proto"
	"vehicle-log-grpc-api/internal/model"
)

func ConvertToListVehicleGrpc(vehicles []model.Vehicle) *grpc.VehicleList {
	var vehicleListGrpc grpc.VehicleList
	vehicleGrpcSlice := make([]*grpc.Vehicle, 0)
	for _, v := range vehicles {
		convert := ConvertToVehicleGrpc(v)
		vehicleGrpcSlice = append(vehicleGrpcSlice, &convert)
	}
	vehicleListGrpc.List = vehicleGrpcSlice
	return &vehicleListGrpc
}

func ConvertToVehicleGrpc(v model.Vehicle) grpc.Vehicle {
	return grpc.Vehicle{
		Id:         v.ID,
		PlatNumber: v.PlatNumber,
	}
}

func ConvertFromListLogTOVehicleGrpcList(logs []model.Log) *grpc.VehicleList {
	var vehicleList grpc.VehicleList
	vehicleConvertList := make([]*grpc.Vehicle, 0)
	for _, v := range logs {
		convert := ConvertFromLogToVehicleGrpc(v)
		vehicleConvertList = append(vehicleConvertList, &convert)
	}
	vehicleList.List = vehicleConvertList
	return &vehicleList
}

func ConvertFromLogToVehicleGrpc(v model.Log) grpc.Vehicle {
	timeOut := "-"
	dateOut := "-"
	if v.TimeOut != nil {
		timeOut = *v.TimeOut
		dateOut = *v.DateOutAt
	}
	return grpc.Vehicle{
		Id:           v.VehicleID,
		PlatNumber:   v.Vehicle.PlatNumber,
		VehicleType:  v.TownSlot.Type.Name,
		TimeCheckIn:  v.TimeIn,
		TimeCheckOut: timeOut,
		DateCheckIn:  v.DateAt,
		DateCheckOut: dateOut,
	}
}
