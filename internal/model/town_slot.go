package model

type TownSlot struct {
	ID       string `json:"id" bun:",pk"`
	TownID   string `json:"town_id"`
	Town     Town   `json:"town" bun:"rel:belongs-to,join:town_id=id"`
	TypeID   string `json:"type_id"`
	Type     Type   `json:"type" bun:"rel:belongs-to,join:type_id=id"`
	Quantity int    `json:"quantity"`
	Price    int    `json:"price"`
}
