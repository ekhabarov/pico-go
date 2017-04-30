package api

type Config struct {
	ConsoleWidth  uint
	ConsoleHeight uint
	WindowWidth   uint
	WindowHeight  uint
	FPS           uint8
	Verbose       bool
}

type Color int

/*
	This package tries to replicate the pico8 API as closely as possible
	During development I will be trying to implement more an more of the API
	To achieve feature parity with pico8
	Documented extensively here http://pico-8.wikia.com/wiki/Category:API
*/

type PicoGoAPI interface {
	Cls()                 // Clear screen
	ClsColor(color Color) // Clear screen

	//Flip() // Copy graphics buffer to screen

	//Print(str string, x *int, y *int, color *Color) // Print a string of character to the scree
}