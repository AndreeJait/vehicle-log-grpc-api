package converter

import (
	grpc "vehicle-log-grpc-api/common/grpc/proto"
	"vehicle-log-grpc-api/internal/model"
)

func ConvertToListTownGrpc(towns []model.Town) *grpc.TownList {
	var townList grpc.TownList
	townsGrpc := make([]*grpc.Town, 0)
	for _, v := range towns {
		convert := ConvertToTownGrpc(v)
		townsGrpc = append(townsGrpc, &convert)
	}
	townList.Lists = townsGrpc
	return &townList
}

func ConvertToTownGrpc(t model.Town) grpc.Town {
	return grpc.Town{
		Id:       t.ID,
		TownCode: t.TownCode,
		TownName: t.TownName,
	}
}
