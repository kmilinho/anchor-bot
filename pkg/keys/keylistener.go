package keys

import (
	"os"
	"log"
)

type KeyEvent struct {
	Type uint8
	Char rune
	Err error
}
type TermController interface {
	GetErrorEventType() uint8
	GetKeyEventType() uint8
	Init() error
	Close()
	PollEvent() KeyEvent
	Sync()
}

// struct that holds the state, like the registered event handlers
type KeyListener struct {
	termController TermController
	bindings       map[string]func(string)
	exit           chan bool
	wait           chan bool
}

func NewKeyListener(controller TermController) *KeyListener {
	if _, found := os.LookupEnv("TERM"); !found {
		os.Setenv("TERM", "xterm")
	}

	bindings := make(map[string]func(string))
	exit := make(chan bool)
	wait := make(chan bool)

	return &KeyListener{
		controller,
		bindings,
		exit,
		wait,
	}
}

// creates a new key listener instance with its own state
func NewTermBoxKeyListener() *KeyListener {
	return NewKeyListener(&TermBoxController{})
}

// register event handlers associated to key pressed events
// for example: execute the exit function when the key 'q' is pressed
func (listener *KeyListener) Register(key string, handler func(string)) (*KeyListener, error) {
	listener.bindings[key] = handler
	return listener, nil
}

// stop the key listener, it interrupt the internal event loop
func (listener *KeyListener) Stop() {
	listener.exit <- true
}

// start the key listener, it runs the internal event loop to start listening for key events
func (listener *KeyListener) Start() {
	go keyEventLoop(listener)
}

// calling Wait will block the caller until the Stop() is called
func (listener *KeyListener) Wait() {
	<-listener.wait
}

func keyEventLoop(listener *KeyListener) {
	term := listener.termController

	err := term.Init()
	if err != nil {
		log.Fatalf("starting the termbox app: %v", err)
	}

	defer term.Close()

eventLoop:
	for {
		select {
		case <-listener.exit:
			break eventLoop
		default:
			switch kev := term.PollEvent(); kev.Type {
			case term.GetKeyEventType():
				term.Sync()
				key := string(kev.Char)
				if handler, ok := listener.bindings[key]; ok {
					go handler(key)
				}

				if handler, ok := listener.bindings["*"]; ok {
					go handler(key)
				}
			case term.GetErrorEventType():
				log.Fatalf("processing keyboard events: %v", kev.Err)
			}
		}
	}

	listener.wait <- true
}
