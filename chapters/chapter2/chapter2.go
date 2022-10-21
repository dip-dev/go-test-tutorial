package chapter2

import (
	"github.com/dip-dev/go-test-tutorial/chapters/chapter2/communication"
)

// Chapter2 ..
type Chapter2 struct {
	communication communication.InterfaceCommunication
}

// New ..
func New(communication communication.InterfaceCommunication) *Chapter2 {
	return &Chapter2{
		communication: communication,
	}
}

func (c2 Chapter2) exec() string {
	return c2.communication.Greeting()
}
