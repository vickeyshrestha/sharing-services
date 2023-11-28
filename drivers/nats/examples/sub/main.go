package main

import (
	"fmt"
	"github.com/vickeyshrestha/sharing-services/drivers/nats"
	"log"
)

func main() {

	encodedConnection, err := nats.NewNatsConnectionClient("localhost:5101") // this is your nats msg bus address and port
	if err != nil {
		fmt.Println(err)
		return
	}
	defer encodedConnection.Close()

	// Make sure this type and its properties are exported
	// so the serializer doesn't bork
	type Request struct {
		Id int
	}
	personChanRecv := make(chan *Request)
	encodedConnection.BindRecvChan("request_subject", personChanRecv)

	for {
		// Wait for incoming messages
		req := <-personChanRecv

		log.Printf("Received request: %d", req.Id)
	}
}
