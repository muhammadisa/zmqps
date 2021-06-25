package zmqps

import (
	"fmt"
	"github.com/stretchr/testify/suite"
	"testing"
)

type AcknowledgementTestSuite struct {
	suite.Suite
}

func TestAcknowledgementTestSuite(t *testing.T) {
	suite.Run(t, new(AcknowledgementTestSuite))
}

func (ts *AcknowledgementTestSuite) TestReason() {
	tests := []struct {
		A         Acknowledgement
		WantError bool
	}{
		{
			A: Acknowledgement{
				State:   ACK,
				Message: BlankReason,
			},
			WantError: false,
		},
		{
			A: Acknowledgement{
				State:   NACK,
				Message: Reason("cannot be unmarshalled"),
			},
			WantError: true,
		},
		{
			A: Acknowledgement{
				State:   NACK,
				Message: Reason("message is corrupt"),
			},
			WantError: true,
		},
		{
			A: Acknowledgement{
				State:   State("UNKNOWN"),
				Message: Reason("message is corrupt"),
			},
			WantError: true,
		},
	}

	for i, test := range tests {
		ts.Run(fmt.Sprintf("Test number %d of TestReason", i), func() {
			err := test.A.Reason()
			if test.WantError {
				ts.Assert().Error(err)
				ts.Assert().Equal(tests[i].A.Message, test.A.Message, "message should be equal")
			} else {
				ts.Assert().NoError(
					err,
					fmt.Sprintf("error should be nil because state is ACK, state : %+v", test.A.State),
				)
			}
		})
	}
}
