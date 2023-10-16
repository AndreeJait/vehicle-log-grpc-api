package model

type Town struct {
	ID       string `json:"id" bun:",pk"`
	TownCode string `json:"town_code"`
	TownName string `json:"town_name"`
}
