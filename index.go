package bus

import "github.com/infrago/bus"

func Driver() bus.Driver {
	return &defaultDriver{}
}

func init() {
	bus.Register("default", Driver())
}
