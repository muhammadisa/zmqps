package zmqps

import "fmt"

// SubscribeAcknowledgement send back to the publisher state and reason of the message that has been received
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
