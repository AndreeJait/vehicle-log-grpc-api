package log

import (
	"context"
	"vehicle-log-grpc-api/internal/model"
	"vehicle-log-grpc-api/internal/usecase/vehicle"
)

type LogsRepository interface {
	GetVehicleInTown(ctx context.Context, towerCode string, filter vehicle.FilterVehicleInTown) ([]model.Log, error)
	GetReportedLog(ctx context.Context, townCode string, date string) ([]model.LogReport, error)
	GetLog(ctx context.Context, townCode string, platNumber string, status []string) (model.Log, error)
	UpdateLog(ctx context.Context, log model.Log) error
	InsertLog(ctx context.Context, log model.Log) error

	GetAvailableCount(ctx context.Context, towerCode string, date string) (int, error)
}
