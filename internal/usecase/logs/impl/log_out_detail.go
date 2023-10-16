package impl

import (
	"context"
	"vehicle-log-grpc-api/internal/common/utils"
	"vehicle-log-grpc-api/internal/model"
	"vehicle-log-grpc-api/internal/shared/ierr"
	"vehicle-log-grpc-api/internal/usecase/logs"
)

func (u useCase) LogOutDetail(ctx context.Context, req logs.RequestLogOut) (model.Log, error) {
	logRepo := u.repo.GetLogRepo()
	foundLog, err := logRepo.GetLog(ctx, req.TownCode, req.PlatNumber, []string{model.LogStatusPending, model.LogStatusOut})
	if err != nil {
		return foundLog, err
	}
	if foundLog.ID == "" {
		return foundLog, ierr.ErrVehicleNotFoundInTown
	}
	timeNow := utils.NowJkrt().Format(utils.DefaultFormatDate)
	foundLog.TimeOut = &req.TimeOut
	foundLog.DateOutAt = &timeNow
	foundLog.Status = model.LogStatusPending
	err = logRepo.UpdateLog(ctx, foundLog)
	_, data, err := u.redis.CheckAndGetTownAvailable(ctx, req.TownCode, foundLog.TownSlot.Type.Name)
	if err != nil {
		return foundLog, err
	}
	data -= 1
	err = u.redis.SetTownAvailable(ctx, req.TownCode, foundLog.TownSlot.Type.Name, data)
	return foundLog, err
}
