package impl

import (
	"vehicle-log-grpc-api/internal/handlers/types"
	typesUc "vehicle-log-grpc-api/internal/usecase/types"
)

type handler struct {
	ucTypes typesUc.TypeUseCase
}

func New(ucTypes typesUc.TypeUseCase) types.Handler {
	return &handler{
		ucTypes: ucTypes,
	}
}
