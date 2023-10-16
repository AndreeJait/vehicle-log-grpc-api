package impl

import (
	"context"
	"log"
	"vehicle-log-grpc-api/internal/model"
	"vehicle-log-grpc-api/internal/usecase/logs"
)

func (u useCase) LogOutProcess(ctx context.Context, req logs.MessageLogOut) error {
	repoLog := u.repo.GetLogRepo()
	l, err := repoLog.GetLog(ctx, req.TownCode, req.PlatNumber, []string{model.LogStatusIn, model.LogStatusOut})
	if err != nil {
		return err
	}
	if l.ID == "" {
		log.Println("Vehicle not exits in town")
		return nil
	}
	err = repoLog.UpdateLog(ctx, model.Log{
		ID:     l.ID,
		Status: model.LogStatusOut,
	})
	return err
}
