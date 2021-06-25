package zmqps

// Subscribe listening to publisherA data and have passing data to Listener parameter
func (ps pubSub) Subscribe(listener Listener) {
	for {
		msg, err := ps.Socket.Recv(DefaultFlag)
		listener([]byte(msg), err)
	}
}
