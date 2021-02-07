package main

// publisher
import (
	"fmt"
	"log"
	"sharing-services/drivers/nats"

	"time"
)

func main() {
	encodedConnection, err := nats.NewNatsConnectionClient("192.168.202.128:4222")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer encodedConnection.Close()

	type Request struct {
		Id int
	}
	personChanSend := make(chan *Request)
	encodedConnection.BindSendChan("request_subject", personChanSend)

	i := 0
	for {

		// Create instance of type Request with Id set to
		// the current value of i
		req := Request{Id: i}

		// Just send to the channel! :)
		log.Printf("Sending request %d", req.Id)

		personChanSend <- &req

		// Pause and increment counter
		time.Sleep(time.Second * 1)
		i = i + 1
	}
}
