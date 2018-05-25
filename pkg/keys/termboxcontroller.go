package keys

import term "github.com/nsf/termbox-go"

type TermBoxController struct {
}

func (*TermBoxController) GetErrorEventType() uint8 {
	return uint8(term.EventError)
}

func (*TermBoxController) GetKeyEventType() uint8 {
	return uint8(term.EventKey)
}

func (*TermBoxController) Init() error{
	return term.Init()
}

func (*TermBoxController) Close() {
	term.Close()
}

func (*TermBoxController) PollEvent() KeyEvent {
	ev := term.PollEvent()
	return KeyEvent{uint8(ev.Type), ev.Ch, ev.Err}
}

func (*TermBoxController) Sync() {
	term.Sync()
}
