package zmqps

import "strings"

func (ps pubSub) PublishAcknowledgement() error {
	var acknowledgement Acknowledgement
	msg, err := ps.Socket.Recv(DefaultFlag)
	if err != nil {
		return err
	}
	messages := strings.Split(msg, separator)
	if len(messages) > 1 {
		acknowledgement.State = State(messages[0])
		acknowledgement.Message = Reason(messages[1])
	}
	acknowledgement.State = State(msg)
	acknowledgement.Message = BlankReason
	return acknowledgement.Reason()
}
