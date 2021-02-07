package nats

import (
	"errors"
	"github.com/nats-io/nats.go"
	"log"
)

func NewNatsConnectionClient(natsURL ...string) (*nats.EncodedConn, error) {
	var msgBusUrl string
	if len(natsURL) > 1 {
		return nil, errors.New("multiple url detected")
	}
	if len(natsURL) == 1 {
		msgBusUrl = natsURL[0]
	} else {
		msgBusUrl = nats.DefaultURL
	}

	nc, err := nats.Connect(msgBusUrl)
	if err != nil {
		panic(err)
	}

	ec, err := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	if err != nil {
		panic(err)
	}
	//defer ec.Close()

	log.Println("connected to nats...")

	return ec, nil

}
