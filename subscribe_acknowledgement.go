package zmqps

import "fmt"

func (ps pubSub) SubscribeAcknowledgement(acknowledgement Acknowledgement) error {
	var state = string(acknowledgement.State)
	var message = string(acknowledgement.Message)
	acknowledgementMessageFormat := fmt.Sprintf("%s%s%s", state, separator, message)
	_, err := ps.Socket.Send(acknowledgementMessageFormat, DefaultFlag)
	if err != nil {
		return err
	}
	return nil
}
