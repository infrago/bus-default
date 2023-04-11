package bus

import (
	"github.com/infrago/bus"
	"github.com/infrago/infra"
)

func Driver() bus.Driver {
	return &defaultDriver{}
}

func init() {
	infra.Register("default", Driver())
}
