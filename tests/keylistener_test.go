package tests

import (
	"testing"
	"github.com/kmilinho/twcli/pkg/keys"
)

func TestRegisterFailsIfListenerIsRunning(t *testing.T) {
	listener := keys.NewKeyListener(&TestTermController{t})
	listener.Start()
	k, err := listener.Register("x", func(s string) {

	})

	if k != nil || err == nil {
		t.Errorf("Register() must fail if listener is running")
	}
}

func TestRegisterDoesNotFailIfListenerIsStopped(t *testing.T) {
	listener := keys.NewKeyListener(&TestTermController{})
	listener.Start()
	listener.Stop()

	k, err := listener.Register("x", func(s string) {

	})

	if k == nil || err != nil {
		t.Errorf("Register() must not fail after stopping the listener")
	}
}

type TestTermController struct {
	t *testing.T
}

func (*TestTermController) GetErrorEventType() uint8 {
	return 0
}

func (*TestTermController) GetKeyEventType() uint8 {
	return 0
}

func (*TestTermController) Init() error{
	return nil
}

func (*TestTermController) Close() {
}

func (*TestTermController) PollEvent() keys.KeyEvent {
	return keys.KeyEvent{}
}

func (*TestTermController) Sync() {
}
