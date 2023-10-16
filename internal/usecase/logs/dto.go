package logs

type RequestLogOut struct {
	TownCode   string
	TimeOut    string
	PlatNumber string
}

type MessageLogIn struct {
	PlatNumber  string `json:"plat_number"`
	TownCode    string `json:"town_code"`
	VehicleType string `json:"vehicle_type"`
}

type MessageLogOut struct {
	PlatNumber  string `json:"plat_number"`
	TownCode    string `json:"town_code"`
	VehicleType string `json:"vehicle_type"`
}
