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
	serial = flag.String("serial", "", "Serial number for pi sensor")
	addr   = "http://localhost:5558"
)

func main() {
	flag.Parse()

	client := &http.Client{}

	for {
		time.Sleep(2 * time.Second)

		temp := float64(rand.Intn(5)) + rand.Float64()
		reading := models.SensorMessage{Num: temp}

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
