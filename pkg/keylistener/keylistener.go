package keylistener

import (
	"os"
	term "github.com/nsf/termbox-go"
	"log"
	"fmt"
)

type KeyListener struct {
	bindings map[string]func(string)
	exit     chan bool
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
	}
}

func (listener *KeyListener) Register(pattern string, handler func(string)) *KeyListener {
	listener.bindings[pattern] = handler
	return listener
}

func (listener *KeyListener) Stop() {
	listener.exit <- true
}

func (listener *KeyListener) Start() {
	go keyEventLoop(*listener)
}

func keyEventLoop(listener KeyListener) {

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
//				term.Sync()
				char string(ev.Ch)
			case term.EventError:
				log.Fatalf("processing keyboard events: %v", ev.Err)
			}
		}
	}
}
