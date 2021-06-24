package zmqps

import "fmt"

func (ps pubSub) SubscribeAcknowledgement(acknowledgement Acknowledgement) error {
	acknowledgementMessageFormat := fmt.Sprintf("%s%s%s", acknowledgement.StateAsString(), separator, string(acknowledgement.Message))
	_, err := ps.Socket.Send(acknowledgementMessageFormat, DefaultFlag)
	if err != nil {
		return err
	}
	return nil
}
