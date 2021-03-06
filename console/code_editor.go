package console

import (
	"github.com/veandco/go-sdl2/sdl"
)

type codeEditor struct {
	console *console
	PixelBuffer
}

func newCodeEditorMode(c *console) Mode {
	codeEditor := &codeEditor{
		console: c,
	}
	pb, _ := newPixelBuffer(c.Config)

	codeEditor.PixelBuffer = pb

	return codeEditor
}

func (c *codeEditor) HandleEvent(event sdl.Event) error {
	switch t := event.(type) {
	case *sdl.KeyDownEvent:
		switch t.Keysym.Sym {
		// case sdl.K_RIGHT:
		// 	fmt.Printf("Switching to runtime\n")
		// 	c.console.SetMode(RUNTIME)
		}
	default:
		//fmt.Printf("Some event: %#v \n", event)
	}

	return nil
}

func (c *codeEditor) Init() error {
	return nil
}

func (c *codeEditor) Update() error {
	return nil
}

func (c *codeEditor) Render() error {
	c.PixelBuffer.ClsWithColor(5)
	c.PixelBuffer.PrintAtWithColor("Code editor Print Test", 10, 10, 7)
	renderModeHeader(c)
	return nil
}
