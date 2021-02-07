package main

import (
	"fmt"
	"log"
	"sharing-services/drivers/nats"
)

func main() {

	encodedConnection, err := nats.NewNatsConnectionClient("192.168.202.128:4222")
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
