package main

import (
	"fmt"
	"github.com/muhammadisa/zmqps"
	zmq "github.com/pebbe/zmq4"
)

func main() {
	pubSub, err := zmqps.New(zmqps.SUB, "127.0.0.1", "5555")
	if err != nil {
		panic(err)
	}

	pubSub.Subscribe(func(msg string, err error, socket *zmq.Socket) {
		fmt.Println(msg)
		_, _ = socket.Send("OK", zmqps.DefaultFlag)
	})

}
