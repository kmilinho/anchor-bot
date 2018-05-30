package tests

import (
	"testing"
	"container/list"
	"github.com/kmilinho/twcli/pkg/keys"
	"time"
	"fmt"
)

func TestListenerWithTermController(t *testing.T) {
	mockTermController := &TestTermController{t, list.New()}
	listener := keys.NewKeyListener(mockTermController)
	listener.Register("x", func(s string) {})
	listener.Start()
	go func() {
		time.Sleep(time.Second)
		listener.Stop()
		t.Log("listener stopped")
	}()
	listener.Wait()

	for e := mockTermController.interaction.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}

func TestRegisterFailsIfListenerIsRunning(t *testing.T) {
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
	interaction *list.List
}

func (tc *TestTermController) GetErrorEventType() uint8 {
	tc.interaction.PushBack("GetErrorEventType")
	return 1
}

func (tc *TestTermController) GetKeyEventType() uint8 {
	tc.interaction.PushBack("GetKeyEventType")
	return 0
}

func (tc *TestTermController) Init() error{
	tc.interaction.PushBack("Init")
	return nil
}

func (tc *TestTermController) Close() {
	tc.interaction.PushBack("Close")
}

func (tc *TestTermController) PollEvent() keys.KeyEvent {
	tc.interaction.PushBack("PollEvent")
	time.Sleep(200 * time.Millisecond)
	return keys.KeyEvent{}
}

func (tc *TestTermController) Sync() {
	tc.interaction.PushBack("Sync")
}
