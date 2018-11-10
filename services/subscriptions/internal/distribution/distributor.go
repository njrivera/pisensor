package distribution

import (
	"errors"
	"fmt"
	"log"
	"sync"

	"github.com/pisensor/pkg/models"
)

type Distributor struct {
	clients map[string]chan<- models.TempReading
	sync.RWMutex
}

func NewDistributor() *Distributor {
	return &Distributor{
		clients: map[string]chan<- models.TempReading{},
	}
}

func (d *Distributor) AddClient(name string, ch chan<- models.TempReading) error {
	d.Lock()
	defer d.Unlock()

	if _, ok := d.clients[name]; ok {
		return errors.New(fmt.Sprintf("Client %s already subscribed to distributor", name))
	}
	d.clients[name] = ch

	return nil
}

func (d *Distributor) RemoveClient(name string) {
	d.Lock()
	delete(d.clients, name)
	d.Unlock()
}

func (d *Distributor) DistributeToClients(inChan <-chan models.TempReading) {
	for msg := range inChan {
		d.RLock()
		for name, ch := range d.clients {
			select {
			case ch <- msg:
			default:
				log.Printf("Client %s not ready - dropping message", name)
			}
		}
		d.RUnlock()
	}
}
