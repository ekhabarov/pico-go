package console

import (
	"sync"

	"github.com/veandco/go-sdl2/sdl"
)

type inputter struct {
	sync.Mutex
	keydownMap map[int]bool
}

type Inputter interface {
	PicoInputAPI
	storeKeyState(event sdl.Event) error
}

func NewInputter() Inputter {
	i := &inputter{}

	i.Lock()
	defer i.Unlock()

	// init key map
	i.keydownMap = make(map[int]bool)

	i.keydownMap[P1_BUTT_RIGHT] = false
	i.keydownMap[P1_BUTT_LEFT] = false
	i.keydownMap[P1_BUTT_UP] = false
	i.keydownMap[P1_BUTT_DOWN] = false
	i.keydownMap[P1_BUTT_01] = false
	i.keydownMap[P1_BUTT_02] = false
	i.keydownMap[P1_BUTT_03] = false
	i.keydownMap[P1_BUTT_04] = false
	i.keydownMap[P1_BUTT_05] = false
	i.keydownMap[P1_BUTT_06] = false

	return i
}

func (i *inputter) storeKeyState(event sdl.Event) error {

	switch event.(type) {
	case *sdl.KeyDownEvent:
		// update keydown map
		keyEvent := event.(*sdl.KeyDownEvent)
		i.Lock()
		defer i.Unlock()
		i.keydownMap[int(keyEvent.Keysym.Scancode)] = true
	case *sdl.KeyUpEvent:
		// update keydown map
		keyEvent := event.(*sdl.KeyUpEvent)
		i.Lock()
		defer i.Unlock()
		i.keydownMap[int(keyEvent.Keysym.Scancode)] = false
	}
	return nil
}

func (i *inputter) Btn(id int) bool {
	i.Lock()
	defer i.Unlock()
	if val, ok := i.keydownMap[id]; ok {
		return val
	}
	return false
}
