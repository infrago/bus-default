package bus

import (
	"errors"
	"sync"
	"time"

	"github.com/infrago/bus"
	"github.com/infrago/util"
)

var (
	errRunning    = errors.New("Bus is running.")
	errNotRunning = errors.New("Bus is not running.")
)

type (
	defaultDriver  struct{}
	defaultConnect struct {
		mutex   sync.RWMutex
		running bool
		actives int64

		instance *bus.Instance

		runner *util.Runner
	}
)

// 连接
func (driver *defaultDriver) Connect(inst *bus.Instance) (bus.Connect, error) {
	return &defaultConnect{
		instance: inst, runner: util.NewRunner(),
	}, nil
}

// 打开连接
func (this *defaultConnect) Open() error {
	return nil
}
func (this *defaultConnect) Health() (bus.Health, error) {
	this.mutex.RLock()
	defer this.mutex.RUnlock()
	return bus.Health{Workload: this.actives}, nil
}

// 关闭连接
func (this *defaultConnect) Close() error {
	this.runner.End()
	return nil
}

// 单机版 不用做处理，因为根本就调用不到这边
func (this *defaultConnect) Register(name string) error {
	this.mutex.Lock()
	defer this.mutex.Unlock()

	return nil
}

// 开始
func (this *defaultConnect) Start() error {
	if this.running {
		return errRunning
	}

	this.running = true
	return nil
}

// 结束
func (this *defaultConnect) Stop() error {
	if this.running {
		return errNotRunning
	}

	this.running = false
	return nil
}

// 本来不会执行到这
func (this *defaultConnect) Request(name string, data []byte, timeout time.Duration) ([]byte, error) {
	return nil, nil
}
