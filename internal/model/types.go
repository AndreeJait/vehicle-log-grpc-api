package model

type Type struct {
	ID   string `json:"id" bun:",pk"`
	Name string `json:"name"`
}
