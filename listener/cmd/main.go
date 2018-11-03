package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	messengerfactory "github.com/Messenger/pkg/factory"
	"github.com/go-martini/martini"
	"github.com/pisensor/pkg/models"
)

const (
	msgChannel = "temp_readings"
)

func main() {
	msnger, err := messengerfactory.NewDefaultMessenger()
	if err != nil {
		log.Fatalf("Error getting messenger: %s", err.Error())
	}

	m := martini.Classic()

	m.Post("/:serial/:model", func(req *http.Request, params martini.Params) {
		serial := params["serial"]
		model := params["model"]

		msg := models.SensorMessage{}
		if err := json.NewDecoder(req.Body).Decode(&msg); err != nil {
			log.Printf("Error reading temperature: %+s", err.Error())
			return
		}

		jsonMsg, err := json.Marshal(models.TempReading{
			Serial:    serial,
			Model:     model,
			Temp:      msg.Num,
			Unit:      "Farenheit",
			Timestamp: time.Now(),
		})
		if err != nil {
			log.Printf("Error encoding message to json: %s", err.Error())
			return
		}

		if err := msnger.Publish(jsonMsg, msgChannel); err != nil {
			log.Printf("Error publishing temp reading: %s", err.Error())
		}
	})

	log.Printf("Running listener service...")

	m.Run()
}
