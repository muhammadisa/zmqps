package main

import (
	"fmt"
	"github.com/muhammadisa/zmqps"
	zmq "github.com/pebbe/zmq4"
	"math/rand"
	"time"
)

func main() {
	pubSub, err := zmqps.New(zmqps.SUB, "127.0.0.1", "5555")
	if err != nil {
		panic(err)
	}

	randomAcknowledgement := []zmqps.Acknowledgement{zmqps.NACK, zmqps.ACK}

	pubSub.Subscribe(func(msg []byte, err error, socket *zmq.Socket) {
		fmt.Println(string(msg))

		rand.Seed(time.Now().Unix())
		acknlgm := randomAcknowledgement[rand.Intn(len(randomAcknowledgement))]
		pubSub.SubscribeAcknowledgement(acknlgm)
		fmt.Println(acknlgm)
	})

}
