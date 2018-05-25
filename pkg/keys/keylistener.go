package keys

import (
	"os"
	term "github.com/nsf/termbox-go"
	"log"
	"errors"
)

type KeyListener struct {
	bindings map[string]func(string)
	exit     chan bool
	running  bool
}

func New() *KeyListener {

	if _, found := os.LookupEnv("TERM"); !found {
		os.Setenv("TERM", "xterm")
	}

	bindings := make(map[string]func(string))
	exit := make(chan bool)

	return &KeyListener{
		bindings,
		exit,
		false,
	}
}

func (listener *KeyListener) Register(key string, handler func(string)) (*KeyListener, error) {
	if listener.running {
		return nil, errors.New("cannot register the key handler, key listener already running")
	}
	listener.bindings[key] = handler
	return listener, nil
}

func (listener *KeyListener) Stop() {
	listener.running = false
	listener.exit <- true
}

func (listener *KeyListener) Start() {
	listener.running = true
	go keyEventLoop(listener)
}

func keyEventLoop(listener *KeyListener) {

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
			switch ev := term.PollEvent(); ev.Type {
			case term.EventKey:
				term.Sync()
				key := string(ev.Ch)
				if handler, ok := listener.bindings[key]; ok {
					go handler(key)
				}

				if handler, ok := listener.bindings["*"]; ok {
					go handler(key)
				}
			case term.EventError:
				log.Fatalf("processing keyboard events: %v", ev.Err)
			}
		}
	}
}
