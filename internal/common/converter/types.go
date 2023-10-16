package converter

import (
	grpc "vehicle-log-grpc-api/common/grpc/proto"
	"vehicle-log-grpc-api/internal/model"
)

func ConvertFromTypeToTypeGrpcList(listFromDB []model.Type) *grpc.TypeList {
	var typesListGrpc grpc.TypeList
	listItemGrpc := make([]*grpc.Type, 0)
	for _, v := range listFromDB {
		convert := ConvertFromTypeToTypeGrpc(v)
		listItemGrpc = append(listItemGrpc, &convert)
	}
	typesListGrpc.List = listItemGrpc
	return &typesListGrpc
}

func ConvertFromTypeToTypeGrpc(typeModel model.Type) grpc.Type {
	return grpc.Type{
		Id:       typeModel.ID,
		TypeName: typeModel.Name,
	}
}
