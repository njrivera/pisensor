package models

import (
	"time"
)

type TempReading struct {
	Serial    string    `json:"serial,omitempty"`
	Model     string    `json:"model,omitempty"`
	Temp      float64   `json:"temp,omitempty"`
	Unit      string    `json:"unit,omitempty"`
	Timestamp time.Time `json:"time,omitempty"`
}
