package impl

import (
	"vehicle-log-grpc-api/internal/handlers/logs"
	logsUc "vehicle-log-grpc-api/internal/usecase/logs"
)

type handler struct {
	ucLog logsUc.UseCase
}

func New(uc logsUc.UseCase) logs.Handler {
	return &handler{ucLog: uc}
}
