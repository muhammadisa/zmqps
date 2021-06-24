package zmqps

import zmq "github.com/pebbe/zmq4"

// Publish data as []byte to the subscriber return return *zmq.Socket and error
func (ps pubSub) Publish(data []byte) (*zmq.Socket, error) {
	_, err := ps.Socket.Send(string(data), DefaultFlag)
	if err != nil {
		return ps.Socket, err
	}
	return ps.Socket, nil
}
