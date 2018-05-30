package tests

import (
	"testing"
	"github.com/kmilinho/twcli/pkg/keys"
	"time"
)

func TestListenerProcessTermControllerPolledEvents(t *testing.T) {
	var expectedCharFromEvent rune = 97

	mockTermController := &TestTermController{t,
		keys.KeyEvent{Char: expectedCharFromEvent},
	}

	listener := keys.NewKeyListener(mockTermController)
	listener.Register("a", func(s string) {
		receivedCharFromEvent := []rune(s)[0]
		if receivedCharFromEvent != expectedCharFromEvent {
			t.Error("expected=" + string(expectedCharFromEvent) + " but received=" + string(receivedCharFromEvent))
		}
	})
	listener.Start()
	go func() {
		time.Sleep(time.Second)
		listener.Stop()
	}()
	listener.Wait()
}

type TestTermController struct {
	t     *testing.T
	event keys.KeyEvent
}

func (tc *TestTermController) GetErrorEventType() uint8 {
	return 1
}

func (tc *TestTermController) GetKeyEventType() uint8 {
	return 0
}

func (tc *TestTermController) Init() error {
	return nil
}

func (tc *TestTermController) Close() {
}

func (tc *TestTermController) PollEvent() keys.KeyEvent {
	time.Sleep(200 * time.Millisecond)
	return tc.event
}

func (tc *TestTermController) Sync() {
}
