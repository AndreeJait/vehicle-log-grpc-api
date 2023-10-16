package impl

import (
	"vehicle-log-grpc-api/internal/handlers/nsqlogout"
	"vehicle-log-grpc-api/internal/usecase/logs"
)

type handler struct {
	logUc logs.UseCase
}

func New(logUc logs.UseCase) nsqlogout.Handler {

	return &handler{
		logUc: logUc,
	}
}
