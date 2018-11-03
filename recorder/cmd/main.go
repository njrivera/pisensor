package main

import (
	"encoding/json"
	"log"

	"github.com/Messenger/pkg/factory"
	"github.com/pisensor/pkg/models"
	"github.com/pisensor/recorder/internal/dbfactory"
)

const (
	msgChannel = "temp_readings"
	dbType     = dbfactory.SQLite
)

func main() {
	msnger, err := factory.NewDefaultMessenger()
	if err != nil {
		log.Fatalf("Error starting messenger: %s.", err.Error())
	}

	msgChan, err := msnger.Subscribe(msgChannel)
	if err != nil {
		log.Fatalf("Error subscribing to %s channel: %s", msgChannel, err.Error())
	}

	db, err := dbfactory.NewDB(dbType)
	if err != nil {
		log.Fatalf("Error initializing database connection: %s", err.Error())
	}

	log.Printf("Running recorder service...")
	defer log.Printf("Stopped recorder service")

	for msg := range msgChan {
		reading := models.TempReading{}
		if err := json.Unmarshal([]byte(msg.Payload), &reading); err != nil {
			log.Printf("Error unmarshalling message to reading: %s.", err.Error())
			continue
		}

		if db.InsertReading(reading); err != nil {
			log.Printf("Error inserting reading: %s", err.Error())
		}
	}
}
