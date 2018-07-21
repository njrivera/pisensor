package main

import (
	"encoding/binary"
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
		var temp float64
		if err := binary.Read(req.Body, binary.LittleEndian, &temp); err != nil {
			log.Fatalf("Error reading temperature: %+v.", err)
		}
		if err := msnger.Publish(models.TempReading{
			Serial: serial,
			Model:  model,
			Temp:   temp,
		}, "sensor"); err != nil {
			log.Printf("Error publishing temp reading: %+v.", err)
		}
	})
}
