package main

// publisher
import (
	"fmt"
	"github.com/vickeyshrestha/sharing-services/drivers/nats"
	"log"

	"time"
)

func main() {
	encodedConnection, err := nats.NewNatsConnectionClient("localhost:5101") // this is your nats msg bus address and port
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
