package impl

import (
	"context"
	grpc "vehicle-log-grpc-api/common/grpc/proto"
	"vehicle-log-grpc-api/internal/common/converter"
)

func (h handler) GetAllType(ctx context.Context, emptyType *grpc.EmptyType) (*grpc.TypeList, error) {
	types, err := h.ucTypes.GetAllTypes(ctx)
	if err != nil {
		return nil, err
	}
	return converter.ConvertFromTypeToTypeGrpcList(types), nil
}
