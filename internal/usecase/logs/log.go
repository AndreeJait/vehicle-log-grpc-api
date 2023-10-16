package logs

import (
	"context"
	"vehicle-log-grpc-api/internal/model"
)

type UseCase interface {
	LogOutDetail(ctx context.Context, req RequestLogOut) (model.Log, error)
	ReportLogged(ctx context.Context, date string, townCode string) ([]model.LogReport, error)

	LogOutProcess(ctx context.Context, req MessageLogOut) error
	LogInProcess(ctx context.Context, req MessageLogIn) error
}
