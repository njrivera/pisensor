package models

type TempReading struct {
	Serial string  `json:"serial,omitempty"`
	Model  string  `json:"model,omitempty"`
	Temp   float64 `json:"temp,omitempty"`
	Unit   string  `json:"unit,omitempty"`
}
