package server

import (
	"log"
	"os"

	messengerfactory "github.com/Messenger/pkg/factory"
	websocketfactory "github.com/Websocket/pkg/factory"
	"github.com/pisensor/subscriptions/internal/control"
	"github.com/pisensor/subscriptions/internal/distribution"
	"github.com/pisensor/subscriptions/internal/receiving"
)

const (
	addrEnv    = "SUB_SERVER_ADDRESS"
	portEnv    = "SUB_PORT"
	msgChannel = "temp_readings"
)

func RunServer() {
	addr := os.Getenv(addrEnv)
	port := os.Getenv(portEnv)

	if addr == "" {
		log.Fatalln("SUB_SERVER_ADDRESS not set")
	}
	if port == "" {
		log.Fatalln("SUB_PORT not set")
	}

	wsServer, err := websocketfactory.NewDefaultWsServer(addr, port)
	if err != nil {
		log.Fatalf("Error starting server: %s.", err.Error())
	}
	msnger, err := messengerfactory.NewDefaultMessenger()
	if err != nil {
		log.Fatalf("Error starting messenger: %s.", err.Error())
	}
	dist := distribution.NewDistributor()

	dataChan := receiving.GoSubscribe(msnger, msgChannel)
	go dist.DistributeToClients(dataChan)

	clientChan := wsServer.ListenForNewClients()
	for c := range clientChan {
		controller := control.NewController(c, dist)
		log.Printf("Starting new websocket client: %s", c.GetName())
		go controller.RunClientController()
	}
}
