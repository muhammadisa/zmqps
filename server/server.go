package main

import (
	zmq "github.com/pebbe/zmq4"
	"log"
)

// SUBSCRIBER
func main() {
	zctx, _ := zmq.NewContext()

	s, _ := zctx.NewSocket(zmq.REP)
	s.Bind("tcp://*:5555")

	for {
		// Wait for next request from client
		msg, _ := s.Recv(0)
		log.Printf("Received %s\n", msg)

		// Send reply back to client
		s.Send("World", 0)
	}
}