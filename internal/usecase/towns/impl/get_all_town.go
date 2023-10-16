package impl

import (
	"context"
	"encoding/json"
	"log"
	"vehicle-log-grpc-api/internal/model"
)

func (u useCase) GetAllTown(ctx context.Context) ([]model.Town, error) {
	result := make([]model.Town, 0)
	check, redis, err := u.redis.CheckAndGetTownList(ctx)
	if err != nil {
		return result, err
	}
	if check {
		err = json.Unmarshal([]byte(redis), &result)
		if err != nil {
			return result, err
		}
		return result, nil
	}
	townFromRepo, err := u.repo.GetTownRepo().GetAllTown(ctx)
	if err != nil {
		return result, err
	}
	jsonStrByte, err := json.Marshal(&townFromRepo)
	if err != nil {
		return townFromRepo, err
	}
	go u.makeCacheForQuantity(townFromRepo)
	err = u.redis.SetListTown(ctx, string(jsonStrByte))
	return townFromRepo, err
}

func (u useCase) makeCacheForQuantity(towns []model.Town) {
	ctx := context.Background()
	repoTownSlot := u.repo.GetTownSlotRepository()
	for _, t := range towns {
		types, err := repoTownSlot.GetAllByTownCode(ctx, t.TownCode)
		if err != nil {
			log.Println(err)
			return
		}
		for _, v := range types {
			check, _, err := u.redis.CheckAndGetTownQuantity(ctx, t.TownCode, v.Type.Name)
			if err != nil {
				log.Println(err)
				return
			}
			if !check {
				err = u.redis.SetTownQuantity(ctx, t.TownCode, v.Type.Name, v.Quantity)
			}
			continue
		}
	}
}
