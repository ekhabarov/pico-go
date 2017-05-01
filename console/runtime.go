package console

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

type runtime struct {
	console *console
}

func newRuntimeMode(c *console) Mode {
	runtime := &runtime{
		console: c,
	}
	return runtime
}

func (r *runtime) PollEvents() error {
	for event = sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch t := event.(type) {
		case *sdl.QuitEvent:
			r.console.hasQuit = true
		case *sdl.KeyDownEvent:
			switch t.Keysym.Sym {
			case sdl.K_RIGHT:
				fmt.Printf("Switching to cli\n")
				r.console.SetMode(CLI)
			}
		// case *sdl.MouseMotionEvent:
		// 	fmt.Printf("[%d ms] MouseMotion\ttype:%d\tid:%d\tx:%d\ty:%d\txrel:%d\tyrel:%d\n",
		// 		t.Timestamp, t.Type, t.Which, t.X, t.Y, t.XRel, t.YRel)
		// case *sdl.MouseButtonEvent:
		// 	fmt.Printf("[%d ms] MouseButton\ttype:%d\tid:%d\tx:%d\ty:%d\tbutton:%d\tstate:%d\n",
		// 		t.Timestamp, t.Type, t.Which, t.X, t.Y, t.Button, t.State)
		// case *sdl.MouseWheelEvent:
		// 	fmt.Printf("[%d ms] MouseWheel\ttype:%d\tid:%d\tx:%d\ty:%d\n",
		// 		t.Timestamp, t.Type, t.Which, t.X, t.Y)
		// case *sdl.KeyDownEvent:
		// 	fmt.Printf("[%d ms] Keyboard\ttype:%d\tsym:%c\tmodifiers:%d\tstate:%d\trepeat:%d\n",
		// 		t.Timestamp, t.Type, t.Keysym.Sym, t.Keysym.Mod, t.State, t.Repeat)
		// case *sdl.KeyUpEvent:
		// 	fmt.Printf("[%d ms] Keyboard\ttype:%d\tsym:%c\tmodifiers:%d\tstate:%d\trepeat:%d\n",
		// 		t.Timestamp, t.Type, t.Keysym.Sym, t.Keysym.Mod, t.State, t.Repeat)
		// case *sdl.JoyAxisEvent:
		// 	fmt.Printf("[%d ms] JoyAxis\ttype:%d\twhich:%c\taxis:%d\tvalue:%d\n",
		// 		t.Timestamp, t.Type, t.Which, t.Axis, t.Value)
		// case *sdl.JoyBallEvent:
		// 	fmt.Printf("[%d ms] JoyBall\ttype:%d\twhich:%d\tball:%d\txrel:%d\tyrel:%d\n",
		// 		t.Timestamp, t.Type, t.Which, t.Ball, t.XRel, t.YRel)
		// case *sdl.JoyButtonEvent:
		// 	fmt.Printf("[%d ms] JoyButton\ttype:%d\twhich:%d\tbutton:%d\tstate:%d\n",
		// 		t.Timestamp, t.Type, t.Which, t.Button, t.State)
		// case *sdl.JoyHatEvent:
		// 	fmt.Printf("[%d ms] JoyHat\ttype:%d\twhich:%d\that:%d\tvalue:%d\n",
		// 		t.Timestamp, t.Type, t.Which, t.Hat, t.Value)
		// case *sdl.JoyDeviceEvent:
		// 	if t.Type == sdl.JOYDEVICEADDED {
		// 		joysticks[int(t.Which)] = sdl.JoystickOpen(t.Which)
		// 		if joysticks[int(t.Which)] != nil {
		// 			fmt.Printf("Joystick %d connected\n", t.Which)
		// 		}
		// 	} else if t.Type == sdl.JOYDEVICEREMOVED {
		// 		if joystick := joysticks[int(t.Which)]; joystick != nil {
		// 			joystick.Close()
		// 		}
		// 		fmt.Printf("Joystick %d disconnected\n", t.Which)
		//	}
		default:
			fmt.Printf("Some event: %#v \n", event)
		}
	}

	return nil

}

func (r *runtime) Update() error {
	return nil
}

func (r *runtime) Render() error {
	r.console.Display.ClsColor(15)
	return nil
}