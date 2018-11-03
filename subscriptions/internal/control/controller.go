package control

import (
	"encoding/json"
	"log"

	websocketTypes "github.com/Websocket/pkg/types"
	"github.com/pisensor/pkg/models"
	"github.com/pisensor/subscriptions/internal/distribution"
	"github.com/pisensor/subscriptions/pkg/types"
)

type Controller struct {
	clientConn  websocketTypes.Client
	receiveChan chan models.TempReading
	distributor *distribution.Distributor
	filter      *types.ServerFilter
	stopChan    chan struct{}
}

func NewController(client websocketTypes.Client, dist *distribution.Distributor) *Controller {
	return &Controller{
		clientConn:  client,
		receiveChan: make(chan models.TempReading),
		distributor: dist,
		filter:      types.NewServerFilter(types.ClientFilter{}),
		stopChan:    make(chan struct{}),
	}
}

func (c *Controller) RunClientController() {
	filterChan := ListenForFilters(c.clientConn, c.stopChan)
	c.distributor.AddClient(c.clientConn.GetName(), c.receiveChan)

	defer c.Close()

	for {
		select {
		case filter, ok := <-filterChan:
			if !ok {
				return
			}
			c.filter = filter
		case reading := <-c.receiveChan:
			if err := SendIfPassFilter(c.clientConn, c.filter, reading); err != nil {
				log.Printf("Error sending reading: %s", err.Error())
				return
			}
		}
	}

}

func (c *Controller) Close() {
	name := c.clientConn.GetName()
	log.Printf("Closing websocket connection for %s", name)

	c.distributor.RemoveClient(name)
	close(c.stopChan)
	c.clientConn.Close()
}

func ListenForFilters(client websocketTypes.Client, stopChan <-chan struct{}) <-chan *types.ServerFilter {
	filterChan := make(chan *types.ServerFilter)

	go func() {
		defer close(filterChan)
		for {
			msg, err := client.Receive()
			if err != nil {
				log.Printf("Error receiving message from client: %s", err.Error())
				return
			}

			var cFilter types.ClientFilter
			if err := json.Unmarshal(msg, &cFilter); err != nil {
				log.Printf("Error unmarshaling message into filter: %s", err.Error())
				continue
			}

			select {
			case <-stopChan:
				return
			case filterChan <- types.NewServerFilter(cFilter):
			}
		}
	}()

	return filterChan
}

func SendIfPassFilter(client websocketTypes.Client, filter *types.ServerFilter, r models.TempReading) error {
	if filter.Check(r) {
		return client.Send(r)
	}

	return nil
}
