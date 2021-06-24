package zmqps

import (
	"errors"
	"fmt"
	zmq "github.com/pebbe/zmq4"
)

// DefaultFlag is int for set a flag when Recv and Send of zero mq library
const DefaultFlag = 0

// Type is type for represent mode of message queue PUB or SUB
type Type int

const (
	// PUB is type for Publisher mode
	PUB Type = 0
	// SUB is type for subscriber mode
	SUB Type = 1
)

// Listener is function type for subscriber mode for receiving message and error
type Listener func(msg []byte, err error)

// pubSub struct for PubSub interface implementation
type pubSub struct {
	Socket *zmq.Socket
}

// PubSub zmqps main public function definitions
type PubSub interface {
	Publish([]byte) (*zmq.Socket, error)
	Subscribe(Listener)
	PublishAcknowledgement() error
	SubscribeAcknowledgement(Acknowledgement) error
}

// startAsPublisher start socket as publisher mode
func startAsPublisher(zctx *zmq.Context, URL string) (*zmq.Socket, error) {
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

// startAsSubscriber start socket as subscriber mode
func startAsSubscriber(zctx *zmq.Context, URL string) (*zmq.Socket, error) {
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

// New create PubSub zmqps
func New(t Type, host, port string) (PubSub, error) {
	var err error
	var socket *zmq.Socket

	URL := fmt.Sprintf("tcp://%s:%s", host, port)
	zctx, _ := zmq.NewContext()

	switch {
	case t == PUB:
		socket, err = startAsPublisher(zctx, URL)
		if err != nil {
			return nil, err
		}
	case t == SUB:
		socket, err = startAsSubscriber(zctx, URL)
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
