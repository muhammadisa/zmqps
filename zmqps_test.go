package zmqps

import (
	"fmt"
	"github.com/stretchr/testify/suite"
	"testing"
)

type ZmqpsTestSuite struct {
	suite.Suite
}

func TestZmqpsTestSuite(t *testing.T) {
	suite.Run(t, new(ZmqpsTestSuite))
}

func (ts *ZmqpsTestSuite) TestNew() {
	tests := []struct {
		T         Type
		H, P      string
		WantError bool
	}{
		{
			T:         Type(69),
			H:         "127.0.0.1",
			P:         "5555",
			WantError: true,
		},
		// PUB - +
		{
			T:         PUB,
			H:         "127.0.0.1",
			P:         "5555",
			WantError: false,
		},
		{
			T:         PUB,
			H:         "127.0.0.1",
			P:         "asdasd",
			WantError: true,
		},
		// SUB - +
		{
			T:         SUB,
			H:         "127.0.0.1",
			P:         "5555",
			WantError: false,
		},
		{
			T:         SUB,
			H:         "127.0.0.1",
			P:         "asdasd",
			WantError: true,
		},
	}

	for i, test := range tests {
		ts.Run(fmt.Sprintf("Test number %d of TestNew", i), func() {
			_, err := New(test.T, test.H, test.P)
			if test.WantError {
				ts.Assert().Error(err)
			} else {
				ts.Assert().NoError(err)
			}
		})
	}
}
