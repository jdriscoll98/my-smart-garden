package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"my-smart-garden/common"
	"net/http"
	"time"
)


func main() {
	// Initial conditions with reasonable starting points
	temperature := 25.0  // Average temperature starts at 25°C
	humidity := 50.0     // Average humidity starts at 50%
	soilMoisture := 30.0 // Average soil moisture starts at 30%

	for {
		// Simulate sensor data generation every 10 seconds
		time.Sleep(10 * time.Second)

		// Small random change for each sensor reading
		temperature += (rand.Float64() - 0.5) * 0.5 // Fluctuate by up to ±0.25°C
		humidity += (rand.Float64() - 0.5) * 2.0    // Fluctuate by up to ±1%
		soilMoisture += (rand.Float64() - 0.5) * 1.0 // Fluctuate by up to ±0.5%

		// Ensuring values stay within reasonable ranges
		temperature = clamp(temperature, 20, 35)
		humidity = clamp(humidity, 30, 100)
		soilMoisture = clamp(soilMoisture, 10, 50)

		data := common.SensorData{
			Temperature:  temperature,
			Humidity:     humidity,
			SoilMoisture: soilMoisture,
			Timestamp:    time.Now(),
		}

		sendData(data)
	}
}

func clamp(value, min, max float64) float64 {
	if value < min {
		return min
	} else if value > max {
		return max
	}
	return value
}

func sendData(data common.SensorData) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}

	// console.log the data
	resp, err := http.Post("http://localhost:8080/data", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Failed to send data:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Data sent successfully:", data)
}
