package impl

import (
	"vehicle-log-grpc-api/internal/handlers/nsqlogin"
	"vehicle-log-grpc-api/internal/usecase/logs"
)

type handler struct {
	logUc logs.UseCase
}

func New(logUC logs.UseCase) nsqlogin.Handler {

	return &handler{
		logUc: logUC,
	}
}
