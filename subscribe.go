package zmqps

func (ps pubSub) Subscribe(listener Listener) {
	for {
		msg, err := ps.Socket.Recv(DefaultFlag)
		listener([]byte(msg), err, ps.Socket)
	}
}
