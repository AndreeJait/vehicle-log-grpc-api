package model

import (
	"github.com/uptrace/bun"
)

const (
	LogStatusIn      = "in"
	LogStatusPending = "pending"
	LogStatusOut     = "out"
)

type Log struct {
	ID          string   `json:"id" bun:",pk"`
	TownSlotsID string   `json:"town_slots_id"`
	TownSlot    TownSlot `json:"town_slot" bun:"rel:belongs-to,join:town_slots_id=id"`
	VehicleID   string   `json:"vehicle_id"`
	Vehicle     Vehicle  `json:"vehicle" bun:"rel:belongs-to,join:vehicle_id=id"`
	TimeIn      string   `json:"time_in"`
	TimeOut     *string  `json:"time_out"`
	DateAt      string   `json:"date_at"`
	DateOutAt   *string  `json:"date_out_at"`
	Status      string   `json:"status"`
}

type LogReport struct {
	bun.BaseModel `bun:"table:logs,alias:u"`
	VehicleType   string `json:"vehicle_type"`
	Quantity      int    `json:"quantity"`
	TownCode      string `json:"town_code"`
	TotalLogged   int    `json:"total_logged"`
}
