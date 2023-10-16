package model

type Vehicle struct {
	ID         string `json:"id" bun:",pk"`
	PlatNumber string `json:"plat_number"`
	TypeID     string `json:"type_id"`
	Type       Type   `json:"type" bun:"rel:belongs-to,join:type_id=id"`
}
