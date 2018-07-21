package main

import (
	"encoding/json"
	"log"
	"sync"

	"github.com/Websocket/pkg/websockettypes"

	"github.com/Messenger/pkg/messenger"
	"github.com/Messenger/pkg/messengertypes"
	"github.com/Websocket/pkg/websocket"
	"github.com/pisensor/pkg/models"
)

var (
	clients  = newClientMap()
	wsServer websockettypes.Server
)

type clientMap struct {
	clients map[int]struct{}
	sync.RWMutex
}

func newClientMap() *clientMap {
	return &clientMap{
		clients: map[int]struct{}{},
	}
}

func (m *clientMap) add(id int) {
	m.Lock()
	m.clients[id] = struct{}{}
	m.Unlock()
}

func (m *clientMap) getIds() []int {
	ids := []int{}
	m.RLock()
	for id := range m.clients {
		ids = append(ids, id)
	}
	m.RUnlock()

	return ids
}

func (m *clientMap) remove(id int) {
	m.Lock()
	delete(m.clients, id)
	m.Unlock()
}

func main() {
	msnger, err := messenger.NewMessenger()
	if err != nil {
		log.Fatalf("Error getting messenger: %+v.", err)
	}

	var sessionChan <-chan int
	wsServer, sessionChan = websocket.NewWsServer("/temp", "5556")

	go manageSessions(sessionChan)
	receiveReadings(subscribe(msnger))
}

func subscribe(msnger messengertypes.Messenger) <-chan models.TempReading {
	msgChan, err := msnger.Subscribe("sensor")
	if err != nil {
		log.Fatalf("Error subscribing to sensor channel: %+v.", err)
	}

	readingsChan := make(chan models.TempReading)

	go func() {
		for {
			reading := models.TempReading{}
			msg := <-msgChan
			if err := json.Unmarshal([]byte(msg.Payload), &reading); err != nil {
				log.Printf("Error unmarshalling message to reading: %+v.", err)
			}

			readingsChan <- reading
		}
	}()

	return readingsChan
}

func receiveReadings(readingsChan <-chan models.TempReading) {
	for {
		reading := <-readingsChan
		sendReadingToClients(reading)
	}
}

func manageSessions(sessionChan <-chan int) {
	for {
		client := <-sessionChan
		clients.add(client)
	}
}

func sendReadingToClients(reading models.TempReading) {
	for id := range clients.getIds() {
		wsServer.Send(id, reading)
	}
}
