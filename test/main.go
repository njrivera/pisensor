package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/pisensor/pkg/models"
)

var (
	numReadings = flag.Int("numReadings", 1000, "Number of readings to generate")
	serial      = flag.String("serial", "", "Serial number for pi sensor")
	addr        = "http://localhost:5558"
)

func main() {
	flag.Parse()

	client := &http.Client{}

	for _, reading := range generateReadings(*numReadings) {
		time.Sleep(2 * time.Second)
		log.Printf("Sending temp %f for serial %s", reading.Num, *serial)

		msg, _ := json.Marshal(reading)

		url := addr + "/" + *serial + "/MODEL-B"

		r, _ := http.NewRequest("POST", url, bytes.NewBuffer(msg))
		r.Header.Set("Content-Type", "application/json")

		resp, err := client.Do(r)
		if err != nil {
			log.Printf("Error sending reading to listener: %s", err.Error())
		} else {
			resp.Body.Close()
		}
	}
}

func generateReadings(n int) []models.SensorMessage {
	readings := []models.SensorMessage{}

	for i := 0; i < n; i++ {
		temp := float64(rand.Intn(200)) + rand.Float64()
		readings = append(readings, models.SensorMessage{Num: temp})
	}

	return readings
}
