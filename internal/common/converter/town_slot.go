package converter

import (
	grpc "vehicle-log-grpc-api/common/grpc/proto"
	"vehicle-log-grpc-api/internal/model"
)

func ConvertToTownSlotGrpc(townSlot model.TownSlot) grpc.TownSlot {
	return grpc.TownSlot{
		TownCode:   townSlot.Town.TownCode,
		TownName:   townSlot.Town.TownName,
		TownSlot:   int64(townSlot.Quantity),
		TypeName:   townSlot.Type.Name,
		TypeId:     townSlot.TypeID,
		TownSlotId: townSlot.ID,
	}
}

func ConvertToTownSlotListGrpc(logs []model.TownSlot) *grpc.TownSlotList {
	var townList grpc.TownSlotList
	var townItemList = make([]*grpc.TownSlot, 0)
	for _, v := range logs {
		convert := ConvertToTownSlotGrpc(v)
		townItemList = append(townItemList, &convert)
	}
	townList.List = townItemList
	return &townList
}
