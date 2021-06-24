package zmqps

import (
	"errors"
	"fmt"
	zmq "github.com/pebbe/zmq4"
)

const (
	DefaultFlag = 0
)

type Type int

const (
	PUB Type = 0
	SUB Type = 1
)

type pubSub struct {
	Socket *zmq.Socket
}

type Listener func(msg []byte, err error, socket *zmq.Socket)

type PubSub interface {
	Publish([]byte) (*zmq.Socket, error)
	Subscribe(Listener)
	PublishAcknowledgement() error
	SubscribeAcknowledgement(Acknowledgement) error
}

func connectAsPublisher(zctx *zmq.Context, URL string) (*zmq.Socket, error) {
	socket, err := zctx.NewSocket(zmq.REQ)
	if err != nil {
		return nil, err
	}
	err = socket.Connect(URL)
	if err != nil {
		return nil, err
	}
	return socket, nil
}

func connectAsSubscriber(zctx *zmq.Context, URL string) (*zmq.Socket, error) {
	socket, err := zctx.NewSocket(zmq.REP)
	if err != nil {
		return nil, err
	}
	err = socket.Bind(URL)
	if err != nil {
		return nil, err
	}
	return socket, nil
}

func New(t Type, host, port string) (PubSub, error) {
	var err error
	var socket *zmq.Socket

	URL := fmt.Sprintf("tcp://%s:%s", host, port)
	zctx, _ := zmq.NewContext()

	switch {
	case t == PUB:
		socket, err = connectAsPublisher(zctx, URL)
		if err != nil {
			return nil, err
		}
	case t == SUB:
		socket, err = connectAsSubscriber(zctx, URL)
		if err != nil {
			return nil, err
		}
	default:
		return nil, errors.New("unknown type operation")
	}

	return &pubSub{
		Socket: socket,
	}, nil
}
