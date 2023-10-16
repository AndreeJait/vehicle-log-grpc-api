package types

import (
	"context"
	grpc "vehicle-log-grpc-api/common/grpc/proto"
)

type Handler interface {
	GetAllType(context.Context, *grpc.EmptyType) (*grpc.TypeList, error)
	BulkAddType(context.Context, *grpc.TypesAddRequestList) (*grpc.EmptyType, error)
}
