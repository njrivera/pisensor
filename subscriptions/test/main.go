package main

import (
	"encoding/json"
	"log"

	"github.com/pisensor/pkg/models"

	"github.com/Messenger/pkg/factory"
)

func main() {
	msnger, err := factory.NewDefaultMessenger()
	if err != nil {
		log.Fatalf("Error starting messenger: %s.", err.Error())
	}

	msg, _ := json.Marshal(models.TempReading{Temp: 5, Serial: "yo"})

	msnger.Publish(msg, "temp_readings")
}
