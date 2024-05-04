package common

import (
	"time"
)

type SensorData struct {
	Temperature  float64   `json:"temperature"`
	Humidity     float64   `json:"humidity"`
	SoilMoisture float64   `json:"soil_moisture"`
	Timestamp    time.Time `json:"timestamp"`
}
