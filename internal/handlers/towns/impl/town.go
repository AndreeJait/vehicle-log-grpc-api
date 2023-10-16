package impl

import (
	"vehicle-log-grpc-api/internal/handlers/towns"
	townsUC "vehicle-log-grpc-api/internal/usecase/towns"
)

type handler struct {
	townUc townsUC.TownUseCase
}

func New(tc townsUC.TownUseCase) towns.Handler {
	return &handler{
		townUc: tc,
	}
}
