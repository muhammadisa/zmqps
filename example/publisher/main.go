package main

import (
	"encoding/json"
	fmt "fmt"
	"github.com/google/uuid"
	"github.com/muhammadisa/zmqps"
	"time"
)

func sendEventTime(pubSub zmqps.PubSub) {
	jsonByte, _ := json.Marshal(struct {
		UUID string `json:"uuid"`
		Time int64  `json:"time"`
	}{
		UUID: uuid.New().String(),
		Time: time.Now().Unix(),
	})

	_, err := pubSub.Publish(jsonByte)
	if err != nil {
		fmt.Println(err)
		sendEventTime(pubSub)
	}
}

func main() {
	pubSub, err := zmqps.New(zmqps.PUB, "127.0.0.1", "5555")
	if err != nil {
		panic(err)
	}

	sendEventTime(pubSub)
}
