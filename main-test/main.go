package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

// Define a Go struct AirQualityReading to represent sensor readings:
type AirQualityReading struct {
	SensorID  string    `json:"sensor_id"`
	Timestamp time.Time `json:"timestamp"`
	PM25      float64   `json:"pm25"`
	CO2       float64   `json:"co2"`
}

// parseReadings parses a byte array of JSON data into a slice of AirQualityReading.
func parseReadings(data []byte) ([]AirQualityReading, error) {
	var readings []AirQualityReading
	err := json.Unmarshal(data, &readings)
	if err != nil {
		return nil, fmt.Errorf("error parsing JSON: %v", err)
	}
	return readings, nil
}

// calculate and return the average value of each pollutant.
func calculateAverage(readings []AirQualityReading) map[string]float64 {
	var totalPM25, totalCO2 float64
	for _, reading := range readings {
		totalPM25 += reading.PM25
		totalCO2 += reading.CO2
	}
	count := float64(len(readings))
	averages := map[string]float64{
		"pm25": totalPM25 / count,
		"co2":  totalCO2 / count,
	}
	return averages
}

func main() {
	// Example JSON data
	jsonData := `[
        {"sensor_id": "S001", "timestamp": "2023-12-28T10:00:00Z", "pm25": 25.5, "co2": 410.2},
		{"sensor_id": "S002", "timestamp": "2023-12-28T10:05:00Z", "pm25": 30.8, "co2": 405.7},
		{"sensor_id": "S001", "timestamp": "2023-12-28T11:00:00Z", "pm25": 18.2, "co2": 395.1}
    ]`

	readings, err := parseReadings([]byte(jsonData))
	if err != nil {
		log.Fatalf("Failed to parse readings: %v", err)
	}

	averages := calculateAverage(readings)
	fmt.Printf("Average PM2.5: %.2f\n", averages["pm25"])
	fmt.Printf("Average CO2: %.2f\n", averages["co2"])
}
