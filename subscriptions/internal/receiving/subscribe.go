package receiving

import (
	"encoding/json"
	"log"

	messengertypes "github.com/Messenger/pkg/types"
	"github.com/pisensor/pkg/models"
)

func GoSubscribe(msnger messengertypes.Messenger, msgChannel string) <-chan models.TempReading {
	msgChan, err := msnger.Subscribe(msgChannel)
	if err != nil {
		log.Fatalf("Error subscribing to sensor channel: %s.", err.Error())
	}

	readingsChan := make(chan models.TempReading)

	go func() {
		for {
			reading := models.TempReading{}
			msg := <-msgChan
			if err := json.Unmarshal([]byte(msg.Payload), &reading); err != nil {
				log.Printf("Error unmarshalling message to reading: %s.", err.Error())
				continue
			}

			readingsChan <- reading
		}
	}()

	return readingsChan
}
