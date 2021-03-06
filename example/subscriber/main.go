package main

import (
	"fmt"
	"github.com/muhammadisa/zmqps"
	"math/rand"
	"time"
)

const (
	messageCorrupt              zmqps.Reason = `message is corrupt`
	messageIsNotAcceptable      zmqps.Reason = `message is not acceptable`
	messageCannotBeUnmarshalled zmqps.Reason = `message is cannot be unmarshalled`
)

func mockingAcknowledgement() zmqps.Acknowledgement {
	rand.Seed(time.Now().Unix())
	s := []zmqps.Acknowledgement{
		{
			State: zmqps.ACK,
		},
		{
			State:   zmqps.NACK,
			Message: messageCorrupt,
		},
		{
			State:   zmqps.NACK,
			Message: messageIsNotAcceptable,
		},
		{
			State:   zmqps.NACK,
			Message: messageCannotBeUnmarshalled,
		},
	}
	return s[rand.Intn(len(s))]
}

func main() {
	pubSub, err := zmqps.New(zmqps.SUB, "127.0.0.1", "5555")
	if err != nil {
		panic(err)
	}

	pubSub.Subscribe(func(msg []byte, err error) {
		if err != nil {
			fmt.Println("ERROR FROM SUBSCRIBER", err)
		}
		fmt.Println(string(msg))
		rand.Seed(time.Now().Unix())
		knowledge := mockingAcknowledgement()
		err = pubSub.SubscribeAcknowledgement(zmqps.Acknowledgement{
			State:   zmqps.NACK,
			Message: messageCannotBeUnmarshalled,
		})
		if err != nil {
			fmt.Println("ERROR SEND ACKNOWLEDGEMENT", err)
		}
		fmt.Println(knowledge)
	})

}
