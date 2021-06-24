package zmqps

import "errors"

// separator is string that split State and Reason
const separator = `|`

// Reason is type string for reason of a State and explain what's is going wrong
type Reason string

// Reason constants
const (
	// BlankReason is a default Reason of ACK
	BlankReason Reason = ``
)

// State is type string for represent a condition of a msg that has been published
type State string

// State constants
const (
	// ACK is a state where a msg has been received successfully
	ACK State = `ACK`
	// NACK is a state where a msg has been received with Reason
	NACK State = `NACK`
)

// Acknowledgement is type for wrapping State and Reason when confirming a msg has been received/sent
type Acknowledgement struct {
	State   State
	Message Reason
}

// Reason is a function that determines is State should be an error, and use Acknowledgement.Message as message of error
func (a Acknowledgement) Reason() error {
	switch {
	case a.State == ACK:
		return nil
	case a.State == NACK:
		return errors.New(string(a.Message))
	default:
		return errors.New("unknown acknowledgement status")
	}
}

// StateAsString is a function for convert State to string
func (a Acknowledgement) StateAsString() string {
	return string(a.State)
}
