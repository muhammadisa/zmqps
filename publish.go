package zmqps

import zmq "github.com/pebbe/zmq4"

func (ps pubSub) Publish(data []byte) (*zmq.Socket, error) {
	_, err := ps.Socket.Send(string(data), DefaultFlag)
	if err != nil {
		return ps.Socket, err
	}
	return ps.Socket, nil
}
