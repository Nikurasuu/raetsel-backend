package entity

type SolarData struct {
	PvCurrentPower float64 `json:"pv_current_power"`
	PvDailyYield   float64 `json:"pv_daily_yield"`
}
