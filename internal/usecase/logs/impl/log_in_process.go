package impl

import (
	"context"
	"github.com/google/uuid"
	"log"
	"vehicle-log-grpc-api/internal/common/utils"
	"vehicle-log-grpc-api/internal/model"
	"vehicle-log-grpc-api/internal/repository/db"
	"vehicle-log-grpc-api/internal/usecase/logs"
)

func (u useCase) LogInProcess(ctx context.Context, req logs.MessageLogIn) error {
	log.Println("Receive request login")
	_, err := u.repo.DoInTransaction(ctx, func(ctx context.Context, repo db.Repository) (interface{}, error) {
		timeNow := utils.NowJkrt()
		repoVehicle := repo.GetVehicleRepo()
		repoTownSlot := repo.GetTownSlotRepository()
		repoLog := repo.GetLogRepo()
		repoType := repo.GetTypeRepository()

		check, err := repoVehicle.CheckByPlatNumber(ctx, req.PlatNumber)
		if err != nil {
			return nil, err
		}
		if !check {
			types, err := repoType.GetByName(ctx, req.VehicleType)
			if err != nil {
				return nil, err
			}
			err = repoVehicle.InsertVehicle(ctx, model.Vehicle{
				ID:         uuid.NewString(),
				PlatNumber: req.PlatNumber,
				TypeID:     types.ID,
			})
			if err != nil {
				return nil, err
			}
		}

		vehicle, err := repoVehicle.GetVehicleByPlatNumber(ctx, req.PlatNumber)
		if err != nil {
			return nil, err
		}

		townSlot, err := repoTownSlot.GetByTownCodeAndVehicleType(ctx, req.TownCode, vehicle.Type.Name)
		if err != nil {
			return nil, err
		}

		getLog, err := repoLog.GetLog(ctx, req.TownCode, req.PlatNumber, []string{model.LogStatusPending, model.LogStatusOut})
		if err != nil {
			return nil, err
		}

		if getLog.ID != "" {
			log.Println("Vehicle already in town")
			return nil, nil
		}

		l := model.Log{
			ID:          uuid.NewString(),
			TimeIn:      timeNow.Format("15:04"),
			DateAt:      timeNow.Format(utils.DefaultFormatDate),
			Status:      model.LogStatusIn,
			TownSlotsID: townSlot.ID,
			VehicleID:   vehicle.ID,
		}
		err = repoLog.InsertLog(ctx, l)
		if err != nil {
			return nil, err
		}
		_, data, err := u.redis.CheckAndGetTownAvailable(ctx, req.TownCode, townSlot.Type.Name)
		if err != nil {
			return nil, err
		}
		if data == 0 {
			count, err := repoLog.GetAvailableCount(ctx, req.TownCode, timeNow.Format(utils.DefaultFormatDate))
			if err != nil {
				return nil, err
			}
			data += count
		} else {
			data += 1
		}
		err = u.redis.SetTownAvailable(ctx, req.TownCode, townSlot.Type.Name, data)
		return nil, err
	})
	return err
}
