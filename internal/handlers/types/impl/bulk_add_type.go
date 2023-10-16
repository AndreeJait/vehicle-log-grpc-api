package impl

import (
	"context"
	grpc "vehicle-log-grpc-api/common/grpc/proto"
)

func (h handler) BulkAddType(ctx context.Context, list *grpc.TypesAddRequestList) (*grpc.EmptyType, error) {
	var listStr = make([]string, 0)
	for _, v := range list.Towns {
		listStr = append(listStr, v.Name)
	}
	err := h.ucTypes.BulkInsertType(ctx, listStr)
	if err != nil {
		return nil, err
	}
	return &grpc.EmptyType{}, nil
}
