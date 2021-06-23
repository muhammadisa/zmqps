package main

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/muhammadisa/zmqps"
	"time"
)

func main() {
	pubSub, err := zmqps.New(zmqps.PUB, "127.0.0.1", "5555")
	if err != nil {
		panic(err)
	}

	for i := 0; i < 1000000; i++ {
		jsonByte, _ := json.Marshal(struct {
			UUID string `json:"uuid"`
			Time int64  `json:"time"`
		}{
			UUID: uuid.New().String(),
			Time: time.Now().Unix(),
		})

		socket, _ := pubSub.Publish(jsonByte)
		msg, _ := socket.Recv(zmqps.DefaultFlag)
		fmt.Println(msg)
	}
}
