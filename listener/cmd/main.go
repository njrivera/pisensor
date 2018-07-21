package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Messenger/pkg/messenger"
	"github.com/go-martini/martini"
	"github.com/pisensor/pkg/models"
)

func main() {
	msnger, err := messenger.NewMessenger()
	if err != nil {
		log.Fatalf("Error getting messenger: %+v.", err)
	}

	m := martini.Classic()

	m.Post("/:serial/:model", func(req *http.Request, params martini.Params) {
		serial := params["serial"]
		model := params["model"]
		msg := models.SensorMessage{}
		if err := json.NewDecoder(req.Body).Decode(&msg); err != nil {
			log.Fatalf("Error reading temperature: %+v.", err)
		}
		jsonMsg, err := json.Marshal(models.TempReading{
			Serial: serial,
			Model:  model,
			Temp:   msg.Num,
		})
		if err != nil {
			log.Printf("Error encoding message to json: %+v.", err)
		}
		if err := msnger.Publish(jsonMsg, "sensor"); err != nil {
			log.Printf("Error publishing temp reading: %+v.", err)
		}
	})

	m.Run()
}
